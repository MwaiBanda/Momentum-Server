package message

import (
	"context"
	"fmt"

	domainmessage "Momentum/domain/message"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Create(ctx context.Context, params domainmessage.CreateParams) error {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin message creation: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	messageQuery := persistence.LoadFile("repository/message/sql/create.sql")
	if _, err := tx.ExecContext(
		ctx,
		messageQuery,
		params.ID,
		params.Thumbnail,
		params.Series,
		params.Title,
		params.Preacher,
		params.Date,
	); err != nil {
		return fmt.Errorf("create message %q: %w", params.ID, err)
	}

	passageQuery := persistence.LoadFile("repository/message/sql/create_passage.sql")
	for _, passage := range params.Passages {
		if _, err := tx.ExecContext(
			ctx,
			passageQuery,
			passage.ID,
			passage.Header,
			passage.Verse,
			passage.Message,
			params.ID,
			passage.Order,
		); err != nil {
			return fmt.Errorf("create passage %q for message %q: %w", passage.ID, params.ID, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit message %q: %w", params.ID, err)
	}
	committed = true
	return nil
}
