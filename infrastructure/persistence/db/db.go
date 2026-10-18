package db

import (
	"context"
	"fmt"
	"time"

	"github.com/XSAM/otelsql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/attribute"
)

const driverName = "pgx"

func Open(ctx context.Context, databaseURL string) (*sqlx.DB, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("database URL is required")
	}

	rawDB, err := otelsql.Open(
		driverName,
		databaseURL,
		otelsql.WithAttributes(
			attribute.String("db.system.name", "postgresql"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	db := sqlx.NewDb(rawDB, driverName)

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

func Close(db *sqlx.DB) error {
	if db == nil {
		return nil
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("close database: %w", err)
	}

	return nil
}
