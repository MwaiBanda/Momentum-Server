package message

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	domainmessage "Momentum/domain/message"
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

func TestRepositoryIntegrationMessageAggregate(t *testing.T) {
	database := openIntegrationDatabase(t)
	repository := GetInstance(database)
	operationCtx, cancelOperation := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelOperation()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	messageID := "repository-test-message-" + suffix
	headerPassageID := "repository-test-header-" + suffix
	bodyPassageID := "repository-test-body-" + suffix
	noteID := "repository-test-note-" + suffix
	userID := "repository-test-user-" + suffix

	t.Cleanup(func() {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCleanup()

		queries := []struct {
			query string
			args  []any
		}{
			{
				query: `DELETE FROM "Note" WHERE "passageId" IN ($1, $2)`,
				args:  []any{headerPassageID, bodyPassageID},
			},
			{
				query: `DELETE FROM "Passage" WHERE id IN ($1, $2)`,
				args:  []any{headerPassageID, bodyPassageID},
			},
			{
				query: `DELETE FROM "Message" WHERE id = $1`,
				args:  []any{messageID},
			},
			{
				query: `DELETE FROM "User" WHERE id = $1`,
				args:  []any{userID},
			},
		}
		for index, cleanup := range queries {
			if _, err := database.ExecContext(cleanupCtx, cleanup.query, cleanup.args...); err != nil {
				t.Errorf("cleanup query %d: %v", index, err)
			}
		}
	})

	if _, err := database.ExecContext(
		operationCtx,
		`INSERT INTO "User" (id, email, fullname, phone) VALUES ($1, $2, $3, $4)`,
		userID,
		"message-test@example.com",
		"Message Test User",
		"555-0100",
	); err != nil {
		t.Fatalf("insert note user: %v", err)
	}

	header := "Introduction"
	verse := "John 1:1"
	body := "In the beginning"
	if err := repository.Create(operationCtx, domainmessage.CreateParams{
		ID:        messageID,
		Thumbnail: "thumbnail.png",
		Series:    "Series",
		Title:     "Repository Test Message",
		Preacher:  "Repository Test Preacher",
		Date:      "2026-10-16",
		Passages: []domainmessage.CreatePassageParams{
			{ID: headerPassageID, Header: &header, Order: 1},
			{ID: bodyPassageID, Verse: &verse, Message: &body, Order: 2},
		},
	}); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	unpublished, err := repository.ListUnpublished(operationCtx)
	if err != nil {
		t.Fatalf("ListUnpublished() error = %v", err)
	}
	if len(unpublished) != 1 || len(unpublished[0].Passages) != 2 {
		t.Fatalf("ListUnpublished() = %#v, want one message with two passages", unpublished)
	}

	if err := repository.Update(operationCtx, domainmessage.UpdateParams{
		ID:        messageID,
		Thumbnail: "updated-thumbnail.png",
		Series:    "Updated Series",
		Title:     "Updated Message",
		Preacher:  "Updated Preacher",
		Date:      "2026-10-17",
		Published: true,
		Passages: []domainmessage.UpdatePassageParams{{
			ID:      bodyPassageID,
			Header:  "",
			Verse:   "John 1:2",
			Message: "The word was with God",
		}},
	}); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if err := repository.CreateNote(operationCtx, domainmessage.CreateNoteParams{
		ID:        noteID,
		Content:   "Remember this",
		UserID:    userID,
		PassageID: bodyPassageID,
	}); err != nil {
		t.Fatalf("CreateNote() error = %v", err)
	}

	published, err := repository.ListPublishedByUser(operationCtx, userID)
	if err != nil {
		t.Fatalf("ListPublishedByUser() error = %v", err)
	}
	if len(published) != 1 || len(published[0].Passages) != 2 {
		t.Fatalf("ListPublishedByUser() = %#v, want one message with two passages", published)
	}
	var noteCount int
	for _, passage := range published[0].Passages {
		noteCount += len(passage.Notes)
	}
	if noteCount != 1 || published[0].Passages[1].Notes[0].Content != "Remember this" {
		t.Fatalf("published notes = %#v, want user's note", published[0].Passages)
	}

	updatedNote, err := repository.UpdateNote(operationCtx, domainmessage.UpdateNoteParams{
		ID:      noteID,
		Content: "Updated note",
	})
	if err != nil {
		t.Fatalf("UpdateNote() error = %v", err)
	}
	if updatedNote.Content != "Updated note" || updatedNote.UserID != userID {
		t.Fatalf("UpdateNote() = %#v, want updated note", updatedNote)
	}

	deletedNote, err := repository.DeleteNote(operationCtx, noteID)
	if err != nil {
		t.Fatalf("DeleteNote() error = %v", err)
	}
	if deletedNote.ID != noteID || deletedNote.UserID != userID {
		t.Fatalf("DeleteNote() = %#v, want deleted note", deletedNote)
	}

	if err := repository.CreateNote(operationCtx, domainmessage.CreateNoteParams{
		ID:        noteID,
		Content:   "Delete with message",
		UserID:    userID,
		PassageID: bodyPassageID,
	}); err != nil {
		t.Fatalf("recreate note: %v", err)
	}

	if err := repository.Delete(operationCtx, messageID); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	var remaining int
	for index, query := range []string{
		`SELECT count(*) FROM "Message" WHERE id = $1`,
		`SELECT count(*) FROM "Passage" WHERE id IN ($1, $2)`,
		`SELECT count(*) FROM "Note" WHERE "passageId" IN ($1, $2)`,
	} {
		var err error
		switch index {
		case 0:
			err = database.GetContext(operationCtx, &remaining, query, messageID)
		default:
			err = database.GetContext(operationCtx, &remaining, query, headerPassageID, bodyPassageID)
		}
		if err != nil {
			t.Fatalf("count remaining records %d: %v", index, err)
		}
		if remaining != 0 {
			t.Fatalf("remaining records for query %d = %d, want 0", index, remaining)
		}
	}
}

func TestRepositoryIntegrationCreateRollsBackChildren(t *testing.T) {
	database := openIntegrationDatabase(t)
	repository := GetInstance(database)
	operationCtx, cancelOperation := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelOperation()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	messageID := "repository-test-rollback-message-" + suffix
	passageID := "repository-test-rollback-passage-" + suffix
	t.Cleanup(func() {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCleanup()
		_, _ = database.ExecContext(cleanupCtx, `DELETE FROM "Passage" WHERE id = $1`, passageID)
		_, _ = database.ExecContext(cleanupCtx, `DELETE FROM "Message" WHERE id = $1`, messageID)
	})

	header := "Header"
	err := repository.Create(operationCtx, domainmessage.CreateParams{
		ID:    messageID,
		Title: "Rollback",
		Passages: []domainmessage.CreatePassageParams{
			{ID: passageID, Header: &header, Order: 1},
			{ID: passageID, Header: &header, Order: 2},
		},
	})
	if err == nil {
		t.Fatal("Create() error = nil, want duplicate child failure")
	}

	var messageCount, passageCount int
	if err := database.GetContext(operationCtx, &messageCount, `SELECT count(*) FROM "Message" WHERE id = $1`, messageID); err != nil {
		t.Fatalf("count rollback message: %v", err)
	}
	if err := database.GetContext(operationCtx, &passageCount, `SELECT count(*) FROM "Passage" WHERE id = $1`, passageID); err != nil {
		t.Fatalf("count rollback passage: %v", err)
	}
	if messageCount != 0 || passageCount != 0 {
		t.Fatalf("rollback counts = message %d, passage %d; want 0, 0", messageCount, passageCount)
	}
}
