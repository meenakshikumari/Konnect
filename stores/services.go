package stores

import (
	"api/internal/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func NewServiceStore(db *DataBase) *ServiceStore {
	return &ServiceStore{
		db: db,
	}
}

type ServiceStore struct {
	db *DataBase
}

type Service struct {
	ID           int64      `db:"id"`
	Name         string     `db:"name"`
	Description  string     `db:"description"`
	Published    bool       `db:"published"`
	VersionCount int        `db:"version_count"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}

func (s *ServiceStore) FindByID(id int64) (*models.Service, error) {
	query := `SELECT id, name, description, published, version_count 
		FROM services 
		WHERE id = $1`

	var service Service
	err := s.db.DB.GetContext(context.Background(), &service, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Infof("[ServiceStore][FindByID] No Service Data found")
			return nil, fmt.Errorf("no Service Data found")
		}
		logrus.Errorf("[ServiceStore][FindAll] Failed to execute query: %v", err)
		return nil, err
	}

	resp := &models.Service{
		ID:           service.ID,
		Name:         service.Name,
		Description:  service.Description,
		Published:    service.Published,
		VersionCount: service.VersionCount,
	}
	return resp, nil
}

func (s *ServiceStore) FindAll(limit, offset int, sortOn, sortBy, nameContains string) ([]models.Service, error) {
	var res []models.Service
	offset = (offset - 1) * limit

	query := `SELECT id, name, description, published, version_count 
		FROM services 
		WHERE name ILIKE $1
		ORDER BY %s %s LIMIT $2 OFFSET $3`

	nameContains = "%" + nameContains + "%"
	fullQuery := fmt.Sprintf(query, sortOn, sortBy)
	rows, err := s.db.DB.Query(fullQuery, nameContains, limit, offset)

	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Infof("[ServiceStore][FindAll] No Service Data found")
			return res, nil
		}
		logrus.Errorf("[ServiceStore][FindAll] Failed to execute query: %v", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var svc models.Service
		err = rows.Scan(&svc.ID, &svc.Name, &svc.Description, &svc.Published, &svc.VersionCount)
		if err != nil {
			logrus.Errorf("[ServiceStore][FindAll] Failed to scan row: %v", err)
			return nil, err
		}

		svcModel := models.Service{
			ID:           svc.ID,
			Name:         svc.Name,
			Description:  svc.Description,
			Published:    svc.Published,
			VersionCount: svc.VersionCount,
		}
		res = append(res, svcModel)
	}
	return res, nil
}
