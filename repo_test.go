package service_catalog_test

import (
	"context"
	"services-catalog/internal/repo"
	"services-catalog/internal/testcontainer"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRepoIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Setup test container
	postgres, err := testcontainer.SetupPostgres(t)
	require.NoError(t, err)

	ctx := context.Background()
	defer postgres.Cleanup(ctx)

	// Setup schema
	require.NoError(t, postgres.DB.AutoMigrate(&repo.Service{}, &repo.Version{}))

	// Create repository
	repository, err := repo.Open(postgres.DSN)
	require.NoError(t, err)

	// Test empty list
	t.Run("EmptyList", func(t *testing.T) {
		services, total, err := repository.ListServices(ctx, repo.ListOpts{})
		require.NoError(t, err)
		assert.Equal(t, int64(0), total)
		assert.Empty(t, services)
	})

	// Seed data for remaining tests
	testServices := seedTestData(t, postgres.DB)

	t.Run("ListAllServices", func(t *testing.T) {
		services, total, err := repository.ListServices(ctx, repo.ListOpts{
			Page:     1,
			PageSize: 10,
		})
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Len(t, services, 2)
	})

	t.Run("SearchServices", func(t *testing.T) {
		services, total, err := repository.ListServices(ctx, repo.ListOpts{
			Query:    "API",
			Page:     1,
			PageSize: 10,
		})
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Len(t, services, 1)
		assert.Equal(t, "User API", services[0].Name)
	})

	t.Run("PaginateServices", func(t *testing.T) {
		services, total, err := repository.ListServices(ctx, repo.ListOpts{
			Page:     1,
			PageSize: 1,
		})
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Len(t, services, 1)
	})

	t.Run("SortServices", func(t *testing.T) {
		services, _, err := repository.ListServices(ctx, repo.ListOpts{
			SortBy: "created_at",
			Order:  "desc",
		})
		require.NoError(t, err)
		assert.Len(t, services, 2)
		assert.True(t, services[0].CreatedAt.After(services[1].CreatedAt))
	})

	t.Run("GetExistingService", func(t *testing.T) {
		service, err := repository.GetService(ctx, testServices[0].ServiceUUID)
		require.NoError(t, err)
		assert.Equal(t, testServices[0].Name, service.Name)
		assert.Equal(t, testServices[0].Description, service.Description)
		assert.Len(t, service.Versions, 2)
	})

	t.Run("GetNonExistentService", func(t *testing.T) {
		randomUUID := uuid.New()
		service, err := repository.GetService(ctx, randomUUID)
		assert.Error(t, err)
		assert.Nil(t, service)
	})
}

func seedTestData(t *testing.T, db *gorm.DB) []repo.Service {
	now := time.Now()

	// Create services with versions
	services := []repo.Service{
		{
			ServiceUUID: uuid.New(),
			Name:        "User API",
			Description: "User management API service",
			CreatedAt:   now.Add(-48 * time.Hour),
			UpdatedAt:   now.Add(-24 * time.Hour),
			Versions: []repo.Version{
				{
					ID:          uuid.New(),
					Name:        "v1.0.0",
					PublishedOn: now.Add(-48 * time.Hour),
				},
				{
					ID:          uuid.New(),
					Name:        "v1.1.0",
					PublishedOn: now.Add(-24 * time.Hour),
				},
			},
		},
		{
			ServiceUUID: uuid.New(),
			Name:        "Payment Service",
			Description: "Payment processing service",
			CreatedAt:   now.Add(-12 * time.Hour),
			UpdatedAt:   now.Add(-6 * time.Hour),
			Versions: []repo.Version{
				{
					ID:          uuid.New(),
					Name:        "v1.0.0",
					PublishedOn: now.Add(-12 * time.Hour),
				},
			},
		},
	}

	for _, service := range services {
		require.NoError(t, db.Create(&service).Error)
	}

	return services
}
