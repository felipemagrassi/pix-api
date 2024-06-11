package main

import (
	"context"
	"log"
	"os"

	"github.com/felipemagrassi/pix-api/configuration/database/postgres"
	"github.com/felipemagrassi/pix-api/configuration/env"
	"github.com/felipemagrassi/pix-api/internal/infra/api/web/controller/receiver_controller"
	"github.com/felipemagrassi/pix-api/internal/infra/database/receiver_repository"
	"github.com/felipemagrassi/pix-api/internal/usecase/receiver_usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()

	config, err := env.LoadConfig("cmd/api/.env")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db, err := postgres.InitializeDatabase(ctx, config.DBUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer db.Close()

	receiverController := initDependencies(db)

	router := gin.Default()

	router.GET("/receiver", receiverController.FindReceivers)
	router.GET("/receiver/:receiverId", receiverController.FindReceiverById)
	router.POST("/receiver", receiverController.CreateReceiver)
	router.PUT("/receiver/:receiverId", receiverController.UpdateReceiver)
	router.DELETE("/receiver", receiverController.DeleteReceivers)

	router.Run(config.WebServerPort)
}

func initDependencies(database *sqlx.DB) *receiver_controller.ReceiverController {
	receiverRepo := receiver_repository.NewReceiverRepository(database)
	receiverUseCase := receiver_usecase.NewReceiverUseCase(receiverRepo)
	receiverController := receiver_controller.NewReceiverController(receiverUseCase)

	return receiverController
}
