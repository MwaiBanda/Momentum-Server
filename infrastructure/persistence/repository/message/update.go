package message

import (
	"context"
	"fmt"

	domainmessage "Momentum/domain/message"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Update(ctx context.Context, params domainmessage.UpdateParams) error {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin message update: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	messageQuery := persistence.LoadFile("repository/message/sql/update.sql")
	result, err := tx.ExecContext(
		ctx,
		messageQuery,
		params.ID,
		params.Thumbnail,
		params.Series,
		params.Title,
		params.Preacher,
		params.Date,
		params.Published,
	)
	if err != nil {
		return fmt.Errorf("update message %q: %w", params.ID, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("inspect update message %q: %w", params.ID, err)
	}
	if affected == 0 {
		return fmt.Errorf("update message %q: %w", params.ID, domainmessage.ErrMessageNotFound)
	}

	passageQuery := persistence.LoadFile("repository/message/sql/update_passage.sql")
	for _, passage := range params.Passages {
		result, err := tx.ExecContext(
			ctx,
			passageQuery,
			passage.ID,
			passage.Header,
			passage.Verse,
			passage.Message,
		)
		if err != nil {
			return fmt.Errorf("update passage %q: %w", passage.ID, err)
		}

		affected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("inspect update passage %q: %w", passage.ID, err)
		}
		if affected == 0 {
			return fmt.Errorf("update passage %q: %w", passage.ID, domainmessage.ErrPassageNotFound)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit message %q update: %w", params.ID, err)
	}
	committed = true
	return nil
}
