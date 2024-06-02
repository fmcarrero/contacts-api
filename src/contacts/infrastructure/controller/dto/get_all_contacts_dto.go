package dto

import (
	"github.com/fmcarrero/contacts-api/src/contacts/domain"
	"time"
)

type GetAllContactsDTO struct {
	Data  []contact `json:"data"`
	Total int       `json:"total"`
}
type contact struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Indicator   string    `json:"indicator"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func ToContactsDTO(contacts []domain.Contact) GetAllContactsDTO {
	contactsDTO := make([]contact, 0)
	for _, c := range contacts {
		contactsDTO = append(contactsDTO, ToContactDTO(c))
	}
	return GetAllContactsDTO{
		Data:  contactsDTO,
		Total: len(contactsDTO),
	}
}
func ToContactDTO(c domain.Contact) contact {
	return contact{
		ID:          c.ID,
		FullName:    c.FullName,
		PhoneNumber: c.PhoneNumber,
		Indicator:   c.PhoneNumber[0:3],
		Email:       c.Email,
		CreatedAt:   c.CreatedAt,
		UpdateAt:    c.UpdateAt,
	}
}
