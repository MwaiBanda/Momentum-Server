package volunteer_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	domainservice "Momentum/domain/volunteer_service"
	"Momentum/infrastructure/persistence"
)

type updatedDayRow struct {
	ID        string         `db:"id"`
	Notes     string         `db:"notes"`
	UserID    sql.NullString `db:"user_id"`
	Date      string         `db:"date"`
	CreatedOn time.Time      `db:"created_on"`
}

func (repository *Repository) UpdateDay(ctx context.Context, params domainservice.UpdateDayParams) (domainservice.VolunteerServiceDay, error) {
	var row updatedDayRow
	query := persistence.LoadFile("repository/volunteer_service/sql/update_day.sql")
	if err := repository.db.GetContext(
		ctx,
		&row,
		query,
		params.ID,
		params.Notes,
		nullableString(params.UserID),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainservice.VolunteerServiceDay{}, fmt.Errorf("update volunteer service day %q: %w", params.ID, domainservice.ErrDayNotFound)
		}
		return domainservice.VolunteerServiceDay{}, fmt.Errorf("update volunteer service day %q: %w", params.ID, err)
	}

	return domainservice.VolunteerServiceDay{
		ID:        row.ID,
		Notes:     row.Notes,
		UserID:    nullStringValue(row.UserID),
		Date:      row.Date,
		CreatedOn: row.CreatedOn,
	}, nil
}
