package domain

import (
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/pkg/domain"
)

type (
	ReceiverStatus int
)

const (
	Valid ReceiverStatus = iota
	Draft
)

type Receiver struct {
	Id     domain.ID
	Name   string
	Email  string
	Status ReceiverStatus
	PixKey PixKey
}

func NewReceiver(name, email string, pixKey PixKey) (*Receiver, *internal_error.InternalError) {
	receiver := &Receiver{
		Name:   name,
		Email:  email,
		Status: Draft,
		PixKey: pixKey,
	}

	if err := receiver.Validate(); err != nil {
		return nil, internal_error.NewBadRequestError("Invalid Receiver")
	}

	return receiver, nil
}

func (r *Receiver) Validate() *internal_error.InternalError {
	return nil
}
