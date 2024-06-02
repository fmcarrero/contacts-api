package domain

import (
	errors "github.com/fmcarrero/contacts-api/src/contacts/domain/error"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContactSuccessfully(t *testing.T) {

	got, err := NewContact(1797303785932984320, "John Doe", "+573142805241", "mm@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", got.FullName)
	assert.Equal(t, "+573142805241", got.PhoneNumber)
	assert.Equal(t, "mm@gmail.com", got.Email)
	assert.NotZero(t, got.CreatedAt)
	assert.NotZero(t, got.UpdateAt)
	assert.Equal(t, got.CreatedAt, got.UpdateAt)
}
func TestNewContactErrorEmailFormat(t *testing.T) {

	_, err := NewContact(1797303785932984320, "John Doe", "+573142805241", "mmgmail.com")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "email validation error: invalid email format 'mmgmail.com'. Expected format: 'example@domain.com'", err.Error())
}
func TestNewContactErrorEmailIsEmpty(t *testing.T) {

	_, err := NewContact(1797303785932984320, "John Doe", "+573142805241", "")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "invalid email, email should not be empty", err.Error())
}
func TestNewContactErrorEmptyFullName(t *testing.T) {

	_, err := NewContact(1797303785932984320, "", "+573142805241", "mm@gmail.com")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "invalid fullName, fullName should not be empty", err.Error())
}
func TestNewContactErrorEmptyPhoneNumber(t *testing.T) {

	_, err := NewContact(1797303785932984320, "John Doe", "", "mm@gmail.com")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "invalid phoneNumber, phoneNumber should not be empty", err.Error())
}
func TestNewContactErrorFormatPhoneNumber(t *testing.T) {

	_, err := NewContact(1797303785932984320, "John Doe", "3142805241", "mm@gmail.com")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "phone number validation error: invalid phone number format '3142805241'. Expected format: '+573133159000' where '+' is followed by 11 to 15 digits", err.Error())
}
func TestNewContactErrorIsLessThanZero(t *testing.T) {

	_, err := NewContact(-1, "John Doe", "+573142805241", "m@gmail.com")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "invalid id -1, id should be greater than 0", err.Error())
}

func TestEditContactSuccessfully(t *testing.T) {

	got, err := NewContactEdit(1797303785932984320, "John Doe", "+573142805241", "m@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", got.FullName)
	assert.Equal(t, "+573142805241", got.PhoneNumber)
	assert.Equal(t, "m@gmail.com", got.Email)
}

func TestEditContactErrorEmailFormat(t *testing.T) {

	_, err := NewContactEdit(1797303785932984320, "John Doe", "+573142805241", "mmgmail.com")

	assert.Error(t, err)
	assert.ErrorAs(t, err, &errors.ContactValidationError{})
	assert.Equal(t, "email validation error: invalid email format 'mmgmail.com'. Expected format: 'example@domain.com'", err.Error())
}
