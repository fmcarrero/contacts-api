package command

import (
	"context"
	"github.com/fmcarrero/contacts-api/pkg/snowflake"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"github.com/fmcarrero/contacts-api/src/contacts/domain/repository"
)

type AddContactCommand struct {
	ContactRepository repository.ContactRepository
}

func NewAddContact(contactRepository repository.ContactRepository) AddContactCommand {
	return AddContactCommand{
		ContactRepository: contactRepository,
	}
}

type AddContactRequest struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (c AddContactCommand) Executes(ctx context.Context, request AddContactRequest) (domain.Contact, error) {
	contact, err := domain.NewContact(snowflake.NewSnowflakeID(), request.FullName, request.PhoneNumber, request.Email)
	if err != nil {
		return domain.Contact{}, err
	}
	if err = c.ContactRepository.AddContact(ctx, contact); err != nil {
		return domain.Contact{}, err
	}
	return contact, nil
}
