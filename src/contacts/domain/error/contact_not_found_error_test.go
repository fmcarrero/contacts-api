package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContactNotFoundError(t *testing.T) {

	err := NewContactNotFoundError("message", "type")

	assert.Error(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t, "type", err.Type)
}
