package app

import (
	"net"
	"net/http"
	"os"

	"github.com/cleysonph/hyperprof/config"
	"github.com/cleysonph/hyperprof/internal/database"
	"github.com/cleysonph/hyperprof/internal/transport/middleware"
	"github.com/cleysonph/hyperprof/internal/transport/rest"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run() {
	if config.IsDev() {
		if err := godotenv.Load(); err != nil {
			log.Fatal().Err(err).Msg("Error loading .env file")
		}
	}

	config.Init()

	database.InitMySQL(config.Dsn)
	defer database.Close()

	if config.Env == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	listener, err := net.Listen("tcp", config.Addr())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	log.Info().Msgf("Listening on %s", listener.Addr().String())
	handler := middleware.HttpCors(
		middleware.HttpLogger(
			rest.NewRouter(),
		),
	)

	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
