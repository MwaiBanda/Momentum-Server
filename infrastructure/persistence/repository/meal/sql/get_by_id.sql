SELECT
    m.id AS meal_id,
    m.recipient,
    m.reason,
    m.email,
    m.phone,
    m.city,
    m."numOfAdults" AS num_of_adults,
    m."numberOfKids" AS number_of_kids,
    m.street,
    m."preferredTime" AS preferred_time,
    m.favourites,
    m."leastFavourites" AS least_favourites,
    m.allergies,
    m.instructions,
    m."userId" AS meal_user_id,
    m."createdOn" AS created_on,
    meal_user.email AS meal_user_email,
    meal_user.fullname AS meal_user_fullname,
    meal_user.phone AS meal_user_phone,
    meal_user."createdOn" AS meal_user_created_on,
    participant.id AS participant_id,
    participant."userId" AS participant_user_id,
    participant."assignedAt" AS participant_assigned_at,
    participant_user.email AS participant_email,
    participant_user.fullname AS participant_fullname,
    participant_user.phone AS participant_phone,
    participant_user."createdOn" AS participant_created_on,
    volunteered.id AS volunteer_id,
    volunteered.description AS volunteer_description,
    volunteered.notes AS volunteer_notes,
    volunteered."userId" AS volunteer_user_id,
    volunteered."mealId" AS volunteer_meal_id,
    volunteered."createdOn" AS volunteer_created_on,
    volunteered.date AS volunteer_date,
    volunteer_user.email AS volunteer_email,
    volunteer_user.fullname AS volunteer_fullname,
    volunteer_user.phone AS volunteer_phone,
    volunteer_user."createdOn" AS volunteer_user_created_on,
    updates.id AS update_id,
    updates.message AS update_message,
    updates."createdOn" AS update_created_on
FROM "Meal" AS m
LEFT JOIN "User" AS meal_user ON meal_user.id = m."userId"
LEFT JOIN "MealParticipant" AS participant ON participant."mealId" = m.id
LEFT JOIN "User" AS participant_user ON participant_user.id = participant."userId"
LEFT JOIN "VolunteeredMeal" AS volunteered ON volunteered."mealId" = m.id
LEFT JOIN "User" AS volunteer_user ON volunteer_user.id = volunteered."userId"
LEFT JOIN "MealUpdates" AS updates ON updates."mealId" = volunteered.id
WHERE m.id = $1
ORDER BY m.id, participant.id, volunteered.date, volunteered.id, updates."createdOn", updates.id;
