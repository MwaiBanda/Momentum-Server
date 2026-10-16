package transaction

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	domaintransaction "Momentum/domain/transaction"
	database "Momentum/infrastructure/persistence/db"

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

	connection, err := database.Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open integration database: %v", err)
	}
	t.Cleanup(func() {
		if err := database.Close(connection); err != nil {
			t.Errorf("close integration database: %v", err)
		}
	})

	if err := database.Migrate(ctx, connection); err != nil {
		t.Fatalf("migrate integration database: %v", err)
	}

	return connection
}

func TestRepositoryIntegrationCRUD(t *testing.T) {
	connection := openIntegrationDatabase(t)
	repository := GetInstance(connection)
	userID := fmt.Sprintf("transaction-user-%d", time.Now().UnixNano())
	transactionID := fmt.Sprintf("transaction-%d", time.Now().UnixNano())
	operationCtx, cancelOperation := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelOperation()
	t.Cleanup(func() {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCleanup()
		_, _ = connection.ExecContext(cleanupCtx, `DELETE FROM "Transaction" WHERE id = $1`, transactionID)
		_, _ = connection.ExecContext(cleanupCtx, `DELETE FROM "User" WHERE id = $1`, userID)
	})

	created, err := repository.Create(operationCtx, domaintransaction.CreateParams{
		ID:          transactionID,
		Amount:      125,
		Date:        "2026-10-16",
		Description: "Meals, Outreach",
		UserID:      userID,
		Fullname:    "Transaction Test",
		Email:       "transaction-test@example.com",
		Phone:       "555-0123",
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if created.ID != transactionID || created.User.ID != userID || created.Amount != 125 {
		t.Fatalf("Create() = %#v, want transaction and user fields", created)
	}
	if created.CreatedOn.IsZero() || created.User.CreatedOn.IsZero() {
		t.Fatal("Create() returned zero timestamps")
	}

	all, err := repository.List(operationCtx)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if !containsTransaction(all, transactionID) {
		t.Fatalf("List() did not return created transaction: %#v", all)
	}

	byUser, err := repository.ListByUser(operationCtx, userID)
	if err != nil {
		t.Fatalf("ListByUser() error = %v", err)
	}
	if len(byUser) != 1 || byUser[0].ID != transactionID {
		t.Fatalf("ListByUser() = %#v, want created transaction", byUser)
	}

	deleted, err := repository.Delete(operationCtx, transactionID)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	if deleted.ID != transactionID {
		t.Fatalf("Delete() = %#v, want deleted transaction", deleted)
	}

	_, err = repository.Delete(operationCtx, transactionID)
	if !errors.Is(err, domaintransaction.ErrNotFound) {
		t.Fatalf("Delete() missing error = %v, want %v", err, domaintransaction.ErrNotFound)
	}
}

func containsTransaction(transactions []domaintransaction.Transaction, id string) bool {
	for _, transaction := range transactions {
		if transaction.ID == id {
			return true
		}
	}
	return false
}
