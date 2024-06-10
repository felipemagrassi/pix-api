package entity

import (
	"time"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/internal/value_object"
	"github.com/felipemagrassi/pix-api/pkg/entity"
)

type (
	ReceiverStatus int
)

const (
	Valid ReceiverStatus = iota
	Draft
)

type Receiver struct {
	Id        entity.ID
	Name      string
	Document  value_object.Document
	Email     value_object.Email
	Status    ReceiverStatus
	PixKey    *PixKey
	Timestamp time.Time
}

func NewReceiver(name, email, document string, pixKey *PixKey) (*Receiver, *internal_error.InternalError) {
	newDocument, err := value_object.NewDocument(document)
	if err != nil {
		return nil, err
	}

	receiver := &Receiver{
		Id:       entity.NewID(),
		Name:     name,
		Document: newDocument,
		Email:    value_object.Email(email),
		Status:   Draft,
		PixKey:   pixKey,
	}

	if err := receiver.Validate(); err != nil {
		return nil, internal_error.NewBadRequestError("Invalid Receiver")
	}

	return receiver, nil
}

func (r *Receiver) Validate() *internal_error.InternalError {
	if r.Email.String() != "" {
		err := r.Email.Validate()
		if err != nil {
			return err
		}
	}

	if r.PixKey == nil {
		return internal_error.NewBadRequestError("Invalid Receiver")
	}

	return nil
}
