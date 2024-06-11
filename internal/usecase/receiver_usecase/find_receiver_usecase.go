package receiver_usecase

import (
	"context"

	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type FindReceiversInput struct {
	Status      entity.ReceiverStatus `json:"status"`
	Name        string                `json:"name"`
	PixKeyValue string                `json:"pix_key_value"`
	PixKeyType  entity.PixKeyType     `json:"pix_key_type"`
	Page        int                   `json:"page"`
}

type FindReceiversOutput struct {
	CurrentPage int                  `json:"current_page"`
	Receivers   []FindReceiverOutput `json:"receivers,omitempty"`
}

type FindReceiverOutput struct {
	ReceiverId    string                `json:"receiver_id,omitempty"`
	Name          string                `json:"name,omitempty"`
	Document      string                `json:"document,omitempty"`
	Email         string                `json:"email,omitempty"`
	Status        entity.ReceiverStatus `json:"status,omitempty"`
	Bank          string                `json:"bank,omitempty"`
	Office        string                `json:"office,omitempty"`
	AccountNumber string                `json:"account_number,omitempty"`
	PixKey        *PixKeyOutput         `json:"pix_key,omitempty"`
	CreatedAt     string                `json:"created_at" time_format:"2006-01-02T15:04:05Z07:00"`
	UpdatedAt     string                `json:"updated_at" time_format:"2006-01-02T15:04:05Z07:00"`
}

type PixKeyOutput struct {
	KeyValue string `json:"value,omitempty"`
	KeyType  string `json:"type,omitempty"`
}

func (uc *ReceiverUseCase) FindReceivers(ctx context.Context, input FindReceiversInput) (*FindReceiversOutput, *internal_error.InternalError) {
	receivers, err := uc.receiverRepository.FindReceivers(ctx, input.Status, input.Name, input.PixKeyValue, input.PixKeyType, input.Page)
	if err != nil {
		return nil, err
	}

	var receiversOutput []FindReceiverOutput
	for _, receiver := range receivers {
		output := FindReceiverOutput{
			ReceiverId:    receiver.ReceiverId.String(),
			Name:          receiver.Name,
			Document:      receiver.Document.String(),
			Email:         receiver.Email.String(),
			Status:        receiver.GetStatus(),
			Bank:          receiver.Bank,
			Office:        receiver.Office,
			AccountNumber: receiver.AccountNumber,
			CreatedAt:     receiver.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:     receiver.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		if receiver.PixKey != nil {
			output.PixKey = &PixKeyOutput{
				KeyValue: receiver.PixKey.KeyValue,
				KeyType:  receiver.PixKey.KeyType.GetTypeName(),
			}
		}

		receiversOutput = append(receiversOutput, output)
	}

	output := &FindReceiversOutput{
		CurrentPage: input.Page,
		Receivers:   receiversOutput,
	}

	return output, nil
}

func (uc *ReceiverUseCase) FindReceiverById(ctx context.Context, receiverId pkg_entity.ID) (*FindReceiverOutput, *internal_error.InternalError) {
	receiver, err := uc.receiverRepository.FindReceiver(ctx, receiverId)
	if err != nil {
		return nil, err
	}

	return &FindReceiverOutput{
		ReceiverId:    receiver.ReceiverId.String(),
		Name:          receiver.Name,
		Document:      receiver.Document.String(),
		Email:         receiver.Email.String(),
		Status:        receiver.GetStatus(),
		Bank:          receiver.Bank,
		Office:        receiver.Office,
		AccountNumber: receiver.AccountNumber,
		PixKey: &PixKeyOutput{
			KeyValue: receiver.PixKey.KeyValue,
			KeyType:  receiver.PixKey.KeyType.GetTypeName(),
		},
		CreatedAt: receiver.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: receiver.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
