package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContactValidationError(t *testing.T) {
	err := NewContactValidationError("message", "type")

	assert.Error(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t, "type", err.Type)
}
