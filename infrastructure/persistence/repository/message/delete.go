package message

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domainmessage "Momentum/domain/message"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Delete(ctx context.Context, id string) error {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin message deletion: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	existsQuery := persistence.LoadFile("repository/message/sql/lock_message.sql")
	var existingID string
	if err := tx.GetContext(ctx, &existingID, existsQuery, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("delete message %q: %w", id, domainmessage.ErrMessageNotFound)
		}
		return fmt.Errorf("find message %q for deletion: %w", id, err)
	}

	deleteNotesQuery := persistence.LoadFile("repository/message/sql/delete_message_notes.sql")
	if _, err := tx.ExecContext(ctx, deleteNotesQuery, id); err != nil {
		return fmt.Errorf("delete notes for message %q: %w", id, err)
	}

	deletePassagesQuery := persistence.LoadFile("repository/message/sql/delete_message_passages.sql")
	if _, err := tx.ExecContext(ctx, deletePassagesQuery, id); err != nil {
		return fmt.Errorf("delete passages for message %q: %w", id, err)
	}

	deleteMessageQuery := persistence.LoadFile("repository/message/sql/delete.sql")
	if _, err := tx.ExecContext(ctx, deleteMessageQuery, id); err != nil {
		return fmt.Errorf("delete message %q: %w", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit message %q deletion: %w", id, err)
	}
	committed = true
	return nil
}
