package src

import "fmt"

// Routes build the routes of the server
func (s *Server) Routes() {
	// api gateway
	v1 := "v1"
	rootPrefix := s.Server.Group(fmt.Sprintf("/%s/contacts", v1))

	rootPrefix.GET("", s.dependencies.ContactHandler.GetContacts)
	rootPrefix.POST("", s.dependencies.ContactHandler.AddContact)
	rootPrefix.PUT("/:id", s.dependencies.ContactHandler.EditContact)
}
