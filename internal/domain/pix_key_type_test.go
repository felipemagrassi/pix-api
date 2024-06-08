package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCnpjKeyTypeTest(t *testing.T) {
	cnpjs := []string{
		"41.299.131/0001-07",
		"41299131000107",
	}

	for _, cnpj := range cnpjs {
		err := NewCnpjPixKeyType(cnpj).ValidateKey()
		assert.Nil(t, err)
	}
}

func TestInvalidCnpjKeyTypeTest(t *testing.T) {
}

func TestValidCpjKeyTypeTest(t *testing.T) {
}

func TestInvalidCpjKeyTypeTest(t *testing.T) {
}

func TestValidEmailKeyTypeTest(t *testing.T) {
}

func TestInvalidEmailKeyTypeTest(t *testing.T) {
}

func TestValidPhoneKeyTypeTest(t *testing.T) {
}

func TestInvalidPhoneKeyTypeTest(t *testing.T) {
}

func TestValidRandomKeyTypeTest(t *testing.T) {
}

func TestInvalidRandomKeyTypeTest(t *testing.T) {
}
