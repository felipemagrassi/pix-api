package domain

import (
	"regexp"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
)

type PixKeyType interface {
	ValidateKey() *internal_error.InternalError
}

type CnpjPixKeyType struct {
	Name string `json:"name"`
	key  string
}
type CpfPixKeyType struct {
	Name string `json:"name"`
	key  string
}
type EmailPixKeyType struct {
	Name string `json:"name"`
	key  string
}
type PhonePixKeyType struct {
	Name string `json:"name"`
	key  string
}
type RandomPixKeyType struct {
	Name string `json:"name"`
	key  string
}

func NewCnpjPixKeyType(key string) *CnpjPixKeyType {
	return &CnpjPixKeyType{
		Name: "CNPJ",
		key:  key,
	}
}

func NewCpfPixKeyType(key string) *CpfPixKeyType {
	return &CpfPixKeyType{
		Name: "CPF",
		key:  key,
	}
}

func NewEmailPixKeyType(key string) *EmailPixKeyType {
	return &EmailPixKeyType{
		Name: "Email",
		key:  key,
	}
}

func NewPhonePixKeyType(key string) *PhonePixKeyType {
	return &PhonePixKeyType{
		Name: "Phone",
		key:  key,
	}
}

func NewRandomPixKeyType(key string) *RandomPixKeyType {
	return &RandomPixKeyType{
		Name: "Random",
		key:  key,
	}
}

func (kt *CnpjPixKeyType) ValidateKey() *internal_error.InternalError {
	pattern := `^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return internal_error.NewBadRequestError("Invalid CNPJ matching pattern")
	}

	if !re.MatchString(kt.key) {
		return internal_error.NewBadRequestError("Invalid CNPJ Key")
	}

	return nil
}

func (kt *CpfPixKeyType) ValidateKey() *internal_error.InternalError {
	return nil
}

func (kt *EmailPixKeyType) ValidateKey() *internal_error.InternalError {
	return nil
}

func (kt *PhonePixKeyType) ValidateKey() *internal_error.InternalError {
	return nil
}

func (kt *RandomPixKeyType) ValidateKey() *internal_error.InternalError {
	return nil
}
