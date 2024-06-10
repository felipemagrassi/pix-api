package postgres

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/felipemagrassi/pix-api/configuration/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitializeDatabase(ctx context.Context, databaseURL string) (*sqlx.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		logger.Error("error connecting to database", err)
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		logger.Error("error pinging database", err)
		return nil, err
	}

	_, path, _, ok := runtime.Caller(0)
	if !ok {
		logger.Info("error getting migration path")
		return nil, err
	}

	migrationsSource := "file://" + path + "/db/migrations"
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error("error creating migration driver", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsSource, "postgres", driver)
	if err != nil {
		logger.Error("error creating migration", err)
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("error running migration", err)
		return nil, err
	}

	return sqlx.NewDb(db, "postgres"), nil
}
