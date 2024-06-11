package entity

import (
	"testing"

	"github.com/felipemagrassi/pix-api/internal/value_object"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateReceiver(t *testing.T) {
	input := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "felipe@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	receiver, err := NewReceiver(input["document"], input["pixKeyValue"], input["pixKeyType"], input["name"], input["email"])
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, input["name"])
	assert.Equal(t, receiver.Email, value_object.Email(input["email"]))
	assert.Equal(t, receiver.GetStatus(), Draft)
	assert.Equal(t, receiver.Document.String(), input["document"])
	assert.Equal(t, receiver.PixKey.KeyValue, input["pixKeyValue"])
	assert.Equal(t, receiver.PixKey.KeyType.GetTypeName(), input["pixKeyType"])
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
	assert.Equal(t, receiver.CreatedAt, receiver.UpdatedAt)
}

func TestCanCreateReceiverWithCnpj(t *testing.T) {
	input := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901234",
		"email":       "felipe@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	receiver, err := NewReceiver(input["document"], input["pixKeyValue"], input["pixKeyType"], input["name"], input["email"])
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, input["name"])
	assert.Equal(t, receiver.Email, value_object.Email(input["email"]))
	assert.Equal(t, receiver.GetStatus(), Draft)
	assert.Equal(t, receiver.Document.String(), input["document"])
	assert.Equal(t, receiver.PixKey.KeyValue, input["pixKeyValue"])
	assert.Equal(t, receiver.PixKey.KeyType.GetTypeName(), input["pixKeyType"])
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCanCreateReceiverWithoutEmail(t *testing.T) {
	input := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	receiver, err := NewReceiver(input["document"], input["pixKeyValue"], input["pixKeyType"], input["name"], input["email"])
	assert.Nil(t, err)

	assert.NotEmpty(t, receiver.Id)
	assert.Equal(t, receiver.Name, input["name"])
	assert.Equal(t, receiver.Email, value_object.Email(input["email"]))
	assert.Equal(t, receiver.GetStatus(), Draft)
	assert.Equal(t, receiver.Document.String(), input["document"])
	assert.Equal(t, receiver.PixKey.KeyValue, input["pixKeyValue"])
	assert.Equal(t, receiver.PixKey.KeyType.GetTypeName(), input["pixKeyType"])
	assert.Empty(t, receiver.Bank)
	assert.Empty(t, receiver.Office)
	assert.Empty(t, receiver.AccountNumber)
	assert.NotEmpty(t, receiver.CreatedAt)
	assert.NotEmpty(t, receiver.UpdatedAt)
}

func TestCannotCreateReceiverWithInvalidEmail(t *testing.T) {
	input := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "123",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	receiver, err := NewReceiver(input["document"], input["pixKeyValue"], input["pixKeyType"], input["name"], input["email"])

	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCannotCreateReceiverWithInvalidCpf(t *testing.T) {
	input := map[string]string{
		"name":        "Felipe",
		"document":    "1234",
		"email":       "123",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	receiver, err := NewReceiver(input["document"], input["pixKeyValue"], input["pixKeyType"], input["name"], input["email"])
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCannotCreateReceiverWithLongEmail(t *testing.T) {
	input := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "rgKycw8zmuIlnR6eRATh98RtPVKJDvJkW6utF584mUMLrIreqtjWVeyCoEa1Y2AtYUDpeeFJSlAuu9b8Svdg1hSKIQcZLV25miSPRR6ZifeRJahDQDkkBgfgi4CWP7LbQWxWFvitZ1r26WlFDnSggsoQKyAUXdyK7srhgvCM1abYHn3WYMJ5m3XxwunSR3n8wRJvGN0T2wYKlEpMXSCZ0RSIxZj8YAbXqkJG1T4oOfVkppenJ9U661t4qJJ@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	receiver, err := NewReceiver(input["document"], input["pixKeyValue"], input["pixKeyType"], input["name"], input["email"])
	assert.Error(t, err)
	assert.Nil(t, receiver)
}

func TestCanUpdateDraftedReceiver(t *testing.T) {
	createInput := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "felipe@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	createdReceiver, err := NewReceiver(createInput["document"], createInput["pixKeyValue"], createInput["pixKeyType"], createInput["name"], createInput["email"])
	assert.Nil(t, err)

	updateInput := map[string]string{
		"name":        "Teste",
		"document":    "12345678902",
		"email":       "test@email.com",
		"pixKeyValue": "12345678902",
		"pixKeyType":  "Cpf",
	}

	err = createdReceiver.UpdateReceiver(updateInput["document"], updateInput["pixKeyValue"], updateInput["pixKeyType"], updateInput["name"], updateInput["email"])
	assert.Nil(t, err)

	assert.Equal(t, createdReceiver.Name, updateInput["name"])
	assert.Equal(t, createdReceiver.Email, value_object.Email(updateInput["email"]))
	assert.Equal(t, createdReceiver.Document.String(), updateInput["document"])
	assert.Equal(t, createdReceiver.GetStatus(), Draft)
	assert.Equal(t, createdReceiver.PixKey.KeyValue, updateInput["pixKeyValue"])
	assert.Equal(t, createdReceiver.PixKey.KeyType.GetTypeName(), updateInput["pixKeyType"])
	assert.NotEqual(t, createdReceiver.CreatedAt, createdReceiver.UpdatedAt)
}

func TestOnlyUpdatesEmailOnValidReceiver(t *testing.T) {
	createInput := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "felipe@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	createdReceiver, err := NewReceiver(createInput["document"], createInput["pixKeyValue"], createInput["pixKeyType"], createInput["name"], createInput["email"])
	assert.Nil(t, err)

	createdReceiver.ValidateReceiverStatus()

	updateInput := map[string]string{
		"name":        "Teste",
		"document":    "12345678902",
		"email":       "test@email.com",
		"pixKeyValue": "12345678902",
		"pixKeyType":  "Cpf",
	}

	err = createdReceiver.UpdateReceiver(updateInput["document"], updateInput["pixKeyValue"], updateInput["pixKeyType"], updateInput["name"], updateInput["email"])
	assert.Error(t, err)

	assert.Equal(t, createdReceiver.Name, createInput["name"])
	assert.Equal(t, createdReceiver.Email, value_object.Email(updateInput["email"]))
	assert.Equal(t, createdReceiver.Document.String(), createInput["document"])
	assert.Equal(t, createdReceiver.GetStatus(), Valid)
	assert.Equal(t, createdReceiver.PixKey.KeyValue, createInput["pixKeyValue"])
	assert.Equal(t, createdReceiver.PixKey.KeyType.GetTypeName(), createInput["pixKeyType"])
	assert.NotEqual(t, createdReceiver.CreatedAt, createdReceiver.UpdatedAt)
}

func TestCanUpdateDraftReceiverEmail(t *testing.T) {
	createInput := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "felipe@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	createdReceiver, err := NewReceiver(createInput["document"], createInput["pixKeyValue"], createInput["pixKeyType"], createInput["name"], createInput["email"])
	assert.Nil(t, err)

	newEmail := "felipe1@email.com"
	err = createdReceiver.UpdateReceiver("", "", "", "", newEmail)
	assert.Nil(t, err)

	assert.Equal(t, createdReceiver.Email, value_object.Email(newEmail))
	assert.Equal(t, createdReceiver.GetStatus(), Draft)
}

func TestCanUpdateValidReceiverEmail(t *testing.T) {
	createInput := map[string]string{
		"name":        "Felipe",
		"document":    "12345678901",
		"email":       "felipe@email.com",
		"pixKeyValue": "felipe@email.com",
		"pixKeyType":  "Email",
	}

	createdReceiver, err := NewReceiver(createInput["document"], createInput["pixKeyValue"], createInput["pixKeyType"], createInput["name"], createInput["email"])
	assert.Nil(t, err)

	createdReceiver.ValidateReceiverStatus()

	newEmail := "felipe1@email.com"
	err = createdReceiver.UpdateReceiver("", "", "", "", newEmail)
	assert.Nil(t, err)

	assert.Equal(t, createdReceiver.Email, value_object.Email(newEmail))
	assert.Equal(t, createdReceiver.GetStatus(), Valid)
}
