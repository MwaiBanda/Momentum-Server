DELETE FROM "User"
WHERE id = $1
RETURNING id, email, fullname, phone, "createdOn";
