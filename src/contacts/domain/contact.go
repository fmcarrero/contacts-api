package domain

import (
	"fmt"
	errors "github.com/fmcarrero/contacts-api/src/contacts/domain/error"
	"regexp"
	"time"
)

type Contact struct {
	ID          int64
	FullName    string
	PhoneNumber string
	Email       string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func NewContact(id int64, fullName, phoneNumber, email string) (Contact, error) {

	now := time.Now().UTC()
	contact := Contact{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Email:       email,
		CreatedAt:   now,
		UpdateAt:    now,
	}
	if err := contact.Validate(); err != nil {
		return Contact{}, err
	}
	return contact, nil
}
func NewContactEdit(id int64, fullName, phoneNumber, email string) (Contact, error) {

	contact := Contact{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Email:       email,
		UpdateAt:    time.Now().UTC(),
	}

	if err := contact.Validate(); err != nil {
		return Contact{}, err
	}

	return contact, nil
}

// Validate validates the email and phone number formats
func (c *Contact) Validate() error {
	if c.ID <= 0 {
		return fmt.Errorf("invalid id %d, id should be greather than 0", c.ID)
	}
	if c.FullName == "" {
		return errors.NewContactValidationError("invalid fullName, fullName should not be empty", "contact.validation.full_name_empty.error")
	}
	if c.PhoneNumber == "" {
		return errors.NewContactValidationError("invalid phoneNumber, phoneNumber should not be empty", "contact.validation.phone_number_empty.error")
	}
	if c.Email == "" {
		return errors.NewContactValidationError("invalid email, email should not be empty", "contact.validation.email_empty.error")
	}
	if err := validateEmail(c.Email); err != nil {
		return errors.NewContactValidationError(fmt.Errorf("email validation error: %w", err).Error(), "contact.validation.email_format.error")
	}
	if err := validatePhoneNumber(c.PhoneNumber); err != nil {
		return errors.NewContactValidationError(fmt.Errorf("phone number validation error: %w", err).Error(), "contact.validation.phone_number_format.error")
	}
	return nil
}

// validateEmail validates the email format
func validateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return fmt.Errorf("invalid email format '%s'. Expected format: 'example@domain.com'", email)
	}
	return nil
}

// validatePhoneNumber validates the phone number format
func validatePhoneNumber(phone string) error {
	re := regexp.MustCompile(`^\+\d{11,15}$`)
	if !re.MatchString(phone) {
		return fmt.Errorf("invalid phone number format '%s'. Expected format: '+573133159000' where '+' is followed by 11 to 15 digits", phone)
	}
	return nil
}
