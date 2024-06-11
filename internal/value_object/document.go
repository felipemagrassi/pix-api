package value_object

import (
	"fmt"
	"regexp"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
)

type (
	CPF  string
	CNPJ string
)

type Document interface {
	Validate() *internal_error.InternalError
	String() string
}

func (cpf CPF) String() string {
	return string(cpf)
}

func (cnpj CNPJ) String() string {
	return string(cnpj)
}

const (
	CnpjKeyPattern = `^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`
	CpfKeyPattern  = `^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`
)

func NewDocument(document string) (Document, *internal_error.InternalError) {
	if isValidCpf(document) {
		cpf, err := NewCPF(document)
		if err != nil {
			return nil, err
		}
		return cpf, nil

	} else if isValidCnpj(document) {
		cnpj, err := NewCNPJ(document)
		if err != nil {
			return nil, err
		}
		return cnpj, nil
	}

	return nil, internal_error.NewBadRequestError("Invalid Document", internal_error.Causes{Field: "document", Message: "Invalid Document"})
}

func NewCPF(cpf string) (CPF, *internal_error.InternalError) {
	newCpf := CPF(cpf)

	if err := newCpf.Validate(); err != nil {
		return "", err
	}

	return newCpf, nil
}

func NewCNPJ(cnpj string) (CNPJ, *internal_error.InternalError) {
	newCnpj := CNPJ(cnpj)

	if err := newCnpj.Validate(); err != nil {
		return "", err
	}

	return newCnpj, nil
}

func (cpf CPF) Validate() *internal_error.InternalError {
	re, err := regexp.Compile(CpfKeyPattern)
	if err != nil {
		message := fmt.Sprintf("Error compiling regex: %s", err.Error())
		return internal_error.NewInternalServerError(message, err)
	}

	if !re.MatchString(cpf.String()) {
		cause := internal_error.Causes{
			Field:   "cpf",
			Message: "Invalid CPF",
		}

		return internal_error.NewBadRequestError("Invalid CPF", cause)
	}

	return nil
}

func (cnpj CNPJ) Validate() *internal_error.InternalError {
	re, err := regexp.Compile(CnpjKeyPattern)
	if err != nil {
		message := fmt.Sprintf("Error compiling regex: %s", err.Error())
		return internal_error.NewInternalServerError(message, err)
	}

	if !re.MatchString(cnpj.String()) {
		cause := internal_error.Causes{
			Field:   "cnpj",
			Message: "Invalid CNPJ",
		}

		return internal_error.NewBadRequestError("Invalid CNPJ", cause)
	}

	return nil
}

func isValidCpf(cpf string) bool {
	_, err := NewCPF(cpf)
	if err != nil {
		return false
	}
	return true
}

func isValidCnpj(cnpj string) bool {
	_, err := NewCNPJ(cnpj)
	if err != nil {
		return false
	}

	return true
}
