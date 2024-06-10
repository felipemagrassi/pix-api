package database

import "database/sql"

type ReceiverRepository struct {
	Db *sql.DB
}

func NewReceiverRepository(db *sql.DB) *ReceiverRepository {
	return &ReceiverRepository{Db: db}
}
