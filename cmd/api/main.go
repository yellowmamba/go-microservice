package main

import (
	"flag"
	"fmt"
	"myob/internal/api/openapi"
	"myob/internal/api/server"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

const DefaultAppPort int = 1111

var CommitHash string

func main() {
	inMaintenance := flag.Bool("maintenance", false, "app under maintenance")
	runPort := flag.Int("port", DefaultAppPort, "port that the app is running on")
	runLocal := flag.Bool("local", false, "localhost testing")
	flag.Parse()

	api, err := server.NewApiServer(
		&server.ServerConfig{
			InMaintenance: *inMaintenance,
			LastCommit:    CommitHash,
			ApiVersion:    server.AppVersion,
		},
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to bind api implementation")
	}
	router := echo.New()
	openapi.RegisterHandlers(router, api)

	if *runLocal {
		// start a local server
		router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", *runPort)))
		return
	}

	// run it on production
}
