package stores

import (
	"api/internal/models"
	"context"
	"time"
)

func NewVersionStore(db *DataBase) *VersionStore {
	return &VersionStore{
		db: db,
	}
}

type VersionStore struct {
	db *DataBase
}

type Version struct {
	ID          int64  `db:"id"`
	Name        string 	`db:"name"`
	Description string `db:"description"`
	Published   bool `db:"published"`
	ServiceID     int64 `db:"service_id"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

func (v *VersionStore)FindByID(ctx context.Context, id int64) (*models.Version, error){
	return nil, nil

}

//func (v *VersionStore)FindAllVersions(ctx context.Context, serviceID int64) (models.Versions, error){
//	return nil, nil
//}