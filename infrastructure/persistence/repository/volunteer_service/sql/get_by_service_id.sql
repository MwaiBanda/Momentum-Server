SELECT
    vs.id AS volunteer_id,
    vs.description,
    vs.email,
    vs.phone,
    vs.organizer,
    vs."serviceId" AS volunteer_service_id,
    vs."userId" AS volunteer_user_id,
    vs."createdOn" AS volunteer_created_on,
    day.id AS day_id,
    day.notes AS day_notes,
    day.date AS day_date,
    day."userId" AS day_user_id,
    day."createdOn" AS day_created_on,
    u.id AS user_id,
    u.email AS user_email,
    u.fullname AS user_fullname,
    u.phone AS user_phone,
    u."createdOn" AS user_created_on
FROM "VolunteerService" AS vs
LEFT JOIN "VolunteeredServiceDay" AS day ON day."serviceId" = vs.id
LEFT JOIN "User" AS u ON u.id = day."userId"
WHERE vs."serviceId" = $1
ORDER BY vs."createdOn", vs.id, day.date, day.id;
