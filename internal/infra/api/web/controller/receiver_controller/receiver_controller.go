package receiver_controller

import (
	"github.com/felipemagrassi/pix-api/internal/usecase/receiver_usecase"
	"github.com/gin-gonic/gin"
)

type ReceiverController struct {
	receiverUseCase receiver_usecase.ReceiverUseCaseInterface
}

func NewReceiverController(receiverUseCase receiver_usecase.ReceiverUseCaseInterface) *ReceiverController {
	return &ReceiverController{
		receiverUseCase: receiverUseCase,
	}
}

func (r *ReceiverController) FindReceivers(c *gin.Context) {
}

func (r *ReceiverController) FindReceiverById(c *gin.Context) {
}

func (r *ReceiverController) CreateReceiver(c *gin.Context) {
}

func (r *ReceiverController) UpdateDraftReceiver(c *gin.Context) {
}

func (r *ReceiverController) UpdateReceiverEmail(c *gin.Context) {
}

func (r *ReceiverController) DeleteReceivers(c *gin.Context) {
}
