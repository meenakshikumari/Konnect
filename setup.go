package main

import (
	"api/config"
	"api/handlers"
	"api/internal/services"
	"api/pkg/server"
	"api/stores"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	config.Load()

	deps := InitializeDependencies()
	srv := server.New(handlers.NewRouter(*deps))
	logrus.Infof("server listening on %s.", config.Addr())
	srv.Serve(config.Addr())
}

func InitializeDependencies() *handlers.Dependencies {
	db := InitializeDB()
	repos := initializeRepositories(db)
	deps := &handlers.Dependencies{
		Db:             db.DB,
		APIService:     repos.Services,
		ServiceVersion: repos.Versions,
	}
	return deps
}

func InitializeDB() *stores.DataBase {
	db, err := stores.NewDBHandle()
	if err != nil {
		logrus.Fatalf("failed to initialize: %s", err)
	}
	return db
}

func initializeRepositories(db *stores.DataBase) *repositories {
	return &repositories{
		Services: services.APIService{ServiceRepo: stores.NewServiceStore(db)},
		Versions: services.ServiceVersion{VersionRepo: stores.NewVersionStore(db)},
	}
}

type repositories struct {
	Services services.APIService
	Versions services.ServiceVersion
}
