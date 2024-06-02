package query

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
)

type GetAllContactsQuery struct {
	ContactRepository repository.ContactRepository
}

func NewGetAllContacts(contactRepository repository.ContactRepository) GetAllContactsQuery {
	return GetAllContactsQuery{
		ContactRepository: contactRepository,
	}
}

func (c GetAllContactsQuery) Execute(ctx context.Context) ([]domain.Contact, error) {
	return c.ContactRepository.GetAllContacts(ctx)
}
