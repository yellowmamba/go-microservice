package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &ApiServer{
		Config: &ServerConfig{},
	}

	// Assertions
	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "hello world", rec.Body.String())
	}
}

func TestGetHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")
	h := &ApiServer{
		Config: &ServerConfig{
			InMaintenance: true,
		},
	}

	// Assertions
	if assert.NoError(t, h.GetHealth(c)) {
		assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
	}
}

func TestGetMetadata(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/metadata")
	h := &ApiServer{
		Config: &ServerConfig{
			LastCommit: "latest",
			ApiVersion: "version",
		},
	}

	expected := `{
		"description": "This is a web api.",
		"version": "version",
		"lastcommitsha": "latest"
	}`
	// Assertions
	if assert.NoError(t, h.GetMetadata(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
