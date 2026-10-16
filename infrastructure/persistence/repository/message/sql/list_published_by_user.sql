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
    p."order" AS passage_order,
    n.id AS note_id,
    n.content AS note_content,
    n."userId" AS note_user_id,
    n."passageId" AS note_passage_id,
    n."createdOn" AS note_created_on,
    n."updatedOn" AS note_updated_on
FROM "Message" AS m
LEFT JOIN "Passage" AS p ON p."messageId" = m.id
LEFT JOIN "Note" AS n ON n."passageId" = p.id AND n."userId" = $1
WHERE m.published = true
ORDER BY m."createdOn" DESC, p."createdOn" ASC NULLS LAST, p."order" ASC NULLS LAST, n."createdOn" ASC NULLS LAST;
