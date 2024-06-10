package entity

import (
	"time"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/internal/value_object"
	"github.com/felipemagrassi/pix-api/pkg/entity"
)

type CreateReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}

type UpdateDraftedReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}

type FindReceiversInput struct {
	Status      ReceiverStatus
	Name        string
	PixKeyValue string
	PixKeyType  PixKeyType
}

type DeleteReceiversInput struct {
	ReceiverIds []entity.ID
}

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

func NewReceiver(
	createReceiverInput CreateReceiverInput,
) (*Receiver, *internal_error.InternalError) {
	newDocument, err := value_object.NewDocument(createReceiverInput.Document)
	if err != nil {
		return nil, err
	}

	pixKey, err := NewPixKey(createReceiverInput.PixKeyValue, createReceiverInput.PixKeyType)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()

	receiver := &Receiver{
		Id:        entity.NewID(),
		Name:      createReceiverInput.Name,
		Document:  newDocument,
		Email:     value_object.Email(createReceiverInput.Email),
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
	receiverUpdateInput UpdateDraftedReceiverInput,
) *internal_error.InternalError {
	if r.GetStatus() == Valid {
		return internal_error.NewBadRequestError("Receiver is already valid")
	}

	if receiverUpdateInput.Name != "" {
		r.Name = receiverUpdateInput.Name
	}

	if receiverUpdateInput.Document != "" {
		newDocument, err := value_object.NewDocument(receiverUpdateInput.Document)
		if err != nil {
			return err
		}
		r.Document = newDocument
	}

	if receiverUpdateInput.Email != "" {
		r.Email = value_object.Email(receiverUpdateInput.Email)
	}

	if receiverUpdateInput.PixKeyValue != "" && receiverUpdateInput.PixKeyType != "" {
		pixKey, err := NewPixKey(receiverUpdateInput.PixKeyValue, receiverUpdateInput.PixKeyType)
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
