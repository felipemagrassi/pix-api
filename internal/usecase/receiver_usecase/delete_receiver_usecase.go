package receiver_usecase

import (
	"context"

	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type DeleteReceiversInput struct {
	ReceiverIds []pkg_entity.ID
}

func (uc *ReceiverUseCase) DeleteReceivers(ctx context.Context, input DeleteReceiversInput) *internal_error.InternalError {
	return uc.receiverRepository.DeleteManyReceivers(ctx, input.ReceiverIds)
}
