SELECT
    m.id AS message_id,
    m.thumbnail,
    m.series,
    m.title,
    m.preacher,
    m.date,
    m."createdOn" AS message_created_on,
    m.published,
    m."hasOrder" AS has_order,
    p.id AS passage_id,
    p.header,
    p.verse,
    p.message AS passage_message,
    p."messageId" AS passage_message_id,
    p."createdOn" AS passage_created_on,
    p."order" AS passage_order
FROM "Message" AS m
LEFT JOIN "Passage" AS p ON p."messageId" = m.id
WHERE m.published = false
ORDER BY m."createdOn" DESC, p."order" ASC NULLS LAST;
