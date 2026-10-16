UPDATE "VolunteeredMeal"
SET
    description = $2,
    notes = $3
WHERE id = $1
RETURNING id, description, notes, "userId" AS user_id, "mealId" AS meal_id, "createdOn" AS created_on, date;
