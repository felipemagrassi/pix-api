package receiver_repository

import (
	"context"

	"github.com/felipemagrassi/pix-api/internal/entity"
	"github.com/felipemagrassi/pix-api/internal/internal_error"
	pkg_entity "github.com/felipemagrassi/pix-api/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type ReceiverRepository struct {
	Db *sqlx.DB
}

func NewReceiverRepository(db *sqlx.DB) *ReceiverRepository {
	return &ReceiverRepository{Db: db}
}

func (r *ReceiverRepository) FindReceiver(ctx context.Context, id pkg_entity.ID) (*entity.Receiver, *internal_error.InternalError) {
	var receiver entity.Receiver
	err := r.Db.GetContext(ctx, &receiver, "SELECT * FROM receivers WHERE id = $1", id)
	if err != nil {
		return nil, internal_error.NewNotFoundError("receiver not found")
	}

	return &receiver, nil
}

func (r *ReceiverRepository) CreateReceiver(ctx context.Context, receiver *entity.Receiver) *internal_error.InternalError {
	_, err := r.Db.ExecContext(ctx, "INSERT INTO receivers (receiver_id, name, document, email, status, pix_key, pix_key_type, bank, office, account_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11, $12", receiver.Id, receiver.Name, receiver.Document.String(), receiver.Email.String(), receiver.GetStatus(), receiver.PixKey.KeyValue, receiver.PixKey.KeyType, receiver.Bank, receiver.Office, receiver.AccountNumber, receiver.CreatedAt, receiver.UpdatedAt)
	if err != nil {
		return internal_error.NewInternalServerError("error creating receiver")
	}

	return nil
}

func (r *ReceiverRepository) UpdateReceiver(ctx context.Context, receiver *entity.Receiver) *internal_error.InternalError {
	_, err := r.Db.ExecContext(ctx, "UPDATE receivers SET name = $1, document = $2, email = $3, status = $4, pix_key = $5, pix_key_type = $6, bank = $7, office = $8, account_number = $9, updated_at = $10 WHERE id = $11", receiver.Name, receiver.Document.String(), receiver.Email.String(), receiver.GetStatus(), receiver.PixKey.KeyValue, receiver.PixKey.KeyType, receiver.Bank, receiver.Office, receiver.AccountNumber, receiver.UpdatedAt, receiver.Id)
	if err != nil {
		return internal_error.NewInternalServerError("error updating receiver")
	}

	return nil
}

func (r *ReceiverRepository) DeleteManyReceivers(ctx context.Context, ids []pkg_entity.ID) *internal_error.InternalError {
	_, err := r.Db.ExecContext(ctx, "DELETE FROM receivers WHERE id = ANY($1)", ids)
	if err != nil {
		return internal_error.NewInternalServerError("error deleting receivers")
	}

	return nil
}
