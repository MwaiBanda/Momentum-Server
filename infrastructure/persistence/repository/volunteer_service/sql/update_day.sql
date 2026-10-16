UPDATE "VolunteeredServiceDay"
SET notes = $2,
    "userId" = $3
WHERE id = $1
RETURNING id, notes, "userId" AS user_id, date, "createdOn" AS created_on;
