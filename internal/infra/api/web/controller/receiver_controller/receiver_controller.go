package receiver_controller

import (
	"log/slog"
	"strconv"

	"github.com/felipemagrassi/pix-api/internal/entity"
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
	intStatus, convErr := strconv.Atoi(c.Query("status"))
	if convErr != nil {
		intStatus = -1
	}
	name := c.Query("name")
	pixKeyValue := c.Query("pix_key_value")
	pixKeyType, convErr := strconv.Atoi(c.Query("pix_key_type"))
	if convErr != nil {
		pixKeyType = -1
	}

	findReceiverInput := receiver_usecase.FindReceiversInput{
		Status:      entity.ReceiverStatus(intStatus),
		Name:        name,
		PixKeyValue: pixKeyValue,
		PixKeyType:  entity.PixKeyType(pixKeyType),
	}

	receivers, err := r.receiverUseCase.FindReceivers(c.Request.Context(), findReceiverInput)
	if err != nil {
		slog.Error("error finding receivers", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, receivers)
}

func (r *ReceiverController) FindReceiverById(c *gin.Context) {
	id := c.Param("receiverId")

	receiverId, parseErr := pkg_entity.ParseID(id)
	if parseErr != nil {
		slog.Error("error parsing id", parseErr)
		c.JSON(400, parseErr)
		return
	}

	receiver, err := r.receiverUseCase.FindReceiverById(c.Request.Context(), receiverId)
	if err != nil {
		slog.Error("error finding receiver", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, receiver)
}

func (r *ReceiverController) CreateReceiver(c *gin.Context) {
	var createReceiverInput receiver_usecase.CreateReceiverInput

	if err := c.ShouldBindJSON(&createReceiverInput); err != nil {
		slog.Error("error binding json", err)
		c.JSON(400, err)
		return
	}

	err := r.receiverUseCase.CreateReceiver(c.Request.Context(), createReceiverInput)
	if err != nil {
		slog.Error("error creating receiver", err)
		c.JSON(500, err)
		return
	}

	c.JSON(201, gin.H{"message": "Receiver created successfully"})
}

func (r *ReceiverController) UpdateReceiver(c *gin.Context) {
	id := c.Param("receiverId")

	receiverId, parseErr := pkg_entity.ParseID(id)
	if parseErr != nil {
		c.JSON(400, parseErr)
		return
	}

	var updateReceiverInput receiver_usecase.UpdateReceiverInput

	if err := c.ShouldBindJSON(&updateReceiverInput); err != nil {
		c.JSON(400, err)
		return
	}

	err := r.receiverUseCase.UpdateReceiver(c.Request.Context(), receiverId, updateReceiverInput)
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
		receiverId, parseErr := pkg_entity.ParseID(id)
		if parseErr != nil {
			c.JSON(400, parseErr)
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
