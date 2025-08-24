package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Service / represents a service with its versions
type Service struct {
	ServiceUUID uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Versions    []Version
}

// Version represents a version of a service
type Version struct {
	ID          uuid.UUID
	ServiceUUID uuid.UUID
	Name        string
	PublishedOn time.Time
}

// Validate checks if the Service fields meet the required constraints
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
