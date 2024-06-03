package src

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Server       *echo.Echo
	dependencies Dependencies
}

func NewServer(dependencies Dependencies) *Server {
	return &Server{
		Server:       echo.New(),
		dependencies: dependencies,
	}
}

// Start run the server
func (s *Server) Start() {
	s.Server.Logger.Fatal(s.Server.Start(fmt.Sprintf(":%s", s.dependencies.Config.Port)))
}
func (s *Server) GetDependencies() Dependencies {
	return s.dependencies
}
