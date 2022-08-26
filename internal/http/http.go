package http

import (
	"github.com/alvaroeds/test-boletia/internal/config"
	"github.com/alvaroeds/test-boletia/internal/db/postgres"
)

func Start(conf *config.Config, dbClient *postgres.Client) error {
	r := routes(dbClient)

	server := newServer(conf.HttpPort, r)

	server.Start()

	return nil
}
