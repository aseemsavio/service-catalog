package service

import (
	"context"
	"errors"
	"services-catalog/internal/repo"
	"time"

	"github.com/google/uuid"
)

type Svc struct{ repo *repo.Repo }

func New(repo *repo.Repo) *Svc { return &Svc{repo: repo} }

type ListOpts = repo.ListOpts

type ServiceDTO struct {
	ServiceUUID uuid.UUID    `json:"service_uuid"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Versions    []VersionDTO `json:"versions"`
}

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

type VersionDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	PublishedOn string    `json:"published_on"`
}

func ParseUUID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, errors.New("invalid uuid")
	}
	return id, nil
}

// Utility mainly for tests
func Ptr[T any](v T) *T { return &v }

// (Optionally) a small helper to normalize date to midnight for input paths later
func normalizeDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
