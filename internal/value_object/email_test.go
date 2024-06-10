package value_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateEmail(t *testing.T) {
	email, err := NewEmail("felipe@email")
	assert.Nil(t, err)
	assert.NotNil(t, email)
}

func TestEmptyEmail(t *testing.T) {
	email, err := NewEmail("")
	assert.Error(t, err)
	assert.Empty(t, email)
}
