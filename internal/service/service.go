package service

import (
	"context"
	"errors"
	"services-catalog/internal/repo"

	"github.com/google/uuid"
)

// Svc provides service layer methods
type Svc struct{ repo *repo.Repo }

// New creates a new service instance
func New(repo *repo.Repo) *Svc { return &Svc{repo: repo} }

// ListOpts defines options for listing services
type ListOpts = repo.ListOpts

// ServiceDTO is the data transfer object for a service
type ServiceDTO struct {
	ServiceUUID uuid.UUID    `json:"service_uuid"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Versions    []VersionDTO `json:"versions"`
}

// List retrieves a paginated list of services based on the provided options
func (s *Svc) List(ctx context.Context, o ListOpts) ([]ServiceDTO, int64, error) {
	items, total, err := s.repo.ListServices(ctx, o)
	if err != nil {
		return nil, 0, err
	}
	out := make([]ServiceDTO, 0, len(items))
	for _, it := range items {
		versions := make([]VersionDTO, 0, len(it.Versions))
		for _, v := range it.Versions {
			versions = append(versions, VersionDTO{
				ID:          v.ID,
				Name:        v.Name,
				PublishedOn: v.PublishedOn.Format("2006-01-02"),
			})
		}
		out = append(out, ServiceDTO{
			ServiceUUID: it.ServiceUUID,
			Name:        it.Name,
			Description: it.Description,
			Versions:    versions,
		})
	}
	return out, total, nil
}

// Get retrieves a service by its UUID
func (s *Svc) Get(ctx context.Context, id uuid.UUID) (ServiceDTO, error) {
	it, err := s.repo.GetService(ctx, id)
	if err != nil {
		return ServiceDTO{}, err
	}
	versions := make([]VersionDTO, 0, len(it.Versions))
	for _, v := range it.Versions {
		versions = append(versions, VersionDTO{
			ID:          v.ID,
			Name:        v.Name,
			PublishedOn: v.PublishedOn.Format("2006-01-02"),
		})
	}
	return ServiceDTO{
		ServiceUUID: it.ServiceUUID,
		Name:        it.Name,
		Description: it.Description,
		Versions:    versions,
	}, nil
}

// VersionDTO is the data transfer object for a version
type VersionDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	PublishedOn string    `json:"published_on"`
}

// ParseUUID parses a UUID from string and returns an error if invalid
func ParseUUID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, errors.New("invalid uuid")
	}
	return id, nil
}
