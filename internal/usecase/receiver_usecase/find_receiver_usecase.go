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
}

type FindReceiverOutput struct {
	Id            string                `json:"id"`
	Name          string                `json:"name"`
	Document      string                `json:"document"`
	Email         string                `json:"email"`
	Status        entity.ReceiverStatus `json:"status"`
	Bank          string                `json:"bank"`
	Office        string                `json:"office"`
	AccountNumber string                `json:"account_number"`
	PixKey        PixKeyOutput          `json:"pix_key"`
	CreatedAt     string                `json:"created_at" time_format:"2006-01-02T15:04:05Z07:00"`
	UpdatedAt     string                `json:"updated_at" time_format:"2006-01-02T15:04:05Z07:00"`
}

type PixKeyOutput struct {
	KeyValue string `json:"value"`
	KeyType  string `json:"type"`
}

func (uc *ReceiverUseCase) FindReceivers(ctx context.Context, input FindReceiversInput) ([]FindReceiverOutput, *internal_error.InternalError) {
	return nil, nil
}

func (uc *ReceiverUseCase) FindReceiverById(ctx context.Context, receiverId pkg_entity.ID) (*FindReceiverOutput, *internal_error.InternalError) {
	receiver, err := uc.receiverRepository.FindReceiver(ctx, receiverId)
	if err != nil {
		return nil, err
	}

	return &FindReceiverOutput{
		Id:            receiver.Id.String(),
		Name:          receiver.Name,
		Document:      receiver.Document.String(),
		Email:         receiver.Email.String(),
		Status:        receiver.GetStatus(),
		Bank:          receiver.Bank,
		Office:        receiver.Office,
		AccountNumber: receiver.AccountNumber,
		PixKey: PixKeyOutput{
			KeyValue: receiver.PixKey.KeyValue,
			KeyType:  receiver.PixKey.KeyType.GetTypeName(),
		},
		CreatedAt: receiver.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: receiver.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
