package domain

import (
	"fmt"
	"time"
)

type Contact struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func NewContact(id int64, fullName, phoneNumber, email string) (Contact, error) {
	if id <= 0 {
		return Contact{}, fmt.Errorf("invalid id %d, id should be greather than 0", id)
	}
	if fullName == "" {
		return Contact{}, fmt.Errorf("invalid fullName %s, fullName should not be empty", fullName)
	}
	if phoneNumber == "" {
		return Contact{}, fmt.Errorf("invalid phoneNumber %s, phoneNumber should not be empty", phoneNumber)
	}
	if email == "" {
		return Contact{}, fmt.Errorf("invalid email %s, email should not be empty", email)
	}
	now := time.Now().UTC()
	return Contact{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Email:       email,
		CreatedAt:   now,
		UpdateAt:    now,
	}, nil
}
func NewContactEdit(id int64, fullName, phoneNumber, email string) (Contact, error) {
	if id <= 0 {
		return Contact{}, fmt.Errorf("invalid id %d, id should be greather than 0", id)
	}
	if fullName == "" {
		return Contact{}, fmt.Errorf("invalid fullName %s, fullName should not be empty", fullName)
	}
	if phoneNumber == "" {
		return Contact{}, fmt.Errorf("invalid phoneNumber %s, phoneNumber should not be empty", phoneNumber)
	}
	if email == "" {
		return Contact{}, fmt.Errorf("invalid email %s, email should not be empty", email)
	}

	return Contact{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Email:       email,
		UpdateAt:    time.Now().UTC(),
	}, nil
}
