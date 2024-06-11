package main

import (
	"context"
	"log"
	"log/slog"
	"math/rand"
	"os"

	"github.com/felipemagrassi/pix-api/configuration/database/postgres"
	"github.com/felipemagrassi/pix-api/configuration/env"
	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/infra/database/receiver_repository"
)

var (
	Banks          = []string{"Bradesco", "Itau", "Caixa", "Nubank", "Inter"}
	Status         = []entity.ReceiverStatus{entity.Valid, entity.Draft}
	Offices        = []string{"0001", "0002", "0003", "0004", "0005"}
	AccountNumbers = []string{"123456", "654321", "987654", "456789", "321654"}
	PixKey         = map[string]string{"email": "test@email.com", "cpf": "12345678901", "phone": "+5511999999999", "random": "7c7a2ba0-3fda-4f76-8c44-df1f8c1289ba", "cnpj": "41299131000107"}
	RandomNames    = []string{"John Doe", "Jane Doe", "John Smith", "Jane Smith", "John Johnson", "Jane Johnson"}
	RandomEmails   = []string{"johndoe@email.com", "janedoe@email.com", "janesmith@email.com", "johnjohnson@email.com"}
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
	receiverRepo := receiver_repository.NewReceiverRepository(db)

	_, err = db.Exec("DELETE FROM receivers")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 15; i++ {
		seedReceiver(receiverRepo, PixKey["cpf"], PixKey["cpf"], "cpf", randomArrayElement(RandomNames), randomArrayElement(RandomEmails), Status[rand.Intn(len(Status))])
		seedReceiver(receiverRepo, PixKey["cpf"], PixKey["email"], "email", randomArrayElement(RandomNames), randomArrayElement(RandomEmails), Status[rand.Intn(len(Status))])
		seedReceiver(receiverRepo, PixKey["cnpj"], PixKey["phone"], "phone", randomArrayElement(RandomNames), randomArrayElement(RandomEmails), Status[rand.Intn(len(Status))])
		seedReceiver(receiverRepo, PixKey["cnpj"], PixKey["random"], "random", randomArrayElement(RandomNames), randomArrayElement(RandomEmails), Status[rand.Intn(len(Status))])
		seedReceiver(receiverRepo, PixKey["cnpj"], PixKey["cnpj"], "cnpj", randomArrayElement(RandomNames), randomArrayElement(RandomEmails), Status[rand.Intn(len(Status))])
	}

	res, err := db.Query("SELECT COUNT(*) FROM receivers")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	for res.Next() {
		err = res.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Total of receivers: %d\n", count)
}

func seedReceiver(repo entity.ReceiverRepositoryInterface, document, pixKeyValue, pixKeyType, name, email string, status entity.ReceiverStatus) {
	receiver, err := entity.NewReceiver(
		document,
		pixKeyValue,
		pixKeyType,
		name,
		email,
	)
	if err != nil {
		log.Fatal(err)
	}

	receiver.Bank, receiver.Office, receiver.AccountNumber = seedBank()
	receiver.PixKey, err = entity.NewPixKey(pixKeyValue, pixKeyType)
	receiver.Status = status
	if err != nil {
		log.Fatal(err)
	}

	err = repo.CreateReceiver(context.Background(), receiver)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Receiver created: %s, %s, %s, %s, %s, %s, %s, %s, %s\n", receiver.ReceiverId, receiver.Name, receiver.Email, receiver.Bank, receiver.Office, receiver.AccountNumber, receiver.PixKey.KeyValue, receiver.PixKey.KeyType, receiver.Status)
}

func seedBank() (string, string, string) {
	return randomArrayElement(Banks),
		randomArrayElement(Offices),
		randomArrayElement(AccountNumbers)
}

func randomArrayElement(arr []string) string {
	n := rand.Intn(len(arr) - 1)
	return arr[n]
}
