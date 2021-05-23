package server

import (
	"myob/internal/api/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiServer struct {
	Config *ServerConfig
}

type ServerConfig struct {
	ApiVersion    string
	LastCommit    string
	InMaintenance bool
}

const AppDescription = "This is a web api."

func NewApiServer() (*ApiServer, error) {
	return &ApiServer{
		Config: &ServerConfig{
			ApiVersion:    "1.0.0",
			LastCommit:    "abcdefg",
			InMaintenance: false,
		},
	}, nil
}

func (as *ApiServer) Get(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello world")
}

func (as *ApiServer) GetHealth(ctx echo.Context) error {
	if as.Config.InMaintenance {
		return ctx.JSON(http.StatusServiceUnavailable, "Site under maintenance!")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (as *ApiServer) GetMetadata(ctx echo.Context) error {
	metadata := &openapi.AppMetadata{
		Description:   AppDescription,
		Lastcommitsha: as.Config.LastCommit,
		Version:       as.Config.ApiVersion,
	}
	return ctx.JSON(http.StatusOK, metadata)
}
