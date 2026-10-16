package meal

import (
	"context"
	"fmt"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Create(ctx context.Context, params domainmeal.CreateParams) (domainmeal.Meal, error) {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return domainmeal.Meal{}, fmt.Errorf("begin meal creation: %w", err)
	}
	defer tx.Rollback()

	createQuery := persistence.LoadFile("repository/meal/sql/create.sql")
	if _, err := tx.ExecContext(ctx, createQuery,
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
		params.UserID,
		params.Recipient,
	); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("create meal %q: %w", params.ID, err)
	}

	volunteeredMealQuery := persistence.LoadFile("repository/meal/sql/create_volunteered_meal.sql")
	for _, volunteeredMeal := range params.Meals {
		if _, err := tx.ExecContext(ctx, volunteeredMealQuery, volunteeredMeal.ID, params.ID, volunteeredMeal.Date); err != nil {
			return domainmeal.Meal{}, fmt.Errorf("create volunteered meal %q: %w", volunteeredMeal.ID, err)
		}
	}

	var rows []aggregateRow
	getQuery := persistence.LoadFile("repository/meal/sql/get_by_id.sql")
	if err := tx.SelectContext(ctx, &rows, getQuery, params.ID); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("get created meal %q: %w", params.ID, err)
	}

	if err := tx.Commit(); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("commit meal %q: %w", params.ID, err)
	}

	meals := aggregateRows(rows)
	if len(meals) == 0 {
		return domainmeal.Meal{}, fmt.Errorf("get created meal %q: %w", params.ID, domainmeal.ErrNotFound)
	}
	return meals[0], nil
}
