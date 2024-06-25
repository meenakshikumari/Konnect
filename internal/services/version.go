package services

import (
	"api/internal/models"
	"context"
)

type ServiceVersionRepository interface {
	FindByID(ctx context.Context, id int64) (*models.Version, error)
	//FindAllVersions(ctx context.Context, serviceID int64) (models.Versions, error)
}

type ServiceVersion struct {
	VersionRepo ServiceVersionRepository
}
