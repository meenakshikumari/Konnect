package main

import (
	"api/config"
	"api/handlers"
	"api/internal/services"
	"api/pkg/server"
	"api/stores"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newCLI() *cobra.Command {
	cli := &cobra.Command{
		Use:   "Konnect",
		Short: "Konnect is a Service",
	}

	cli.AddCommand(newServeCmd())
	cli.AddCommand(newMigrateCmd())

	return cli
}

func newServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "server",
		Short:   "Start HTTP API server",
		Aliases: []string{"serve", "start"},
		Run: func(_ *cobra.Command, _ []string) {
			StartServer()
		},
	}
}

func newMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Perform db migration",
		Run: func(_ *cobra.Command, _ []string) {
			config.Load()
			err := runDatabaseMigrations()
			if err != nil {
				return
			}
		},
	}
}

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
