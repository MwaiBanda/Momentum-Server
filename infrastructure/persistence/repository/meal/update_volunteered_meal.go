package meal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence"
)

func (repository *Repository) UpdateVolunteeredMeal(ctx context.Context, params domainmeal.UpdateVolunteeredMealParams) (domainmeal.VolunteeredMeal, error) {
	var row volunteeredMealRow
	query := persistence.LoadFile("repository/meal/sql/update_volunteered_meal.sql")
	if err := repository.db.GetContext(ctx, &row, query, params.ID, params.Description, params.Notes); err != nil {
		if err == sql.ErrNoRows {
			return domainmeal.VolunteeredMeal{}, fmt.Errorf("update volunteered meal %q: %w", params.ID, domainmeal.ErrVolunteeredMealNotFound)
		}
		return domainmeal.VolunteeredMeal{}, fmt.Errorf("update volunteered meal %q: %w", params.ID, err)
	}
	return row.entity(), nil
}

type volunteeredMealRow struct {
	ID          string         `db:"id"`
	Description string         `db:"description"`
	Notes       string         `db:"notes"`
	UserID      sql.NullString `db:"user_id"`
	MealID      sql.NullString `db:"meal_id"`
	CreatedOn   time.Time      `db:"created_on"`
	Date        string         `db:"date"`
}

func (row volunteeredMealRow) entity() domainmeal.VolunteeredMeal {
	return domainmeal.VolunteeredMeal{
		ID:          row.ID,
		Description: row.Description,
		Notes:       row.Notes,
		UserID:      nullStringValue(row.UserID),
		MealID:      nullStringValue(row.MealID),
		CreatedOn:   row.CreatedOn,
		Date:        row.Date,
		Updates:     make([]domainmeal.Update, 0),
	}
}
