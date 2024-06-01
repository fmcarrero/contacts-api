package repository

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
)

type ContactRepository interface {
	GetAllContacts(ctx context.Context) ([]domain.Contact, error)
}
