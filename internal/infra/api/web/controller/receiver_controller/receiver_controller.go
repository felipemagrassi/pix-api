package receiver_controller

import (
	"github.com/felipemagrassi/pix-api/internal/usecase/receiver_usecase"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
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
	var findReceiverInput receiver_usecase.FindReceiversInput

	if err := c.ShouldBindJSON(&findReceiverInput); err != nil {
		c.JSON(400, err)
		return
	}

	receivers, err := r.receiverUseCase.FindReceivers(c.Request.Context(), findReceiverInput)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, receivers)
}

func (r *ReceiverController) FindReceiverById(c *gin.Context) {
	id := c.Param("id")

	receiverId, err := pkg_entity.ParseID(id)
	if err != nil {
		c.JSON(400, err)
		return
	}

	receiver, err := r.receiverUseCase.FindReceiverById(c.Request.Context(), receiverId)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, receiver)
}

func (r *ReceiverController) CreateReceiver(c *gin.Context) {
	var createReceiverInput receiver_usecase.CreateReceiverInput

	if err := c.ShouldBindJSON(&createReceiverInput); err != nil {
		c.JSON(400, err)
		return
	}

	err := r.receiverUseCase.CreateReceiver(c.Request.Context(), createReceiverInput)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(201, gin.H{"message": "Receiver created successfully"})
}

func (r *ReceiverController) UpdateReceiver(c *gin.Context) {
	id := c.Param("id")

	receiverId, err := pkg_entity.ParseID(id)
	if err != nil {
		c.JSON(400, err)
		return
	}

	var updateReceiverInput receiver_usecase.UpdateReceiverInput

	if err := c.ShouldBindJSON(&updateReceiverInput); err != nil {
		c.JSON(400, err)
		return
	}

	err = r.receiverUseCase.UpdateReceiver(c.Request.Context(), receiverId, updateReceiverInput)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{"message": "Receiver updated successfully"})
}

func (r *ReceiverController) DeleteReceivers(c *gin.Context) {
	ids := c.QueryArray("ids")

	receiverIds := make([]pkg_entity.ID, 0)
	for _, id := range ids {
		receiverId, err := pkg_entity.ParseID(id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		receiverIds = append(receiverIds, receiverId)
	}

	err := r.receiverUseCase.DeleteReceivers(c.Request.Context(), receiver_usecase.DeleteReceiversInput{ReceiverIds: receiverIds})
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(204, nil)
}
