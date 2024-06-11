package receiver_usecase

import (
	"context"

	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
)

type ReceiverUseCaseInterface interface {
	CreateReceiver(
		ctx context.Context,
		input CreateReceiverInput,
	) *internal_error.InternalError

	UpdateReceiver(
		ctx context.Context,
		receiverId pkg_entity.ID, input UpdateReceiverInput,
	) *internal_error.InternalError

	FindReceivers(
		ctx context.Context,
		input FindReceiversInput,
	) (*FindReceiversOutput, *internal_error.InternalError)

	FindReceiverById(
		ctx context.Context,
		receiverId pkg_entity.ID,
	) (*FindReceiverOutput, *internal_error.InternalError)

	DeleteReceivers(
		ctx context.Context,
		input DeleteReceiversInput,
	) *internal_error.InternalError
}

type ReceiverUseCase struct {
	receiverRepository entity.ReceiverRepositoryInterface
}

func NewReceiverUseCase(receiverRepository entity.ReceiverRepositoryInterface) *ReceiverUseCase {
	return &ReceiverUseCase{receiverRepository: receiverRepository}
}
