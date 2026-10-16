SELECT id, name, type, "Image" AS image, "createdOn" AS createdon
FROM "Service"
WHERE type = $1;
