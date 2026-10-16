package meal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Update(ctx context.Context, params domainmeal.UpdateParams) (domainmeal.Meal, error) {
	query := persistence.LoadFile("repository/meal/sql/update.sql")
	result, err := repository.db.ExecContext(ctx, query,
		params.ID,
		params.Reason,
		params.Email,
		params.Phone,
		params.City,
		params.NumOfAdults,
		params.NumberOfKids,
		params.Street,
		params.PreferredTime,
		params.Favourites,
		params.LeastFavourites,
		params.Allergies,
		params.Instructions,
		params.Recipient,
	)
	if err != nil {
		return domainmeal.Meal{}, fmt.Errorf("update meal %q: %w", params.ID, err)
	}
	if count, err := result.RowsAffected(); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("check updated meal %q: %w", params.ID, err)
	} else if count == 0 {
		return domainmeal.Meal{}, fmt.Errorf("update meal %q: %w", params.ID, domainmeal.ErrNotFound)
	}

	meal, err := repository.getByID(ctx, params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainmeal.Meal{}, fmt.Errorf("get updated meal %q: %w", params.ID, domainmeal.ErrNotFound)
		}
		return domainmeal.Meal{}, err
	}
	return meal, nil
}
