package entity

import (
	"testing"

	"github.com/felipemagrassi/pix-api/internal/dto"
	"github.com/felipemagrassi/pix-api/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateReceiver(t *testing.T) {
	input := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901",
		Email:       "felipe@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	receiver, err := NewReceiver(input)
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, input.Name)
	assert.Equal(t, receiver.Email, value_object.Email(input.Email))
	assert.Equal(t, receiver.GetStatus(), Draft)
	assert.Equal(t, receiver.Document.String(), input.Document)
	assert.Equal(t, receiver.PixKey.KeyValue, input.PixKeyValue)
	assert.Equal(t, receiver.PixKey.KeyType.GetTypeName(), input.PixKeyType)
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
	assert.Equal(t, receiver.CreatedAt, receiver.UpdatedAt)
}

func TestCanCreateReceiverWithCnpj(t *testing.T) {
	input := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901234",
		Email:       "felipe@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	receiver, err := NewReceiver(input)
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, input.Name)
	assert.Equal(t, receiver.Email, value_object.Email(input.Email))
	assert.Equal(t, receiver.GetStatus(), Draft)
	assert.Equal(t, receiver.Document.String(), input.Document)
	assert.Equal(t, receiver.PixKey.KeyValue, input.PixKeyValue)
	assert.Equal(t, receiver.PixKey.KeyType.GetTypeName(), input.PixKeyType)
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCanCreateReceiverWithoutEmail(t *testing.T) {
	input := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901234",
		Email:       "",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	receiver, err := NewReceiver(input)
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, input.Name)
	assert.Equal(t, receiver.Email, value_object.Email(input.Email))
	assert.Equal(t, receiver.GetStatus(), Draft)
	assert.Equal(t, receiver.Document.String(), input.Document)
	assert.Equal(t, receiver.PixKey.KeyValue, input.PixKeyValue)
	assert.Equal(t, receiver.PixKey.KeyType.GetTypeName(), input.PixKeyType)
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCannotCreateReceiverWithInvalidEmail(t *testing.T) {
	input := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901234",
		Email:       "123",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "email",
	}

	receiver, err := NewReceiver(input)
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCannotCreateReceiverWithInvalidCpf(t *testing.T) {
	input := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "11234",
		Email:       "",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "email",
	}

	receiver, err := NewReceiver(input)
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCannotCreateReceiverWithLongEmail(t *testing.T) {
	input := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901234",
		Email:       "rgKycw8zmuIlnR6eRATh98RtPVKJDvJkW6utF584mUMLrIreqtjWVeyCoEa1Y2AtYUDpeeFJSlAuu9b8Svdg1hSKIQcZLV25miSPRR6ZifeRJahDQDkkBgfgi4CWP7LbQWxWFvitZ1r26WlFDnSggsoQKyAUXdyK7srhgvCM1abYHn3WYMJ5m3XxwunSR3n8wRJvGN0T2wYKlEpMXSCZ0RSIxZj8YAbXqkJG1T4oOfVkppenJ9U661t4qJJ@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "email",
	}

	receiver, err := NewReceiver(input)
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCanUpdateDraftedReceiver(t *testing.T) {
	createInput := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901",
		Email:       "felipe@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	createdReceiver, err := NewReceiver(createInput)
	assert.Nil(t, err)

	updateInput := dto.UpdateDraftedReceiverInput{
		Name:        "Teste",
		Document:    "12345678902",
		Email:       "teste@email.com",
		PixKeyValue: "12345678902",
		PixKeyType:  "Cpf",
	}

	err = createdReceiver.UpdateDraftedReceiver(updateInput)
	assert.Nil(t, err)

	assert.Equal(t, createdReceiver.Name, updateInput.Name)
	assert.Equal(t, createdReceiver.Email, value_object.Email(updateInput.Email))
	assert.Equal(t, createdReceiver.Document.String(), updateInput.Document)
	assert.Equal(t, createdReceiver.GetStatus(), Draft)
	assert.Equal(t, createdReceiver.PixKey.KeyValue, updateInput.PixKeyValue)
	assert.Equal(t, createdReceiver.PixKey.KeyType.GetTypeName(), updateInput.PixKeyType)
	assert.NotEqual(t, createdReceiver.CreatedAt, createdReceiver.UpdatedAt)
}

func TestCannotUpdateValidReceiver(t *testing.T) {
	createInput := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901",
		Email:       "felipe@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	createdReceiver, err := NewReceiver(createInput)
	assert.Nil(t, err)

	createdReceiver.ValidateReceiverStatus()

	updateInput := dto.UpdateDraftedReceiverInput{
		Name:        "Teste",
		Document:    "12345678902",
		Email:       "teste@email.com",
		PixKeyValue: "12345678902",
		PixKeyType:  "Cpf",
	}

	err = createdReceiver.UpdateDraftedReceiver(updateInput)
	assert.Error(t, err)

	assert.Equal(t, createdReceiver.Name, createInput.Name)
	assert.Equal(t, createdReceiver.Email, value_object.Email(createInput.Email))
	assert.Equal(t, createdReceiver.Document.String(), createInput.Document)
	assert.Equal(t, createdReceiver.GetStatus(), Valid)
	assert.Equal(t, createdReceiver.PixKey.KeyValue, createInput.PixKeyValue)
	assert.Equal(t, createdReceiver.PixKey.KeyType.GetTypeName(), createInput.PixKeyType)
	assert.Equal(t, createdReceiver.CreatedAt, createdReceiver.UpdatedAt)
}

func TestCanUpdateDraftReceiverEmail(t *testing.T) {
	createInput := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901",
		Email:       "felipe@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	createdReceiver, err := NewReceiver(createInput)
	assert.Nil(t, err)

	err = createdReceiver.UpdateEmail("felipe1@email.com")
	assert.Nil(t, err)

	assert.Equal(t, createdReceiver.Email, value_object.Email("felipe1@email.com"))
	assert.Equal(t, createdReceiver.GetStatus(), Draft)
}

func TestCanUpdateValidReceiverEmail(t *testing.T) {
	createInput := dto.CreateReceiverInput{
		Name:        "Felipe",
		Document:    "12345678901",
		Email:       "felipe@email.com",
		PixKeyValue: "felipe@email.com",
		PixKeyType:  "Email",
	}

	createdReceiver, err := NewReceiver(createInput)
	assert.Nil(t, err)

	createdReceiver.ValidateReceiverStatus()

	err = createdReceiver.UpdateEmail("felipe1@email.com")
	assert.Nil(t, err)

	assert.Equal(t, createdReceiver.Email, value_object.Email("felipe1@email.com"))
	assert.Equal(t, createdReceiver.GetStatus(), Valid)
}
