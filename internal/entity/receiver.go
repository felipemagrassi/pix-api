package entity

import (
	"context"
	"time"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/internal/value_object"
	"github.com/felipemagrassi/pix-api/pkg/entity"
)

type (
	ReceiverStatus int
)

const (
	_ ReceiverStatus = iota
	Valid
	Draft
)

type Receiver struct {
	ReceiverId    entity.ID
	Name          string
	Document      value_object.Document
	Email         value_object.Email
	Status        ReceiverStatus
	Bank          string
	Office        string
	AccountNumber string
	PixKey        *PixKey
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ReceiverRepositoryInterface interface {
	FindReceiver(ctx context.Context, id entity.ID) (*Receiver, *internal_error.InternalError)
	FindReceivers(ctx context.Context, status ReceiverStatus, name, pixKeyValue string, pixKeyType PixKeyType) ([]Receiver, *internal_error.InternalError)
	CreateReceiver(ctx context.Context, receiver *Receiver) *internal_error.InternalError
	UpdateReceiver(ctx context.Context, receiver *Receiver) *internal_error.InternalError
	DeleteManyReceivers(ctx context.Context, ids []entity.ID) *internal_error.InternalError
}

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
		ReceiverId: entity.NewID(),
		Name:       name,
		Document:   newDocument,
		Email:      value_object.Email(email),
		Status:     Draft,
		PixKey:     pixKey,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	if err := receiver.Validate(); err != nil {
		return nil, err
	}

	return receiver, nil
}

func (r *Receiver) UpdateReceiver(
	document, pixKeyValue, pixKeyType, name, email string,
) *internal_error.InternalError {
	if r.GetStatus() == Valid {
		return r.updateValidReceiver(email)
	}

	return r.updateDraftReceiver(document, pixKeyValue, pixKeyType, name, email)
}

func (r *Receiver) Validate() *internal_error.InternalError {
	if r.Email.String() != "" {
		err := r.Email.Validate()
		if err != nil {
			return err
		}
	}

	if r.PixKey == nil {
		return internal_error.NewBadRequestError("Invalid Receiver", internal_error.Causes{Field: "pix_key", Message: "Pix Key is required"})
	}

	return nil
}

func (r *Receiver) GetStatus() ReceiverStatus {
	return r.Status
}

func (r *Receiver) ValidateReceiverStatus() {
	r.Status = Valid
}

func (r *Receiver) updateValidReceiver(email string) *internal_error.InternalError {
	if email != "" {
		r.Email = value_object.Email(email)
		r.UpdatedAt = time.Now()
	}

	return r.Validate()
}

func (r *Receiver) updateDraftReceiver(document, pixKeyValue, pixKeyType, name, email string) *internal_error.InternalError {
	if r.GetStatus() == Valid {
		return internal_error.NewBadRequestError("Receiver is already valid", internal_error.Causes{Field: "status", Message: "Receiver is already valid"})
	}

	if name == "" && document == "" && email == "" && pixKeyValue == "" && pixKeyType == "" {
		return nil
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
