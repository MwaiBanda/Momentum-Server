package db

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestMigrationIntegrationReleasesAdvisoryLock(t *testing.T) {
	databaseURL := os.Getenv("MOMENTUM_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("MOMENTUM_TEST_DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	database, err := Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open database = %v", err)
	}
	defer database.Close()

	if err := Migrate(ctx, database); err != nil {
		t.Fatalf("migrate database = %v", err)
	}

	verificationDatabase, err := Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open verification database = %v", err)
	}
	defer verificationDatabase.Close()

	var acquired bool
	if err := verificationDatabase.GetContext(
		ctx,
		&acquired,
		`SELECT pg_try_advisory_lock(hashtext('momentum-server-goose-migrations'))`,
	); err != nil {
		t.Fatalf("check migration lock = %v", err)
	}
	if !acquired {
		t.Fatal("migration advisory lock is still held after migration")
	}

	if _, err := verificationDatabase.ExecContext(
		ctx,
		`SELECT pg_advisory_unlock(hashtext('momentum-server-goose-migrations'))`,
	); err != nil {
		t.Fatalf("release verification lock = %v", err)
	}
}
