package stores

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"os"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sqlx.DB
}

func NewDBHandle() (*DataBase, error) {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load the database")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to do ping check to database")
	}

	return &DataBase{
		DB: db,
	}, nil
}
