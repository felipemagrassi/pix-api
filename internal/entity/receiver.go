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
	Id            entity.ID
	Name          string
	Document      value_object.Document
	Email         value_object.Email
	status        ReceiverStatus
	Bank          string
	Office        string
	AccountNumber string
	PixKey        *PixKey
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ReceiverRepositoryInterface interface{}

func NewReceiver(
	document, pixKeyValue, pixKeyType, name, email string,
) (*Receiver, *internal_error.InternalError) {
	newDocument, err := value_object.NewDocument(document)
	if err != nil {
		return nil, err
	}

	pixKey, err := NewPixKey(pixKeyValue, pixKeyType)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()

	receiver := &Receiver{
		Id:        entity.NewID(),
		Name:      name,
		Document:  newDocument,
		Email:     value_object.Email(email),
		status:    Draft,
		PixKey:    pixKey,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := receiver.Validate(); err != nil {
		return nil, internal_error.NewBadRequestError("Invalid Receiver")
	}

	return receiver, nil
}

func (r *Receiver) UpdateEmail(email string) *internal_error.InternalError {
	r.Email = value_object.Email(email)
	return r.Validate()
}

func (r *Receiver) UpdateDraftedReceiver(
	document, pixKeyValue, pixKeyType, name, email string,
) *internal_error.InternalError {
	if r.GetStatus() == Valid {
		return internal_error.NewBadRequestError("Receiver is already valid")
	}

	if name != "" {
		r.Name = name
	}

	if document != "" {
		newDocument, err := value_object.NewDocument(document)
		if err != nil {
			return err
		}
		r.Document = newDocument
	}

	if email != "" {
		r.Email = value_object.Email(email)
	}

	if pixKeyValue != "" && pixKeyType != "" {
		pixKey, err := NewPixKey(pixKeyValue, pixKeyType)
		if err != nil {
			return err
		}
		r.PixKey = pixKey
	}

	r.UpdatedAt = time.Now()
	return r.Validate()
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

func (r *Receiver) GetStatus() ReceiverStatus {
	return r.status
}

func (r *Receiver) ValidateReceiverStatus() {
	r.status = Valid
}
