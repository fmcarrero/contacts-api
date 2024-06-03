package src

import "fmt"

// Routes build the routes of the server
func (s *Server) Routes() {
	v1 := "v1"
	s.Server.GET("/ping", s.dependencies.HandierPing.Ping)
	rootPrefix := s.Server.Group(fmt.Sprintf("/%s/contacts", v1))

	rootPrefix.GET("", s.dependencies.ContactHandler.GetContacts)
	rootPrefix.POST("", s.dependencies.ContactHandler.AddContact)
	rootPrefix.PUT("/:id", s.dependencies.ContactHandler.EditContact)
	rootPrefix.DELETE("/:id", s.dependencies.ContactHandler.RemoveContact)
}
