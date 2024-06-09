package domain

import (
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/pkg/domain"
)

type PixKey struct {
	Id       domain.ID
	KeyValue string
	KeyType  PixKeyTypeInterface
}

func NewPixKey(
	keyValue string,
	keyType PixKeyTypeInterface,
) (*PixKey, *internal_error.InternalError) {
	pixKey := &PixKey{
		Id:       domain.NewID(),
		KeyValue: keyValue,
		KeyType:  keyType,
	}

	if err := pixKey.Validate(); err != nil {
		return nil, internal_error.NewBadRequestError("Invalid pix key")
	}

	return pixKey, nil
}

func (pk *PixKey) Validate() *internal_error.InternalError {
	if pk.KeyValue == "" {
		return internal_error.NewBadRequestError("Invalid pix key")
	}

	if pk.KeyType.GetType() == "" {
		return internal_error.NewBadRequestError("Invalid pix key")
	}

	if len(pk.KeyValue) > 140 {
		return internal_error.NewBadRequestError("Invalid pix key")
	}

	if pk.KeyType.ValidateKeyType(pk.KeyValue) != nil {
		return internal_error.NewBadRequestError("Invalid pix key")
	}

	return nil
}
