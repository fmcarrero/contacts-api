package acceptance

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fmcarrero/contacts-api/src"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/controller/dto"
	"github.com/fmcarrero/contacts-api/tests/builder"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllContacts(t *testing.T) {
	server := getServer()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/contacts", nil)
	assert.NoError(t, err, "err should be null when calling to get all contacts endpoint")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application-ID", "acceptance-test")
	server.Server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var contacts dto.GetAllContactsDTO
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&contacts))
	assert.Len(t, contacts.Data, 1)
	assert.EqualValues(t, contacts.Total, len(contacts.Data))
	assert.EqualValues(t, "ma@gmail.com", contacts.Data[0].Email)
	assert.EqualValues(t, 1795068458442948608, contacts.Data[0].ID)
}

func TestAddContact(t *testing.T) {
	t.Run("Add contact", func(t *testing.T) {
		server := getServer()
		w := httptest.NewRecorder()
		var body bytes.Buffer
		expected := builder.NewAddContactRequestBuilder().Build()
		assert.NoError(t, json.NewEncoder(&body).Encode(expected))
		req, err := http.NewRequest(http.MethodPost, "/v1/contacts", &body)
		assert.NoError(t, err, "err should be null when calling to add contact endpoint")

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Application-ID", "acceptance-test")
		server.Server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var contact dto.Contact
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&contact))
		assert.EqualValues(t, expected.Email, contact.Email)
		_, err = server.GetDependencies().Conn.Exec(context.Background(), "delete from contacts where id = $1", contact.ID)
		assert.NoError(t, err)
	})
	t.Run("Add contact with invalid email", func(t *testing.T) {
		var customErr src.Error
		server := getServer()
		w := httptest.NewRecorder()
		var body bytes.Buffer
		expected := builder.NewAddContactRequestBuilder().Email("invalid-email").Build()
		assert.NoError(t, json.NewEncoder(&body).Encode(expected))
		req, err := http.NewRequest(http.MethodPost, "/v1/contacts", &body)
		assert.NoError(t, err, "err should be null when calling to add contact endpoint")

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Application-ID", "acceptance-test")
		server.Server.ServeHTTP(w, req)
		assert.NoError(t, json.NewDecoder(w.Body).Decode(&customErr))
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.EqualValues(t, "contact.validation.email_format.error", customErr.Type)
	})
}
func getServer() *src.Server {
	server := src.NewServer(src.Build())
	server.Routes()
	server.Middlewares(src.WithRecover(), src.WithRequestID(), src.WithErrorHandler())

	return server
}
