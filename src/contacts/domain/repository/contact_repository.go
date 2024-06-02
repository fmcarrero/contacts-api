package repository

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
)

type ContactRepository interface {
	// GetAllContacts returns all contacts
	GetAllContacts(ctx context.Context) ([]domain.Contact, error)
	// EditContact updates a contact
	EditContact(ctx context.Context, contact domain.Contact) error
	// RemoveContact removes a contact
	RemoveContact(ctx context.Context, id int64) error
	// AddContact adds a contact
	AddContact(ctx context.Context, contact domain.Contact) error
	// GetContactByID returns a contact by id
	GetContactByID(ctx context.Context, id int64) (domain.Contact, error)
}
