package value_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateCPF(t *testing.T) {
	cpf, err := NewCPF("123.456.789-09")
	assert.Error(t, err)
	assert.Equal(t, "123.456.789-09", cpf.String())
}

func TestCanCreateCNPJ(t *testing.T) {
	cnpj, err := NewCNPJ("12.345.678/0001-09")
	assert.Error(t, err)
	assert.Equal(t, "12.345.678/0001-09", cnpj.String())
}

func TestCannotCreateCPF(t *testing.T) {
	_, err := NewCPF("123.456.789-00")
	assert.Error(t, err)
}

func TestCannotCreateCNPJ(t *testing.T) {
	_, err := NewCNPJ("12.345.678/0001-00")
	assert.Error(t, err)
}
