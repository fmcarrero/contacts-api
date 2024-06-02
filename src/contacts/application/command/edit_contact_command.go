package command

import (
	"context"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
)

type EditContactCommand struct {
	ContactRepository repository.ContactRepository
}

func NewEditContact(contactRepository repository.ContactRepository) EditContactCommand {
	return EditContactCommand{
		ContactRepository: contactRepository,
	}
}

type EditContactRequest struct {
	ID          int64  `param:"id" json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (c EditContactCommand) Execute(ctx context.Context, request EditContactRequest) (domain.Contact, error) {
	contact, err := domain.NewContactEdit(request.ID, request.FullName, request.PhoneNumber, request.Email)
	if err != nil {
		return domain.Contact{}, err
	}
	foundedContact, err := c.ContactRepository.GetContactByID(ctx, contact.ID)
	if err != nil {
		return domain.Contact{}, err
	}
	contact.CreatedAt = foundedContact.CreatedAt
	if err = c.ContactRepository.EditContact(ctx, contact); err != nil {
		return domain.Contact{}, err
	}
	return contact, nil
}
