package postgres

import (
	"context"

	"github.com/felipemagrassi/pix-api/configuration/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitializeDatabase(ctx context.Context, databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", databaseURL)
	if err != nil {
		logger.Error("error connecting to database", err)
		return nil, err
	}

	m, err := migrate.New("file://db/migrations", databaseURL)
	if err != nil {
		logger.Error("error creating migration", err)
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("error running migration", err)
		return nil, err
	}

	return db, nil
}
