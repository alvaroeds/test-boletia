package main

import (
	"errors"
	"github.com/alvaroeds/test-boletia/internal/config"
	"github.com/alvaroeds/test-boletia/internal/db/postgres"
	"github.com/alvaroeds/test-boletia/internal/http"
	"github.com/alvaroeds/test-boletia/pkg/provider"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const migrationsFolder = "file://migrations"

func main() {
	conf := config.GetConfig()

	err := doMigrate(conf.PostgresURL)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	postgres, err := postgres.NewPostgresClient(conf.PostgresURL)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	p := provider.New(conf, postgres.DB)
	p.Load()

	err = http.Start(conf, postgres)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func doMigrate(databaseURL string) error {
	m, err := migrate.New(
		migrationsFolder,
		databaseURL,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
