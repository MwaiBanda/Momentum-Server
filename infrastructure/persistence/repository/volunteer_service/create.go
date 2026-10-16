package volunteer_service

import (
	"context"
	"fmt"

	domainservice "Momentum/domain/volunteer_service"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Create(ctx context.Context, params domainservice.CreateParams) error {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin volunteer service creation: %w", err)
	}

	parentQuery := persistence.LoadFile("repository/volunteer_service/sql/create.sql")
	if _, err := tx.ExecContext(
		ctx,
		parentQuery,
		params.ID,
		params.Description,
		params.Email,
		params.Phone,
		params.Organizer,
		nullableString(params.ServiceID),
		nullableString(params.UserID),
	); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("create volunteer service %q: %w", params.ID, err)
	}

	dayQuery := persistence.LoadFile("repository/volunteer_service/sql/create_day.sql")
	for _, day := range params.Days {
		if _, err := tx.ExecContext(ctx, dayQuery, day.ID, "", day.Date, params.ID); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("create volunteer service day %q: %w", day.ID, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit volunteer service %q: %w", params.ID, err)
	}

	return nil
}

func nullableString(value string) any {
	if value == "" {
		return nil
	}
	return value
}
