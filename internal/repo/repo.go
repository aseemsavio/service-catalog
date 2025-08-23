package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service struct {
	ServiceUUID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"uniqueIndex;not null"`
	Description string    `gorm:"type:varchar(150);not null"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null;default:now()"`
	Versions    []Version `gorm:"foreignKey:ServiceUUID;constraint:OnDelete:CASCADE"`
	// Read-side aggregates
	LatestPublishedOn *time.Time `gorm:"-"`
	VersionCount      int64      `gorm:"-"`
}

type Version struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceUUID uuid.UUID `gorm:"type:uuid;index;not null"`
	Name        string    `gorm:"not null"`
	PublishedOn time.Time `gorm:"type:date;not null"`
}

type Repo struct{ db *gorm.DB }

func Open(dsn string) (*Repo, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}
	return &Repo{db: db}, nil
}

type ListOpts struct {
	Query    string
	SortBy   string // name|created_at|updated_at|latest_published_on|version_count
	Order    string // asc|desc
	Page     int
	PageSize int
}

func orderOr(s string) string {
	if s == "desc" {
		return "desc"
	}
	return "asc"
}

func (r *Repo) ListServices(ctx context.Context, o ListOpts) (items []Service, total int64, err error) {
	if o.Page <= 0 {
		o.Page = 1
	}
	if o.PageSize <= 0 || o.PageSize > 100 {
		o.PageSize = 20
	}

	// Count (with query filter)
	countQ := r.db.WithContext(ctx).Model(&Service{})
	if o.Query != "" {
		like := "%" + o.Query + "%"
		countQ = countQ.Where("name ILIKE ? OR description ILIKE ?", like, like)
	}
	if err = countQ.Count(&total).Error; err != nil {
		return
	}

	// Select aggregates and sort
	raw := r.db.WithContext(ctx).Model(&Service{}).
		Select(`services.*,
      (SELECT MAX(published_on) FROM versions v WHERE v.service_uuid = services.service_uuid) AS latest_published_on,
      (SELECT COUNT(1) FROM versions v2 WHERE v2.service_uuid = services.service_uuid) AS version_count`)

	if o.Query != "" {
		like := "%" + o.Query + "%"
		raw = raw.Where("services.name ILIKE ? OR services.description ILIKE ?", like, like)
	}

	switch o.SortBy {
	case "name", "created_at", "updated_at":
		raw = raw.Order("services." + o.SortBy + " " + orderOr(o.Order))
	case "latest_published_on":
		raw = raw.Order("latest_published_on " + orderOr(o.Order))
	case "version_count":
		raw = raw.Order("version_count " + orderOr(o.Order))
	default:
		raw = raw.Order("services.name asc")
	}

	err = raw.Offset((o.Page - 1) * o.PageSize).Limit(o.PageSize).Find(&items).Error
	return
}

func (r *Repo) GetService(ctx context.Context, id uuid.UUID) (*Service, error) {
	var s Service
	if err := r.db.WithContext(ctx).First(&s, "service_uuid = ?", id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repo) ListVersions(ctx context.Context, serviceID uuid.UUID) ([]Version, error) {
	var vs []Version
	if err := r.db.WithContext(ctx).
		Where("service_uuid = ?", serviceID).
		Order("published_on DESC, name DESC").
		Find(&vs).Error; err != nil {
		return nil, err
	}
	return vs, nil
}
