package command

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
)

type GetAllContactsCommand struct {
	ContactRepository repository.ContactRepository
}

func NewGetAllContacts(contactRepository repository.ContactRepository) GetAllContactsCommand {
	return GetAllContactsCommand{
		ContactRepository: contactRepository,
	}
}

func (c GetAllContactsCommand) Execute(ctx context.Context) ([]domain.Contact, error) {
	return c.ContactRepository.GetAllContacts(ctx)
}
