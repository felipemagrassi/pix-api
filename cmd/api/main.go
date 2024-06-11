package main

import (
	"context"

	"github.com/felipemagrassi/pix-api/configuration/database/postgres"
	"github.com/felipemagrassi/pix-api/configuration/env"
	"github.com/felipemagrassi/pix-api/configuration/logger"
	"github.com/felipemagrassi/pix-api/internal/infra/api/web/controller/receiver_controller"
	"github.com/felipemagrassi/pix-api/internal/infra/database/receiver_repository"
	"github.com/felipemagrassi/pix-api/internal/usecase/receiver_usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()

	config, err := env.LoadConfig(".")
	if err != nil {
		logger.Error("error loading config", err)
		return
	}

	db, err := postgres.InitializeDatabase(ctx, config.DBUrl)
	if err != nil {
		logger.Error("error initializing database", err)
		return
	}

	defer db.Close()

	receiverController := initDependencies(db)

	router := gin.Default()
	router.Run(config.WebServerPort)
}

func initDependencies(database *sqlx.DB) *receiver_controller.ReceiverController {
	receiverRepo := receiver_repository.NewReceiverRepository(database)
	receiverUseCase := receiver_usecase.NewReceiverUseCase(receiverRepo)
	receiverController := receiver_controller.NewReceiverController(receiverUseCase)

	return receiverController
}
