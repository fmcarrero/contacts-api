package integration

import (
	"context"
	"fmt"
	customError "github.com/fmcarrero/contacts-api/src/contacts/domain/error"
	"testing"

	"github.com/fmcarrero/contacts-api/src"
	"github.com/fmcarrero/contacts-api/src/contacts/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetAllContacts(t *testing.T) {
	logger, _ := zap.NewProduction()
	conn := src.GetConn(src.NewConfig(), logger)
	instance := repository.NewContactRepository(conn, logger)
	defer conn.Close()

	t.Run("Should return all contacts", func(t *testing.T) {
		ctx := context.Background()
		contacts, err := instance.GetAllContacts(ctx)

		assert.NoError(t, err)
		assert.NotEmpty(t, contacts)
		assert.Len(t, contacts, 1)
		assert.Equal(t, "fran carrero", contacts[0].FullName)
	})
}
func TestGetContactById(t *testing.T) {
	logger, _ := zap.NewProduction()
	conn := src.GetConn(src.NewConfig(), logger)
	instance := repository.NewContactRepository(conn, logger)
	defer conn.Close()

	t.Run("Should return an error querying contact by id", func(t *testing.T) {
		ctx := context.Background()
		id := 1
		contact, err := instance.GetContactByID(ctx, int64(id))

		assert.Error(t, err)
		assert.ErrorIs(t, err, customError.NewContactNotFoundError(fmt.Sprintf("contact with id %d not found", id), "contact.not_found"))
		assert.Empty(t, contact.Email)
	})
}
