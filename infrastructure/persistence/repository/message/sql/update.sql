UPDATE "Message"
SET title = $4,
    preacher = $5,
    date = $6,
    thumbnail = $2,
    series = $3,
    published = $7
WHERE id = $1;
