package message

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domainmessage "Momentum/domain/message"
	"Momentum/infrastructure/persistence"
)

type noteRow struct {
	ID        string         `db:"id"`
	Content   string         `db:"content"`
	UserID    sql.NullString `db:"user_id"`
	PassageID sql.NullString `db:"passage_id"`
	CreatedOn sql.NullTime   `db:"created_on"`
	UpdatedOn sql.NullTime   `db:"updated_on"`
}

func (repository *Repository) CreateNote(ctx context.Context, params domainmessage.CreateNoteParams) error {
	query := persistence.LoadFile("repository/message/sql/create_note.sql")
	if _, err := repository.db.ExecContext(
		ctx,
		query,
		params.ID,
		params.Content,
		params.UserID,
		params.PassageID,
	); err != nil {
		return fmt.Errorf("create note %q: %w", params.ID, err)
	}
	return nil
}

func (repository *Repository) UpdateNote(ctx context.Context, params domainmessage.UpdateNoteParams) (domainmessage.Note, error) {
	var row noteRow
	query := persistence.LoadFile("repository/message/sql/update_note.sql")
	if err := repository.db.GetContext(ctx, &row, query, params.ID, params.Content); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainmessage.Note{}, fmt.Errorf("update note %q: %w", params.ID, domainmessage.ErrNoteNotFound)
		}
		return domainmessage.Note{}, fmt.Errorf("update note %q: %w", params.ID, err)
	}
	return toNote(row), nil
}

func (repository *Repository) DeleteNote(ctx context.Context, id string) (domainmessage.Note, error) {
	var row noteRow
	query := persistence.LoadFile("repository/message/sql/delete_note.sql")
	if err := repository.db.GetContext(ctx, &row, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainmessage.Note{}, fmt.Errorf("delete note %q: %w", id, domainmessage.ErrNoteNotFound)
		}
		return domainmessage.Note{}, fmt.Errorf("delete note %q: %w", id, err)
	}
	return toNote(row), nil
}

func toNote(row noteRow) domainmessage.Note {
	return domainmessage.Note{
		ID:        row.ID,
		Content:   row.Content,
		UserID:    nullStringValue(row.UserID),
		PassageID: nullStringValue(row.PassageID),
		CreatedOn: nullTimeValue(row.CreatedOn),
		UpdatedOn: nullTimeValue(row.UpdatedOn),
	}
}

func nullStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	return value.String
}
