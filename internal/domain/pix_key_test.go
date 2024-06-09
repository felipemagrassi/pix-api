package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidPixKey(t *testing.T) {
	keyType, err := NewPixKeyType(CnpjKeyType)
	assert.Nil(t, err)

	keyValue := "41299131000107"

	key, err := NewPixKey(keyValue, keyType)
	assert.Nil(t, err)

	assert.NotNil(t, key.Id)
	assert.Equal(t, key.KeyValue, keyValue)
	assert.Equal(t, key.KeyType, keyType)
}

func TestCreateEmptyPixKeyValue(t *testing.T) {
	keyType, err := NewPixKeyType(CnpjKeyType)
	assert.Nil(t, err)

	keyValue := ""

	key, err := NewPixKey(keyValue, keyType)
	assert.Error(t, err)
	assert.Nil(t, key)
}

func TestCreateLongPixKeyValue(t *testing.T) {
	keyType, err := NewPixKeyType(CnpjKeyType)
	assert.Nil(t, err)

	keyValue := "niysxfjqaruadbglocyknpqtrejhahgkqikinaoktjlyqpzramuncszgnwdovmmrasnupwleihvsllsbmqluxybkitjmstyiyrzsbilghlfxgkqkntmbceebfpypxixklssycncasuqjy"

	key, err := NewPixKey(keyValue, keyType)
	assert.Error(t, err)
	assert.Nil(t, key)
}

func TestCreateInvalidPixKeyTypeValue(t *testing.T) {
	keyType, err := NewPixKeyType(PhoneKeyType)
	assert.Nil(t, err)

	keyValue := "41299131000107"

	key, err := NewPixKey(keyValue, keyType)
	assert.Error(t, err)
	assert.Nil(t, key)
}
