SELECT
    t.id,
    t.amount,
    t.date,
    t."createdOn" AS created_on,
    t."Description" AS description,
    u.id AS user_id,
    u.email AS user_email,
    u.fullname AS user_fullname,
    u.phone AS user_phone,
    u."createdOn" AS user_created_on
FROM "Transaction" t
JOIN "User" u ON u.id = t."userId"
WHERE t."userId" = $1
ORDER BY t."createdOn" DESC;
