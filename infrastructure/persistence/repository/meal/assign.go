package meal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) AssignVolunteer(ctx context.Context, params domainmeal.AssignVolunteerParams) (domainmeal.Meal, error) {
	tx, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		return domainmeal.Meal{}, fmt.Errorf("begin meal volunteer assignment: %w", err)
	}
	defer tx.Rollback()

	updateQuery := persistence.LoadFile("repository/meal/sql/assign_volunteered_meal.sql")
	result, err := tx.ExecContext(ctx, updateQuery,
		params.VolunteeredMealID,
		params.MealID,
		params.UserID,
		params.Description,
		params.Notes,
	)
	if err != nil {
		return domainmeal.Meal{}, fmt.Errorf("assign volunteered meal %q: %w", params.VolunteeredMealID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return domainmeal.Meal{}, fmt.Errorf("check volunteered meal assignment %q: %w", params.VolunteeredMealID, err)
	}
	if count == 0 {
		var exists bool
		if err := tx.GetContext(ctx, &exists, `SELECT EXISTS (SELECT 1 FROM "Meal" WHERE id = $1)`, params.MealID); err != nil {
			return domainmeal.Meal{}, fmt.Errorf("check meal %q for volunteer assignment: %w", params.MealID, err)
		}
		if !exists {
			return domainmeal.Meal{}, fmt.Errorf("assign volunteered meal %q: %w", params.MealID, domainmeal.ErrNotFound)
		}
		return domainmeal.Meal{}, fmt.Errorf("assign volunteered meal %q: %w", params.VolunteeredMealID, domainmeal.ErrVolunteeredMealNotFound)
	}

	participantQuery := persistence.LoadFile("repository/meal/sql/upsert_participant.sql")
	if _, err := tx.ExecContext(ctx, participantQuery, params.VolunteeredMealID, params.UserID, params.MealID); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("save meal participant for %q: %w", params.MealID, err)
	}

	var rows []aggregateRow
	getQuery := persistence.LoadFile("repository/meal/sql/get_by_id.sql")
	if err := tx.SelectContext(ctx, &rows, getQuery, params.MealID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainmeal.Meal{}, fmt.Errorf("get assigned meal %q: %w", params.MealID, domainmeal.ErrNotFound)
		}
		return domainmeal.Meal{}, fmt.Errorf("get assigned meal %q: %w", params.MealID, err)
	}

	if err := tx.Commit(); err != nil {
		return domainmeal.Meal{}, fmt.Errorf("commit meal volunteer assignment %q: %w", params.VolunteeredMealID, err)
	}
	meals := aggregateRows(rows)
	if len(meals) == 0 {
		return domainmeal.Meal{}, fmt.Errorf("get assigned meal %q: %w", params.MealID, domainmeal.ErrNotFound)
	}
	return meals[0], nil
}
