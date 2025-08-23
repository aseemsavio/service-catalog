package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ServiceUUID       uuid.UUID
	Name              string
	Description       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LatestPublishedOn *time.Time
	VersionCount      int64
}

type Version struct {
	ID          uuid.UUID
	ServiceUUID uuid.UUID
	Name        string
	PublishedOn time.Time
}

func (s *Service) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return errors.New("name is required")
	}
	if len(s.Description) == 0 {
		return errors.New("description is required")
	}
	if len(s.Description) > 150 {
		return errors.New("description must be <= 150 chars")
	}
	return nil
}
