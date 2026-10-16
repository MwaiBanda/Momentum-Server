INSERT INTO "Message" (
    id,
    thumbnail,
    series,
    title,
    preacher,
    date,
    "hasOrder"
)
VALUES ($1, $2, $3, $4, $5, $6, true);
