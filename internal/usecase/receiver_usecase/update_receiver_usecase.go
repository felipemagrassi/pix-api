package receiver_usecase

import (
	"context"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type UpdateDraftedReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}

type UpdateValidReceiverInput struct {
	Email string
}

func (uc *ReceiverUseCase) UpdateDraftedReceiver(ctx context.Context, receiverId pkg_entity.ID, input UpdateDraftedReceiverInput) *internal_error.InternalError {
	receiver, err := uc.receiverRepository.FindReceiver(ctx, receiverId)
	if err != nil {
		return err
	}

	if err := receiver.UpdateDraftedReceiver(
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

func (uc *ReceiverUseCase) UpdateValidReceiver(ctx context.Context, receiverId pkg_entity.ID, input UpdateValidReceiverInput) *internal_error.InternalError {
	receiver, err := uc.receiverRepository.FindReceiver(ctx, receiverId)
	if err != nil {
		return err
	}

	if err := receiver.UpdateEmail(input.Email); err != nil {
		return err
	}

	return uc.receiverRepository.UpdateReceiver(ctx, receiver)
}
