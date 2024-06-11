package receiver_usecase

import (
	"context"
	"log/slog"

	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/internal_error"
)

type CreateReceiverInput struct {
	Name        string `json:"name"`
	Document    string `json:"document"`
	Email       string `json:"email"`
	PixKeyValue string `json:"pix_key_value"`
	PixKeyType  string `json:"pix_key_type"`
}

func (uc *ReceiverUseCase) CreateReceiver(ctx context.Context, input CreateReceiverInput) *internal_error.InternalError {
	entity, err := entity.NewReceiver(
		input.Document,
		input.PixKeyValue,
		input.PixKeyType,
		input.Name,
		input.Email,
	)
	if err != nil {
		slog.Error("error creating receiver entity", err)
		return err
	}

	return uc.receiverRepository.CreateReceiver(ctx, entity)
}
