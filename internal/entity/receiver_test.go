package entity

import (
	"testing"

	"github.com/felipemagrassi/pix-api/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateReceiver(t *testing.T) {
	pixKey, err := NewPixKey("felipe@email.com", "email")
	assert.Nil(t, err)

	receiver, err := NewReceiver("Felipe", "felipe@email.com", "12345678901", pixKey)
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, "Felipe")
	assert.Equal(t, receiver.Email, value_object.Email("felipe@email.com"))
	assert.Equal(t, receiver.Status, Draft)
	assert.Equal(t, receiver.Document.String(), "12345678901")
	assert.Equal(t, receiver.PixKey, pixKey)
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCanCreateReceiverWithCnpj(t *testing.T) {
	pixKey, err := NewPixKey("felipe@email.com", "email")
	assert.Nil(t, err)

	receiver, err := NewReceiver("Felipe", "felipe@email.com", "12345678901234", pixKey)
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, "Felipe")
	assert.Equal(t, receiver.Email, value_object.Email("felipe@email.com"))
	assert.Equal(t, receiver.Status, Draft)
	assert.Equal(t, receiver.Document.String(), "12345678901234")
	assert.Equal(t, receiver.PixKey, pixKey)
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCanCreateReceiverWithoutEmail(t *testing.T) {
	pixKey, err := NewPixKey("felipe@email.com", "email")
	assert.Nil(t, err)

	receiver, err := NewReceiver("Felipe", "", "12345678901", pixKey)
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, "Felipe")
	assert.Equal(t, receiver.Email, value_object.Email(""))
	assert.Equal(t, receiver.Status, Draft)
	assert.Equal(t, receiver.Document.String(), "12345678901")
	assert.Equal(t, receiver.PixKey, pixKey)
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCannotCreateReceiverWithInvalidEmail(t *testing.T) {
	pixKey, err := NewPixKey("felipe@email.com", "email")
	assert.Nil(t, err)

	receiver, err := NewReceiver("Felipe", "123", "12345678901", pixKey)
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCannotCreateReceiverWithInvalidCpf(t *testing.T) {
	pixKey, err := NewPixKey("felipe@email.com", "email")
	assert.Nil(t, err)

	receiver, err := NewReceiver("Felipe", "felipe@email", "112321", pixKey)
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCannotCreateReceiverWithLongEmail(t *testing.T) {
	pixKey, err := NewPixKey("felipe@email.com", "email")
	assert.Nil(t, err)

	email := "rgKycw8zmuIlnR6eRATh98RtPVKJDvJkW6utF584mUMLrIreqtjWVeyCoEa1Y2AtYUDpeeFJSlAuu9b8Svdg1hSKIQcZLV25miSPRR6ZifeRJahDQDkkBgfgi4CWP7LbQWxWFvitZ1r26WlFDnSggsoQKyAUXdyK7srhgvCM1abYHn3WYMJ5m3XxwunSR3n8wRJvGN0T2wYKlEpMXSCZ0RSIxZj8YAbXqkJG1T4oOfVkppenJ9U661t4qJJ@email.com"

	receiver, err := NewReceiver("Felipe", email, "12345678901", pixKey)
	assert.Error(t, err)
	assert.Nil(t, receiver)
}
