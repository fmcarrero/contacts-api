package command

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
)

type RemoveContactCommand struct {
	ContactRepository repository.ContactRepository
}

func NewRemoveContact(contactRepository repository.ContactRepository) RemoveContactCommand {
	return RemoveContactCommand{
		ContactRepository: contactRepository,
	}
}

func (c RemoveContactCommand) Execute(ctx context.Context, id int64) (domain.Contact, error) {
	contact, err := c.ContactRepository.GetContactByID(ctx, id)
	if err != nil {
		return domain.Contact{}, err
	}
	if err = c.ContactRepository.RemoveContact(ctx, id); err != nil {
		return domain.Contact{}, err
	}
	return contact, nil
}
