package repository

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
)

type ContactRepository interface {
	GetAllContacts(ctx context.Context) ([]domain.Contact, error)
	EditContact(ctx context.Context, contact domain.Contact) error
	AddContact(ctx context.Context, contact domain.Contact) error
	GetContactByID(ctx context.Context, id int64) (domain.Contact, error)
}
