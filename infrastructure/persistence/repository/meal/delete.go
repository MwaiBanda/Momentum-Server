package meal

import (
	"context"
	"fmt"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) Delete(ctx context.Context, id string) (domainmeal.Meal, error) {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return domainmeal.Meal{}, fmt.Errorf("begin meal deletion: %w", err)
	}
	defer tx.Rollback()

	var rows []aggregateRow
	getQuery := persistence.LoadFile("repository/meal/sql/get_by_id.sql")
	if err := tx.SelectContext(ctx, &rows, getQuery, id); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("get meal %q before deletion: %w", id, err)
	}
	meals := aggregateRows(rows)
	if len(meals) == 0 {
		return domainmeal.Meal{}, fmt.Errorf("delete meal %q: %w", id, domainmeal.ErrNotFound)
	}

	deleteQuery := persistence.LoadFile("repository/meal/sql/delete.sql")
	if _, err := tx.ExecContext(ctx, deleteQuery, id); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("delete meal %q: %w", id, err)
	}
	if err := tx.Commit(); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("commit meal deletion %q: %w", id, err)
	}
	return meals[0], nil
}
