package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidPixKey(t *testing.T) {
	keyValue := "41299131000107"
	keyType := "cnpj"

	key, err := NewPixKey(keyValue, keyType)
	assert.Nil(t, err)

	assert.Equal(t, key.KeyValue, keyValue)

	parsedKeyType, ok := ParsePixKeyType(keyType)
	assert.True(t, ok)
	associatedKeyType, err := NewPixKeyType(parsedKeyType)

	assert.Nil(t, err)
	assert.Equal(t, key.KeyType, associatedKeyType)
}

func TestCreateEmptyPixKeyValue(t *testing.T) {
	keyValue := ""
	keyType := "cnpj"

	key, err := NewPixKey(keyValue, keyType)
	assert.Error(t, err)
	assert.Nil(t, key)
}

func TestCreateLongPixKeyValue(t *testing.T) {
	keyValue := "niysxfjqaruadbglocyknpqtrejhahgkqikinaoktjlyqpzramuncszgnwdovmmrasnupwleihvsllsbmqluxybkitjmstyiyrzsbilghlfxgkqkntmbceebfpypxixklssycncasuqjy"
	keyType := "cnpj"

	key, err := NewPixKey(keyValue, keyType)
	assert.Error(t, err)
	assert.Nil(t, key)
}

func TestCreateInvalidPixKeyTypeValue(t *testing.T) {
	keyValue := "41299131000107"
	keyType := "invalid"

	key, err := NewPixKey(keyValue, keyType)
	assert.Error(t, err)
	assert.Nil(t, key)
}

func TestValidCnpjKeyTypeTest(t *testing.T) {
	keyType := CnpjKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"41.299.131/0001-07",
		"41299131000107",
	}

	for _, cnpj := range expected {

		err = generatedKeyType.ValidateKeyType(cnpj)
		assert.Nil(t, err)
	}
}

func TestInvalidCnpjKeyTypeTest(t *testing.T) {
	keyType := CnpjKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"abc",
		"123",
		"498.777.520-42",
		"49877752042",
	}
	for _, cnpj := range expected {

		err = generatedKeyType.ValidateKeyType(cnpj)
		assert.Error(t, err)
	}
}

func TestValidCpfKeyTypeTest(t *testing.T) {
	keyType := CpfKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"498.777.520-42",
		"49877752042",
	}

	for _, cpf := range expected {

		err = generatedKeyType.ValidateKeyType(cpf)
		assert.Nil(t, err)
	}
}

func TestInvalidCpfKeyTypeTest(t *testing.T) {
	keyType := CpfKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"abc",
		"123",
		"41.299.131/0001-07",
		"41299131000107",
	}
	for _, cpf := range expected {

		err = generatedKeyType.ValidateKeyType(cpf)
		assert.Error(t, err)
	}
}

func TestValidEmailKeyTypeTest(t *testing.T) {
	keyType := EmailKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"govrada@peakvisionhdtv.com",
		"govrada@gmail.com",
		"govrada@gmail",
	}

	for _, email := range expected {

		err = generatedKeyType.ValidateKeyType(email)
		assert.Nil(t, err)
	}
}

func TestInvalidEmailKeyTypeTest(t *testing.T) {
	keyType := EmailKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"abc",
		"123",
		"41.299.131/0001-07",
		"41299131000107",
	}
	for _, email := range expected {

		err = generatedKeyType.ValidateKeyType(email)
		assert.Error(t, err)
	}
}

func TestValidPhoneKeyTypeTest(t *testing.T) {
	keyType := PhoneKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"5511999999999",
		"5521999999999",
	}

	for _, phone := range expected {
		testname := fmt.Sprintf("%s", phone)
		t.Run(testname, func(t *testing.T) {
			err = generatedKeyType.ValidateKeyType(phone)
			assert.Nil(t, err)
		})

	}
}

func TestInvalidPhoneKeyTypeTest(t *testing.T) {
	keyType := PhoneKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"abc",
		"123",
		"41.299.131/0001-07",
		"41299131000107",
		"0111999999999",
		"55123999999999",
		"(55) 123999999999",
		"(55) 21 999999999",
	}
	for _, phone := range expected {

		err = generatedKeyType.ValidateKeyType(phone)
		assert.Error(t, err)
	}
}

func TestValidRandomKeyTypeTest(t *testing.T) {
	keyType := RandomKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"3f2504e0-4f89-11d3-9a0c-0305e82c3301",
		"7c7a2ba0-3fda-4f76-8c44-df1f8c1289ba",
	}

	for _, random := range expected {

		err = generatedKeyType.ValidateKeyType(random)
		assert.Nil(t, err)
	}
}

func TestInvalidRandomKeyTypeTest(t *testing.T) {
	keyType := RandomKeyType
	generatedKeyType, err := NewPixKeyType(keyType)
	assert.Nil(t, err)
	assert.NotNil(t, generatedKeyType)

	expected := []string{
		"abc",
		"123",
		"41.299.131/0001-07",
		"41299131000107",
		"3f2504e0-4f89-11d3-9a0c-0305e82c331",
	}
	for _, random := range expected {

		err = generatedKeyType.ValidateKeyType(random)
		assert.Error(t, err)
	}
}
