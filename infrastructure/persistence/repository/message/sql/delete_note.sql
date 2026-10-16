DELETE FROM "Note"
WHERE id = $1
RETURNING id, content, "userId" AS user_id, "passageId" AS passage_id,
          "createdOn" AS created_on, "updatedOn" AS updated_on;
