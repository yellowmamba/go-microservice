package main

import (
	"myob/internal/api/openapi"
	"myob/internal/api/server"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

func main() {
	api, err := server.NewApiServer()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to bind api implementation")
	}
	router := echo.New()
	openapi.RegisterHandlers(router, api)
	router.Logger.Fatal(router.Start(":1323"))
}
