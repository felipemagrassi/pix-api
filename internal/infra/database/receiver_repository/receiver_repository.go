package receiver_repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	"github.com/felipemagrassi/pix-api/internal/value_object"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type ReceiverEntity struct {
	ReceiverId    pkg_entity.ID `db:"receiver_id"`
	Name          string        `db:"name"`
	Document      string        `db:"document"`
	Email         string        `db:"email"`
	Status        int           `db:"status"`
	PixKey        string        `db:"pix_key"`
	PixKeyType    int           `db:"pix_key_type"`
	Bank          string        `db:"bank"`
	Office        string        `db:"office"`
	AccountNumber string        `db:"account_number"`
	CreatedAt     string        `db:"created_at"`
	UpdatedAt     string        `db:"updated_at"`
}

type ReceiverRepository struct {
	Db *sqlx.DB
}

func NewReceiverRepository(db *sqlx.DB) *ReceiverRepository {
	return &ReceiverRepository{Db: db}
}

func (r *ReceiverRepository) FindReceiver(ctx context.Context, id pkg_entity.ID) (*entity.Receiver, *internal_error.InternalError) {
	var receiver ReceiverEntity
	err := r.Db.GetContext(ctx, &receiver, "SELECT * FROM receivers WHERE receiver_id = $1", id)
	if err != nil {
		slog.Error("error finding receiver", err)
		return nil, internal_error.NewNotFoundError("receiver not found")
	}

	entity := mapReceiverEntityToReceiver(receiver)
	return &entity, nil
}

func (r *ReceiverRepository) FindReceivers(ctx context.Context, status entity.ReceiverStatus, name, pixKeyValue string, pixKeyType entity.PixKeyType) ([]entity.Receiver, *internal_error.InternalError) {
	var receivers []entity.Receiver

	baseQuery := "SELECT receiver_id, name, document, bank, office, account_number, status FROM receivers WHERE 1=1"

	args := []interface{}{}

	if status != -1 {
		baseQuery += " AND status = $1"
		args = append(args, status)
	}

	if name != "" {
		baseQuery += " AND name ILIKE %$2%"
		args = append(args, name)
	}

	if pixKeyValue != "" {
		baseQuery += " AND pix_key = $3"
		args = append(args, pixKeyValue)
	}

	if pixKeyType != -1 {
		baseQuery += " AND pix_key_type = $4"
		args = append(args, pixKeyType)
	}

	defaultOrder := " ORDER BY created_at DESC"
	baseQuery += defaultOrder

	rows, err := r.Db.QueryxContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, internal_error.NewInternalServerError("error finding receivers")
	}

	for rows.Next() {
		var receiver ReceiverEntity
		err := rows.StructScan(&receiver)
		if err != nil {
			return nil, internal_error.NewInternalServerError("error finding receivers")
		}

		receivers = append(receivers, mapReceiverEntityToReceiver(receiver))
	}

	return receivers, nil
}

func (r *ReceiverRepository) CreateReceiver(ctx context.Context, receiver *entity.Receiver) *internal_error.InternalError {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO receivers (receiver_id, name, document, email, status, pix_key, pix_key_type, bank, office, account_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", receiver.ReceiverId, receiver.Name, receiver.Document.String(), receiver.Email.String(), receiver.GetStatus(), receiver.PixKey.KeyValue, receiver.PixKey.KeyType.Value(), receiver.Bank, receiver.Office, receiver.AccountNumber, receiver.CreatedAt, receiver.UpdatedAt)
	if err != nil {
		slog.Error("error creating receiver", err)
		return internal_error.NewInternalServerError("error creating receiver")
	}

	return nil
}

func (r *ReceiverRepository) UpdateReceiver(ctx context.Context, receiver *entity.Receiver) *internal_error.InternalError {
	_, err := r.Db.ExecContext(ctx, "UPDATE receivers SET name = $1, document = $2, email = $3, status = $4, pix_key = $5, pix_key_type = $6, bank = $7, office = $8, account_number = $9, updated_at = $10 WHERE receiver_id = $11", receiver.Name, receiver.Document.String(), receiver.Email.String(), receiver.GetStatus(), receiver.PixKey.KeyValue, receiver.PixKey.KeyType.Value(), receiver.Bank, receiver.Office, receiver.AccountNumber, receiver.UpdatedAt, receiver.ReceiverId)
	if err != nil {
		slog.Error("error updating receiver", err)
		return internal_error.NewInternalServerError("error updating receiver")
	}

	return nil
}

func (r *ReceiverRepository) DeleteManyReceivers(ctx context.Context, ids []pkg_entity.ID) *internal_error.InternalError {
	idsString := ""
	for i, id := range ids {
		if i == 0 {
			idsString += fmt.Sprintf("\"%s\"", id.String())
		} else {
			idsString += fmt.Sprintf(", \"%s\"", id.String())
		}
	}

	query := fmt.Sprintf("DELETE FROM receivers WHERE receiver_id = ANY('{%s}')", idsString)
	fmt.Println(query)
	res, err := r.Db.ExecContext(ctx, query)
	if err != nil {
		slog.Error("error deleting receivers", err)
		return internal_error.NewInternalServerError("error deleting receivers")
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return internal_error.NewNotFoundError("receivers not found")
	}

	return nil
}

func mapReceiverEntityToReceiver(receiverEntity ReceiverEntity) entity.Receiver {
	document, _ := value_object.NewDocument(receiverEntity.Document)
	email, _ := value_object.NewEmail(receiverEntity.Email)

	status := entity.ReceiverStatus(receiverEntity.Status)
	createdAt, _ := time.Parse(time.RFC3339, receiverEntity.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, receiverEntity.UpdatedAt)

	receiver := entity.Receiver{
		ReceiverId:    receiverEntity.ReceiverId,
		Name:          receiverEntity.Name,
		Document:      document,
		Email:         email,
		Status:        status,
		Bank:          receiverEntity.Bank,
		Office:        receiverEntity.Office,
		AccountNumber: receiverEntity.AccountNumber,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}

	receiver.PixKey = nil

	if receiverEntity.PixKey != "" {
		pixKeyType, _ := entity.NewPixKeyType(entity.PixKeyType(receiverEntity.PixKeyType))
		pixKey := entity.PixKey{
			KeyValue: receiverEntity.PixKey,
			KeyType:  pixKeyType,
		}

		receiver.PixKey = &pixKey
	}

	return receiver
}
