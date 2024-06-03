package builders

import (
	"github.com/fmcarrero/contacts-api/pkg/snowflake"
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"time"
)

type ContactBuilder struct {
	contact domain.Contact
}

func NewContactBuilder() *ContactBuilder {
	now := time.Now().UTC()
	return &ContactBuilder{
		contact: domain.Contact{
			ID:          snowflake.NewSnowflakeID(),
			FullName:    "don quijote de la mancha",
			PhoneNumber: "+583124587845",
			Email:       "default@example.com",
			CreatedAt:   now,
			UpdateAt:    now,
		},
	}
}

func (b *ContactBuilder) WithID(id int64) *ContactBuilder {
	b.contact.ID = id
	return b
}

func (b *ContactBuilder) WithFullName(fullName string) *ContactBuilder {
	b.contact.FullName = fullName
	return b
}

func (b *ContactBuilder) WithPhoneNumber(phoneNumber string) *ContactBuilder {
	b.contact.PhoneNumber = phoneNumber
	return b
}

func (b *ContactBuilder) WithEmail(email string) *ContactBuilder {
	b.contact.Email = email
	return b
}

func (b *ContactBuilder) Build() (domain.Contact, error) {
	if err := b.contact.Validate(); err != nil {
		return domain.Contact{}, err
	}
	return b.contact, nil
}
