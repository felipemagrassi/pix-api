package receiver_repository

import (
	"context"
	"time"

	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type MemoryReceiverRepository struct {
	Receivers []ReceiverEntity
}

func NewMemoryReceiverRepository() *ReceiverRepository {
	return &ReceiverRepository{}
}

func (r *MemoryReceiverRepository) FindReceiver(ctx context.Context, id pkg_entity.ID) (*entity.Receiver, *internal_error.InternalError) {
	var receiver ReceiverEntity
	for _, r := range r.Receivers {
		if r.ReceiverId == id {
			receiver = r
			break
		}
	}

	if receiver.ReceiverId.String() == "" {
		return nil, internal_error.NewNotFoundError("receiver not found")
	}

	entity := mapReceiverEntityToReceiver(receiver)

	return &entity, nil
}

func (r *MemoryReceiverRepository) FindReceivers(ctx context.Context, status entity.ReceiverStatus, name, pixKeyValue string, pixKeyType entity.PixKeyType, page int) ([]entity.Receiver, *internal_error.InternalError) {
	var receivers []entity.Receiver

	for _, receiver := range r.Receivers {
		if status != -1 && entity.ReceiverStatus(receiver.Status) != status {
			continue
		}

		if name != "" && receiver.Name != name {
			continue
		}

		if pixKeyValue != "" && receiver.PixKey != pixKeyValue {
			continue
		}

		if pixKeyType != -1 && entity.PixKeyType(receiver.PixKeyType) != pixKeyType {
			continue
		}

		receivers = append(receivers, mapReceiverEntityToReceiver(receiver))
	}

	return receivers, nil
}

func (r *MemoryReceiverRepository) CreateReceiver(ctx context.Context, receiver *entity.Receiver) *internal_error.InternalError {
	receiverEntity := ReceiverEntity{
		ReceiverId:    receiver.ReceiverId,
		Name:          receiver.Name,
		Document:      receiver.Document.String(),
		Email:         receiver.Email.String(),
		Status:        int(receiver.GetStatus()),
		PixKey:        receiver.PixKey.KeyValue,
		PixKeyType:    int(receiver.PixKey.KeyType.Value()),
		Bank:          receiver.Bank,
		Office:        receiver.Office,
		AccountNumber: receiver.AccountNumber,
		CreatedAt:     receiver.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     receiver.UpdatedAt.Format(time.RFC3339),
	}

	r.Receivers = append(r.Receivers, receiverEntity)

	return nil
}

func (r *MemoryReceiverRepository) UpdateReceiver(ctx context.Context, receiver *entity.Receiver) *internal_error.InternalError {
	receiverIndex := -1

	for i, r := range r.Receivers {
		if r.ReceiverId == receiver.ReceiverId {
			receiverIndex = i
			break
		}
	}

	if receiverIndex == -1 {
		return internal_error.NewNotFoundError("receiver not found")
	}

	receiverEntity := ReceiverEntity{
		ReceiverId:    receiver.ReceiverId,
		Name:          receiver.Name,
		Document:      receiver.Document.String(),
		Email:         receiver.Email.String(),
		Status:        int(receiver.GetStatus()),
		PixKey:        receiver.PixKey.KeyValue,
		PixKeyType:    int(receiver.PixKey.KeyType.Value()),
		Bank:          receiver.Bank,
		Office:        receiver.Office,
		AccountNumber: receiver.AccountNumber,
		CreatedAt:     receiver.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     receiver.UpdatedAt.Format(time.RFC3339),
	}

	r.Receivers[receiverIndex] = receiverEntity

	return nil
}

func (r *MemoryReceiverRepository) DeleteManyReceivers(ctx context.Context, ids []pkg_entity.ID) *internal_error.InternalError {
	receiverIndexes := make([]int, 0)

	for i, receiver := range r.Receivers {
		for _, id := range ids {
			if receiver.ReceiverId == id {
				receiverIndexes = append(receiverIndexes, i)
			}
		}
	}

	for _, index := range receiverIndexes {
		r.Receivers = append(r.Receivers[:index], r.Receivers[index+1:]...)
	}

	return nil
}
