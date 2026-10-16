UPDATE "VolunteeredMeal"
SET
    description = $4,
    notes = $5,
    "userId" = $3
WHERE id = $1
  AND "mealId" = $2;
