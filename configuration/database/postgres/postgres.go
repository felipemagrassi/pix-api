package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitializeDatabase(ctx context.Context, databaseURL string, migrationPath string) (*sqlx.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		slog.Error("error connecting to database", err)
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		slog.Error("error connecting to database", err)
		return nil, err
	}

	_, path, _, ok := runtime.Caller(0)
	if !ok {
		slog.Error("error getting current path")
		return nil, err
	}

	migrationsSource := fmt.Sprintf("file://%s", filepath.Join(path, migrationPath))
	migrationFolder := "db/migrations"

	migrationsSource = fmt.Sprintf("%s/%s", migrationsSource, migrationFolder)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("error creating migration driver", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsSource, "postgres", driver)
	if err != nil {
		slog.Error("error creating migration", err)
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("error running migration", err)
		return nil, err
	}

	return sqlx.NewDb(db, "postgres"), nil
}
