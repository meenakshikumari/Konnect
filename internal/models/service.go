package models

import "time"

type Service struct {
	ID           int64
	Name         string
	Description  string
	Published    bool
	VersionCount int
	CreatedAt    *time.Time
	UpdatedAt    *time.Time `db:"updated_at"`
}

func NewServices(services []Service) *Services {
	if services == nil {
		return &Services{services: []Service{}}
	}
	return &Services{services: services}
}

type Services struct {
	services []Service
}
