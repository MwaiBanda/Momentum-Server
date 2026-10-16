package meal

import (
	"context"
	"fmt"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) List(ctx context.Context) ([]domainmeal.Meal, error) {
	var rows []aggregateRow
	query := persistence.LoadFile("repository/meal/sql/list.sql")
	if err := repository.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, fmt.Errorf("list meals: %w", err)
	}
	return aggregateRows(rows), nil
}

func (repository *Repository) getByID(ctx context.Context, id string) (domainmeal.Meal, error) {
	var rows []aggregateRow
	query := persistence.LoadFile("repository/meal/sql/get_by_id.sql")
	if err := repository.db.SelectContext(ctx, &rows, query, id); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("get meal %q: %w", id, err)
	}
	meals := aggregateRows(rows)
	if len(meals) == 0 {
		return domainmeal.Meal{}, fmt.Errorf("get meal %q: %w", id, domainmeal.ErrNotFound)
	}
	return meals[0], nil
}
