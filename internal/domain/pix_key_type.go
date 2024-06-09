package domain

import (
	"fmt"
	"regexp"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
)

const (
	CnpjKeyPattern   = `^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`
	CpfKeyPattern    = `^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`
	EmailKeyPattern  = `^[a-z0-9+_.-]+@[a-z0-9.-]+$`
	PhoneKeyPattern  = `^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`
	RandomKeyPattern = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
)

type PixKeyType int

const (
	CnpjKeyType PixKeyType = iota
	CpfKeyType
	EmailKeyType
	PhoneKeyType
	RandomKeyType
)

func (pkt PixKeyType) String() string {
	switch pkt {
	case CnpjKeyType:
		return "Cnpj"
	case CpfKeyType:
		return "Cpf"
	case EmailKeyType:
		return "Email"
	case PhoneKeyType:
		return "Phone"
	case RandomKeyType:
		return "Random"
	default:
		panic("Unknown Pix Key Type")

	}
}

type PixKeyTypeInterface interface {
	ValidateKeyType(key string) *internal_error.InternalError
	GetType() string
}

func NewPixKeyType(keyType PixKeyType) (PixKeyTypeInterface, *internal_error.InternalError) {
	switch keyType {
	case CnpjKeyType:
		return &CnpjPixKeyType{keyType: keyType}, nil
	case CpfKeyType:
		return &CpfPixKeyType{keyType: keyType}, nil
	case EmailKeyType:
		return &EmailPixKeyType{keyType: keyType}, nil
	case PhoneKeyType:
		return &PhonePixKeyType{keyType: keyType}, nil
	case RandomKeyType:
		return &RandomPixKeyType{keyType: keyType}, nil
	default:
		return nil, internal_error.NewBadRequestError("Invalid Key Type")
	}
}

type (
	CnpjPixKeyType   struct{ keyType PixKeyType }
	CpfPixKeyType    struct{ keyType PixKeyType }
	EmailPixKeyType  struct{ keyType PixKeyType }
	PhonePixKeyType  struct{ keyType PixKeyType }
	RandomPixKeyType struct{ keyType PixKeyType }
)

func (kt *CnpjPixKeyType) GetType() string {
	return kt.keyType.String()
}

func (kt *CpfPixKeyType) GetType() string {
	return kt.keyType.String()
}

func (kt *EmailPixKeyType) GetType() string {
	return kt.keyType.String()
}

func (kt *PhonePixKeyType) GetType() string {
	return kt.keyType.String()
}

func (kt *RandomPixKeyType) GetType() string {
	return kt.keyType.String()
}

func (kt *CnpjPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(CnpjKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetType()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetType()))
	}

	return nil
}

func (kt *CpfPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(CpfKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetType()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetType()))
	}

	return nil
}

func (kt *EmailPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(EmailKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetType()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetType()))
	}
	return nil
}

func (kt *PhonePixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(PhoneKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetType()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetType()))
	}
	return nil
}

func (kt *RandomPixKeyType) ValidateKeyType(key string) *internal_error.InternalError {
	re, err := regexp.Compile(RandomKeyPattern)
	if err != nil {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s key matching pattern", kt.GetType()))
	}

	if !re.MatchString(key) {
		return internal_error.NewBadRequestError(fmt.Sprintf("Invalid %s Key", kt.GetType()))
	}
	return nil
}
