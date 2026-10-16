INSERT INTO "User" (id, email, fullname, phone)
VALUES ($1, $2, $3, $4)
RETURNING id, email, fullname, phone, "createdOn";
