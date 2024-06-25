package models

import "time"

type Version struct {
	ID          int64
	Name        string
	Description string
	Published   bool
	SessionID     int64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time `db:"updated_at"`
}

func NewVersions(versions []Version) *Versions {
	if versions == nil {
		return &Versions{versions: []Version{}}
	}
	return &Versions{versions: versions}
}

type Versions struct {
	versions []Version
}
