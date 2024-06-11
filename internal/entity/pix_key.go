package entity

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/internal/value_object"
)

type PixKey struct {
	KeyValue string
	KeyType  PixKeyTypeInterface
}

type (
	CnpjPixKeyType   struct{ KeyType PixKeyType }
	CpfPixKeyType    struct{ KeyType PixKeyType }
	EmailPixKeyType  struct{ KeyType PixKeyType }
	PhonePixKeyType  struct{ KeyType PixKeyType }
	RandomPixKeyType struct{ KeyType PixKeyType }
)

const (
	PhoneKeyPattern  = `^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`
	RandomKeyPattern = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
)

type PixKeyType int

const (
	_ PixKeyType = iota
	CnpjKeyType
	CpfKeyType
	EmailKeyType
	PhoneKeyType
	RandomKeyType
)

var pixKeyTypeMap = map[string]PixKeyType{
	"cnpj":   CnpjKeyType,
	"cpf":    CpfKeyType,
	"email":  EmailKeyType,
	"phone":  PhoneKeyType,
	"random": RandomKeyType,
}

func NewPixKey(
	keyValue string,
	keyType string,
) (*PixKey, *internal_error.InternalError) {
	parsedKeyType, ok := ParsePixKeyType(keyType)
	if !ok {
		return nil, internal_error.NewBadRequestError("Invalid pix key type")
	}

	newPixKeyType, err := NewPixKeyType(parsedKeyType)
	if err != nil {
		return nil, err
	}

	newPixKey := &PixKey{
		KeyValue: keyValue,
		KeyType:  newPixKeyType,
	}

	if err := newPixKey.Validate(); err != nil {
		return nil, err
	}

	return newPixKey, nil
}

func (pk *PixKey) Validate() *internal_error.InternalError {
	if pk.KeyValue == "" {
		return internal_error.NewBadRequestError("Invalid pix key")
	}

	if pk.KeyType.GetTypeName() == "" {
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

func ParsePixKeyType(pixKeyTypeStr string) (PixKeyType, bool) {
	lowerPixKeyTypeStr := strings.ToLower(pixKeyTypeStr)
	c, ok := pixKeyTypeMap[lowerPixKeyTypeStr]
	return c, ok
}

func (pkt PixKeyType) String() string {
	return []string{"Cnpj", "Cpf", "Email", "Phone", "Random"}[pkt-1]
}

type PixKeyTypeInterface interface {
	ValidateKeyType(key string) *internal_error.InternalError
	GetTypeName() string
	Value() PixKeyType
}

func NewPixKeyType(keyType PixKeyType) (PixKeyTypeInterface, *internal_error.InternalError) {
	switch keyType {
	case CnpjKeyType:
		return &CnpjPixKeyType{KeyType: keyType}, nil
	case CpfKeyType:
		return &CpfPixKeyType{KeyType: keyType}, nil
	case EmailKeyType:
		return &EmailPixKeyType{KeyType: keyType}, nil
	case PhoneKeyType:
		return &PhonePixKeyType{KeyType: keyType}, nil
	case RandomKeyType:
		return &RandomPixKeyType{KeyType: keyType}, nil
	default:
		return nil, internal_error.NewBadRequestError("Invalid Key Type")
	}
}

func (kt *CnpjPixKeyType) GetTypeName() string {
	return kt.KeyType.String()
}

func (kt *CpfPixKeyType) GetTypeName() string {
	return kt.KeyType.String()
}

func (kt *EmailPixKeyType) GetTypeName() string {
	return kt.KeyType.String()
}

func (kt *PhonePixKeyType) GetTypeName() string {
	return kt.KeyType.String()
}

func (kt *RandomPixKeyType) GetTypeName() string {
	return kt.KeyType.String()
}

func (kt *CnpjPixKeyType) Value() PixKeyType {
	return PixKeyType(CnpjKeyType)
}

func (kt *CpfPixKeyType) Value() PixKeyType {
	return PixKeyType(CpfKeyType)
}

func (kt *EmailPixKeyType) Value() PixKeyType {
	return PixKeyType(EmailKeyType)
}

func (kt *PhonePixKeyType) Value() PixKeyType {
	return PixKeyType(PhoneKeyType)
}

func (kt *RandomPixKeyType) Value() PixKeyType {
	return PixKeyType(RandomKeyType)
}

func (kt *CnpjPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	_, err := value_object.NewCNPJ(key)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetTypeName()))
	}

	return nil
}

func (kt *CpfPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	_, err := value_object.NewCPF(key)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetTypeName()))
	}
	return nil
}

func (kt *EmailPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	_, err := value_object.NewEmail(key)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetTypeName()))
	}

	return nil
}

func (kt *PhonePixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(PhoneKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetTypeName()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetTypeName()))
	}
	return nil
}

func (kt *RandomPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(RandomKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetTypeName()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetTypeName()))
	}
	return nil
}
