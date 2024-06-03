package builder

import (
	"github.com/fmcarrero/contacts-api/src/contacts/application/command"
	"github.com/go-faker/faker/v4"
	"math/rand/v2"
	"strconv"
)

type AddContactRequestBuilder struct {
	fullName    string
	email       string
	phoneNumber string
}

func NewAddContactRequestBuilder() *AddContactRequestBuilder {
	minValue := int64(1000000000)
	maxValue := int64(9999999999)
	randomNumber := rand.Int64N(maxValue-minValue+1) + minValue

	return &AddContactRequestBuilder{
		fullName:    faker.FirstName() + " " + faker.LastName(),
		email:       faker.Email(),
		phoneNumber: "+57" + strconv.FormatInt(randomNumber, 10),
	}
}

func (b *AddContactRequestBuilder) FullName(fullName string) *AddContactRequestBuilder {
	b.fullName = fullName
	return b
}

func (b *AddContactRequestBuilder) Email(email string) *AddContactRequestBuilder {
	b.email = email
	return b
}

func (b *AddContactRequestBuilder) PhoneNumber(phoneNumber string) *AddContactRequestBuilder {
	b.phoneNumber = phoneNumber
	return b
}

func (b *AddContactRequestBuilder) Build() *command.AddContactRequest {
	return &command.AddContactRequest{
		FullName:    b.fullName,
		Email:       b.email,
		PhoneNumber: b.phoneNumber,
	}
}
