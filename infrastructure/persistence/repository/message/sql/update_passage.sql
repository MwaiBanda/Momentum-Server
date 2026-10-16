UPDATE "Passage"
SET header = $2,
    verse = $3,
    message = $4
WHERE id = $1;
