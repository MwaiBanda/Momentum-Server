INSERT INTO "MealParticipant" (
    id,
    "userId",
    "mealId"
) VALUES ($1, $2, $3)
ON CONFLICT (id, "userId", "mealId")
DO UPDATE SET "assignedAt" = CURRENT_TIMESTAMP;
