package integration

import (
	"context"
	"fmt"
	"github.com/fmcarrero/contacts-api/tests/builder"
	"testing"
	"time"

	"github.com/fmcarrero/contacts-api/src"
	customError "github.com/fmcarrero/contacts-api/src/contacts/domain/error"
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

	t.Run("Should return a contact by id", func(t *testing.T) {
		//prepare
		contact, _ := builder.NewContactBuilder().Build()
		assert.NoError(t, instance.AddContact(context.Background(), contact))
		defer instance.RemoveContact(context.Background(), contact.ID)

		foundedContact, err := instance.GetContactByID(context.Background(), contact.ID)

		assert.NoError(t, err)
		assert.EqualValues(t, contact.ID, foundedContact.ID)
		assert.EqualValues(t, contact.FullName, foundedContact.FullName)
		assert.EqualValues(t, contact.PhoneNumber, foundedContact.PhoneNumber)
		assert.EqualValues(t, contact.Email, foundedContact.Email)
		assert.EqualValues(t, contact.CreatedAt.Truncate(time.Second).Local(), foundedContact.CreatedAt.Truncate(time.Second).Local())
		assert.EqualValues(t, contact.UpdateAt.Truncate(time.Second).Local(), foundedContact.UpdateAt.Truncate(time.Second).Local())
	})
}
func TestAddContact(t *testing.T) {
	logger, _ := zap.NewProduction()
	conn := src.GetConn(src.NewConfig(), logger)
	instance := repository.NewContactRepository(conn, logger)
	defer conn.Close()

	t.Run("Should return an error adding contact", func(t *testing.T) {
		ctx := context.Background()
		contact, _ := builder.NewContactBuilder().WithID(-2).Build()
		err := instance.AddContact(ctx, contact)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "violates check constraint")
	})

	t.Run("Should add a contact", func(t *testing.T) {
		ctx := context.Background()
		contact, _ := builder.NewContactBuilder().Build()
		defer instance.RemoveContact(ctx, contact.ID)

		err := instance.AddContact(ctx, contact)
		foundedContact, errSearch := instance.GetContactByID(ctx, contact.ID)

		assert.NoError(t, err)
		assert.NoError(t, errSearch)
		assert.EqualValues(t, contact.ID, foundedContact.ID)
		assert.EqualValues(t, contact.FullName, foundedContact.FullName)
		assert.EqualValues(t, contact.PhoneNumber, foundedContact.PhoneNumber)
		assert.EqualValues(t, contact.Email, foundedContact.Email)
		assert.EqualValues(t, contact.CreatedAt.Truncate(time.Second).Local(), foundedContact.CreatedAt.Truncate(time.Second).Local())
		assert.EqualValues(t, contact.UpdateAt.Truncate(time.Second).Local(), foundedContact.UpdateAt.Truncate(time.Second).Local())
	})
}
func TestRemoveContact(t *testing.T) {
	logger, _ := zap.NewProduction()
	conn := src.GetConn(src.NewConfig(), logger)
	instance := repository.NewContactRepository(conn, logger)
	defer conn.Close()

	t.Run("Should remove a contact", func(t *testing.T) {
		ctx := context.Background()
		contact, _ := builder.NewContactBuilder().Build()
		assert.NoError(t, instance.AddContact(ctx, contact))

		err := instance.RemoveContact(ctx, contact.ID)
		_, errSearch := instance.GetContactByID(ctx, contact.ID)

		assert.NoError(t, err)
		assert.Error(t, errSearch)
	})
}

func TestEditContact(t *testing.T) {
	logger, _ := zap.NewProduction()
	conn := src.GetConn(src.NewConfig(), logger)
	instance := repository.NewContactRepository(conn, logger)
	defer conn.Close()

	t.Run("Should return an error editing contact", func(t *testing.T) {
		ctx := context.Background()
		contact, _ := builder.NewContactBuilder().Build()
		assert.NoError(t, instance.AddContact(ctx, contact))
		contact.FullName = ""
		defer instance.RemoveContact(ctx, contact.ID)

		err := instance.EditContact(ctx, contact)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "contacts_full_name_check")
	})

	t.Run("Should edit a contact", func(t *testing.T) {
		ctx := context.Background()
		contact, _ := builder.NewContactBuilder().Build()
		assert.NoError(t, instance.AddContact(ctx, contact))
		defer instance.RemoveContact(ctx, contact.ID)

		contact.FullName = "new name"
		err := instance.EditContact(ctx, contact)
		foundedContact, errSearch := instance.GetContactByID(ctx, contact.ID)

		assert.NoError(t, err)
		assert.NoError(t, errSearch)
		assert.EqualValues(t, contact.ID, foundedContact.ID)
		assert.EqualValues(t, contact.FullName, foundedContact.FullName)
		assert.EqualValues(t, contact.PhoneNumber, foundedContact.PhoneNumber)
		assert.EqualValues(t, contact.Email, foundedContact.Email)
		assert.EqualValues(t, contact.CreatedAt.Truncate(time.Second).Local(), foundedContact.CreatedAt.Truncate(time.Second).Local())
		assert.EqualValues(t, contact.UpdateAt.Truncate(time.Second).Local(), foundedContact.UpdateAt.Truncate(time.Second).Local())
	})
}
