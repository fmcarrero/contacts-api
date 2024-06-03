package ping

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// HandierPing managements healthcheck endpoints
type HandierPing interface {
	Ping(c echo.Context) error
}

type response struct {
	Version string    `json:"version"`
	Name    string    `json:"name"`
	Uptime  time.Time `json:"uptime"`
}

type statusHandler struct {
	projectName    string
	projectVersion string
}

// NewHandierPing creates a new HandierPing instance
func NewHandierPing(projectName, projectVersion string) HandierPing {
	return &statusHandler{
		projectName:    projectName,
		projectVersion: projectVersion,
	}
}

// Ping checks if service  is up
func (s *statusHandler) Ping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, response{
		Version: s.projectVersion,
		Name:    s.projectName,
		Uptime:  time.Now().UTC(),
	})
}
