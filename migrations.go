package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

const appDatabaseMigrationPath = "file://db/migrations"

func runDatabaseMigrations() error {
	m := createMigrate()
	err := m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			logrus.Warnf("no changes needed")
			return nil
		}
		logrus.Fatalf("migration failed: %s", err)
	}

	logrus.Infof("migration successful")

	return nil
}

func createMigrate() *migrate.Migrate {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatalf("ping to the database host failed: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(appDatabaseMigrationPath, "postgres", driver)
	if err != nil {
		logrus.Fatalf("failed to prepare migration: %s", err)
	}

	return m
}
