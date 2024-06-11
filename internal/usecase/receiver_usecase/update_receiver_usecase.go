package receiver_usecase

import (
	"context"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type UpdateReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}

func (uc *ReceiverUseCase) UpdateReceiver(ctx context.Context, receiverId pkg_entity.ID, input UpdateReceiverInput) *internal_error.InternalError {
	receiver, err := uc.receiverRepository.FindReceiver(ctx, receiverId)
	if err != nil {
		return err
	}

	if err := receiver.UpdateReceiver(
		input.Name,
		input.Document,
		input.Email,
		input.PixKeyValue,
		input.PixKeyType,
	); err != nil {
		return err
	}

	return uc.receiverRepository.UpdateReceiver(ctx, receiver)
}
