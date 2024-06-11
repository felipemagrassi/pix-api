package receiver_usecase

import (
	"context"
	"log/slog"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type UpdateReceiverInput struct {
	Name        string `json:"name"`
	Document    string `json:"document"`
	Email       string `json:"email"`
	PixKeyValue string `json:"pix_key_value"`
	PixKeyType  string `json:"pix_key_type"`
}

func (uc *ReceiverUseCase) UpdateReceiver(ctx context.Context, receiverId pkg_entity.ID, input UpdateReceiverInput) *internal_error.InternalError {
	receiver, err := uc.receiverRepository.FindReceiver(ctx, receiverId)
	if err != nil {
		slog.Error("error finding receiver")
		return err
	}

	if err := receiver.UpdateReceiver(
		input.Document,
		input.PixKeyValue,
		input.PixKeyType,
		input.Name,
		input.Email,
	); err != nil {
		slog.Error("error updating receiver")
		return err
	}

	return uc.receiverRepository.UpdateReceiver(ctx, receiver)
}
