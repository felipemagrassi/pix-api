package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	pg "github.com/felipemagrassi/pix-api/configuration/database/postgres"
	"github.com/felipemagrassi/pix-api/internal/infra/api/web/controller/receiver_controller"
	"github.com/felipemagrassi/pix-api/internal/infra/database/receiver_repository"
	"github.com/felipemagrassi/pix-api/internal/usecase/receiver_usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func (suite *ReceiverTestSuite) TestCanCreateAndFetchReceivers() {
	server := initServer(suite.Db)
	defer server.Close()

	body := []byte(`{"name": "Felipe", "document": "12345678900", "email": "felipe@test.com", "pix_key_value": "12345678901", "pix_key_type": "cpf"}`)

	jsonBody := bytes.NewReader(body)

	client := server.Client()
	req, err := http.NewRequest(http.MethodPost, server.URL+"/receiver", jsonBody)
	assert.NoError(suite.T(), err)

	res, err := client.Do(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, res.StatusCode)

	req, err = http.NewRequest(http.MethodGet, server.URL+"/receiver", nil)
	assert.NoError(suite.T(), err)

	res, err = client.Do(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var receivers_output receiver_usecase.FindReceiversOutput

	err = json.NewDecoder(res.Body).Decode(&receivers_output)
	defer res.Body.Close()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(receivers_output.Receivers))

	id := receivers_output.Receivers[0].ReceiverId
	assert.Equal(suite.T(), "12345678900", receivers_output.Receivers[0].Document)

	updateBody := []byte(`{"name": "Felipe", "document": "12345678901", "email": "felipe@test.com", "pix_key_value": "12345678902", "pix_key_type": "cpf"}`)
	jsonBody = bytes.NewReader(updateBody)

	req, err = http.NewRequest(http.MethodPut, server.URL+"/receiver/"+id, jsonBody)
	assert.NoError(suite.T(), err)

	res, err = client.Do(req)
	assert.NoError(suite.T(), err)

	req, err = http.NewRequest(http.MethodGet, server.URL+"/receiver/"+id, nil)
	assert.NoError(suite.T(), err)

	res, err = client.Do(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var receiver_output receiver_usecase.FindReceiverOutput

	err = json.NewDecoder(res.Body).Decode(&receiver_output)
	defer res.Body.Close()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "12345678901", receiver_output.Document)

	req, err = http.NewRequest(http.MethodDelete, server.URL+"/receiver", nil)
	assert.NoError(suite.T(), err)
	params := req.URL.Query()
	params.Set("ids[0]", id)

	req.URL.RawQuery = params.Encode()

	res, err = client.Do(req)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusNoContent, res.StatusCode)

	req, err = http.NewRequest(http.MethodGet, server.URL+"/receiver/"+id, nil)
	assert.NoError(suite.T(), err)

	res, err = client.Do(req)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusNotFound, res.StatusCode)
}

func initServer(db *sqlx.DB) *httptest.Server {
	controller := initDependencies(db)

	g := gin.New()

	g.GET("/receiver", controller.FindReceivers)
	g.GET("/receiver/:receiverId", controller.FindReceiverById)
	g.POST("/receiver", controller.CreateReceiver)
	g.PUT("/receiver/:receiverId", controller.UpdateReceiver)
	g.DELETE("/receiver", controller.DeleteReceivers)

	return httptest.NewServer(g)
}

func initDependencies(db *sqlx.DB) *receiver_controller.ReceiverController {
	receiverRepo := receiver_repository.NewReceiverRepository(db)
	receiverUseCase := receiver_usecase.NewReceiverUseCase(receiverRepo)
	receiverController := receiver_controller.NewReceiverController(receiverUseCase)

	return receiverController
}

type ReceiverTestSuite struct {
	suite.Suite
	Db        *sqlx.DB
	Container testcontainers.Container
}

func (suite *ReceiverTestSuite) SetupSuite() {
	ctx := context.Background()

	dbName := "receivers"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(suite.T(), err)

	db, err := pg.InitializeDatabase(ctx, connStr, "../../../..")
	suite.NoError(err)
	suite.Db = db
	suite.Container = postgresContainer
}

func (suite *ReceiverTestSuite) TearDownTest() {
	suite.Db.Close()
}

func (suite *ReceiverTestSuite) TearDownSuite() {
	ctx := context.Background()
	suite.NoError(suite.Container.Terminate(ctx))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ReceiverTestSuite))
}
