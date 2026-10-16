package user

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	domainuser "Momentum/domain/user"
	"Momentum/infrastructure/persistence/db"

	"github.com/jmoiron/sqlx"
)

func openIntegrationDatabase(t *testing.T) *sqlx.DB {
	t.Helper()

	databaseURL := os.Getenv("MOMENTUM_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("MOMENTUM_TEST_DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	database, err := db.Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open integration database: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(database); err != nil {
			t.Errorf("close integration database: %v", err)
		}
	})

	if err := db.Migrate(ctx, database); err != nil {
		t.Fatalf("migrate integration database: %v", err)
	}

	return database
}

func TestRepositoryIntegrationCRUD(t *testing.T) {
	database := openIntegrationDatabase(t)
	repository := GetInstance(database)
	userID := fmt.Sprintf("repository-test-%d", time.Now().UnixNano())
	operationCtx, cancelOperation := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelOperation()
	upsertID := ""
	t.Cleanup(func() {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCleanup()

		for _, id := range []string{userID, upsertID} {
			if id == "" {
				continue
			}
			if _, err := database.ExecContext(cleanupCtx, `DELETE FROM "User" WHERE id = $1`, id); err != nil {
				t.Errorf("clean up integration user %q: %v", id, err)
			}
		}
	})

	createParams := domainuser.CreateParams{
		ID:       userID,
		Email:    "repository-test@example.com",
		Fullname: "Repository Test",
		Phone:    "555-0100",
	}

	created, err := repository.Create(operationCtx, createParams)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if created.ID != createParams.ID ||
		created.Email != createParams.Email ||
		created.Fullname != createParams.Fullname ||
		created.Phone != createParams.Phone {
		t.Fatalf("Create() = %#v, want fields from %#v", created, createParams)
	}
	if created.CreatedOn.IsZero() {
		t.Fatal("Create() returned zero CreatedOn")
	}

	fetched, err := repository.GetByID(operationCtx, userID)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if fetched != created {
		t.Fatalf("GetByID() = %#v, want %#v", fetched, created)
	}

	updateParams := domainuser.UpdateParams{
		ID:       userID,
		Email:    "updated-repository-test@example.com",
		Fullname: "Updated Repository Test",
		Phone:    "555-0101",
	}
	updated, err := repository.Update(operationCtx, updateParams)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	if updated.ID != updateParams.ID ||
		updated.Email != updateParams.Email ||
		updated.Fullname != updateParams.Fullname ||
		updated.Phone != updateParams.Phone {
		t.Fatalf("Update() = %#v, want fields from %#v", updated, updateParams)
	}
	if updated.CreatedOn != created.CreatedOn {
		t.Fatalf("Update() changed CreatedOn from %v to %v", created.CreatedOn, updated.CreatedOn)
	}

	upsertID = userID + "-upsert"
	upserted, err := repository.Update(operationCtx, domainuser.UpdateParams{
		ID:       upsertID,
		Email:    "upsert-repository-test@example.com",
		Fullname: "Upsert Repository Test",
		Phone:    "555-0102",
	})
	if err != nil {
		t.Fatalf("Update() upsert error = %v", err)
	}
	if upserted.ID != upsertID {
		t.Fatalf("Update() upsert ID = %q, want %q", upserted.ID, upsertID)
	}

	deleted, err := repository.Delete(operationCtx, userID)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	if deleted.ID != userID || deleted.Email != updateParams.Email {
		t.Fatalf("Delete() = %#v, want deleted updated user", deleted)
	}

	_, err = repository.GetByID(operationCtx, userID)
	if !errors.Is(err, domainuser.ErrNotFound) {
		t.Fatalf("GetByID() after delete error = %v, want %v", err, domainuser.ErrNotFound)
	}

	_, err = repository.Delete(operationCtx, userID)
	if !errors.Is(err, domainuser.ErrNotFound) {
		t.Fatalf("Delete() missing error = %v, want %v", err, domainuser.ErrNotFound)
	}
}
