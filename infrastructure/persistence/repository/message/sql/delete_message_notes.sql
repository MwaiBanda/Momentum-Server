DELETE FROM "Note"
WHERE "passageId" IN (
    SELECT id
    FROM "Passage"
    WHERE "messageId" = $1
);
