UPDATE "Meal"
SET
    reason = $2,
    email = $3,
    phone = $4,
    city = $5,
    "numOfAdults" = $6,
    "numberOfKids" = $7,
    street = $8,
    "preferredTime" = $9,
    favourites = $10,
    "leastFavourites" = $11,
    allergies = $12,
    instructions = $13,
    recipient = $14
WHERE id = $1;
