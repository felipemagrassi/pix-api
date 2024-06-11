package receiver_controller

import (
	"log/slog"
	"strconv"

	"github.com/felipemagrassi/pix-api/configuration/rest_err"
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
	pixKeyValue := c.Query("pix_key")
	pixKeyType, convErr := strconv.Atoi(c.Query("pix_key_type"))
	if convErr != nil {
		pixKeyType = -1
	}
	pageInt, convErr := strconv.Atoi(c.Query("page"))
	if convErr != nil {
		pageInt = 1
	}

	findReceiverInput := receiver_usecase.FindReceiversInput{
		Status:      entity.ReceiverStatus(intStatus),
		Name:        name,
		PixKeyValue: pixKeyValue,
		PixKeyType:  entity.PixKeyType(pixKeyType),
		Page:        pageInt,
	}

	receivers, err := r.receiverUseCase.FindReceivers(c.Request.Context(), findReceiverInput)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		slog.Error("error finding receivers", errRest.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(200, receivers)
}

func (r *ReceiverController) FindReceiverById(c *gin.Context) {
	id := c.Param("receiverId")

	receiverId, parseErr := pkg_entity.ParseID(id)
	if parseErr != nil {
		errRest := rest_err.NewBadRequestError("Invalid ID", rest_err.Causes{Field: "id", Message: "Invalid ID"})
		slog.Error("error parsing id", parseErr)
		c.JSON(errRest.Code, errRest)
		return
	}

	receiver, err := r.receiverUseCase.FindReceiverById(c.Request.Context(), receiverId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		slog.Error("error finding receiver", errRest.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(200, receiver)
}

func (r *ReceiverController) CreateReceiver(c *gin.Context) {
	var createReceiverInput receiver_usecase.CreateReceiverInput

	if err := c.ShouldBindJSON(&createReceiverInput); err != nil {
		restErr := rest_err.NewBadRequestError("Invalid JSON", rest_err.Causes{Field: "json", Message: "Invalid JSON"})
		slog.Error("error binding json", restErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	err := r.receiverUseCase.CreateReceiver(c.Request.Context(), createReceiverInput)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		slog.Error("error creating receiver", restErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(201, gin.H{"message": "Receiver created successfully"})
}

func (r *ReceiverController) UpdateReceiver(c *gin.Context) {
	id := c.Param("receiverId")

	receiverId, parseErr := pkg_entity.ParseID(id)
	if parseErr != nil {
		restErr := rest_err.NewBadRequestError("Invalid ID", rest_err.Causes{Field: "id", Message: "Invalid ID"})
		slog.Error("error parsing id", restErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	var updateReceiverInput receiver_usecase.UpdateReceiverInput

	if err := c.ShouldBindJSON(&updateReceiverInput); err != nil {
		restErr := rest_err.NewBadRequestError("Invalid JSON", rest_err.Causes{Field: "json", Message: "Invalid JSON"})
		slog.Error("error binding json", err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := r.receiverUseCase.UpdateReceiver(c.Request.Context(), receiverId, updateReceiverInput)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		slog.Error("error updating receiver", restErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(200, gin.H{"message": "Receiver updated successfully"})
}

func (r *ReceiverController) DeleteReceivers(c *gin.Context) {
	ids := c.QueryMap("ids")

	receiverIds := make([]pkg_entity.ID, 0)
	for _, id := range ids {
		receiverId, parseErr := pkg_entity.ParseID(id)
		if parseErr != nil {
			restErr := rest_err.NewBadRequestError("Invalid ID", rest_err.Causes{Field: "id", Message: "Invalid ID"})
			slog.Error("error parsing id", restErr.Error())
			c.JSON(restErr.Code, restErr)
			return
		}
		receiverIds = append(receiverIds, receiverId)
	}

	err := r.receiverUseCase.DeleteReceivers(c.Request.Context(), receiver_usecase.DeleteReceiversInput{ReceiverIds: receiverIds})
	if err != nil {
		restErr := rest_err.ConvertError(err)
		slog.Error("error deleting receivers", restErr.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(204, nil)
}
