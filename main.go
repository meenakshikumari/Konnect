package main

import "os"

func main() {
	err := os.Setenv("DATABASE_URL", "host=127.0.0.1 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return
	}
	err = runDatabaseMigrations()
	if err != nil {
		return
	}
	StartServer()
}
