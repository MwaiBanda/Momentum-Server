INSERT INTO "User" (id, email, fullname, phone)
VALUES ($1, $2, $3, $4)
ON CONFLICT (id) DO UPDATE
SET email = EXCLUDED.email,
    fullname = EXCLUDED.fullname,
    phone = EXCLUDED.phone
RETURNING id, email, fullname, phone, "createdOn";
