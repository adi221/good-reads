package main

import (
	"net/http"

	"github.com/adi221/good-reads/pkg/api"
	"github.com/adi221/good-reads/pkg/config"
	"github.com/adi221/good-reads/pkg/db"
	"github.com/adi221/good-reads/pkg/service"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.NewConfig()

	database, err := db.NewDB(conf.General.DatabaseURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init database")
	}
	defer database.Close()

	err = service.Configure(*conf, database)
	if err != nil {
		database.Close()
		log.Fatal().Err(err).Msg("Failed to init service registry")
	}

	server := &http.Server{
		Addr:    conf.General.ListenAddr,
		Handler: api.NewRouter(),
	}

	log.Info().Msg("Server listening on port: " + conf.General.ListenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("could not start the server")
	}
}
