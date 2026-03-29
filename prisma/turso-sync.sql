PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Message" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "thumbnail" TEXT NOT NULL,
    "series" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "preacher" TEXT NOT NULL,
    "date" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "published" BOOLEAN NOT NULL DEFAULT false,
    "hasOrder" BOOLEAN NOT NULL DEFAULT false
);
CREATE TABLE IF NOT EXISTS "Passage" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "header" TEXT,
    "verse" TEXT,
    "message" TEXT,
    "messageId" TEXT,
    "createdOn" DATETIME DEFAULT CURRENT_TIMESTAMP,
    "order" INTEGER DEFAULT 0,
    CONSTRAINT "Passage_messageId_fkey" FOREIGN KEY ("messageId") REFERENCES "Message" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "Service" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "Image" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS "Note" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "content" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "userId" TEXT,
    "passageId" TEXT,
    CONSTRAINT "Note_passageId_fkey" FOREIGN KEY ("passageId") REFERENCES "Passage" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "Note_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "User" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "email" TEXT NOT NULL,
    "fullname" TEXT NOT NULL,
    "phone" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS "VolunteerService" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "description" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "phone" TEXT NOT NULL,
    "organizer" TEXT NOT NULL,
    "serviceId" TEXT,
    "userId" TEXT,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "VolunteerService_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "VolunteerService_serviceId_fkey" FOREIGN KEY ("serviceId") REFERENCES "Service" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "VolunteeredServiceDay" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "notes" TEXT NOT NULL,
    "userId" TEXT,
    "serviceId" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "date" TEXT NOT NULL DEFAULT '',
    CONSTRAINT "VolunteeredServiceDay_serviceId_fkey" FOREIGN KEY ("serviceId") REFERENCES "VolunteerService" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "VolunteeredServiceDay_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "VolunteeredServiceParticipant" (
    "userId" TEXT NOT NULL,
    "serviceId" TEXT NOT NULL,
    "assignedAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "id" TEXT NOT NULL,

    PRIMARY KEY ("id", "userId", "serviceId"),
    CONSTRAINT "VolunteeredServiceParticipant_serviceId_fkey" FOREIGN KEY ("serviceId") REFERENCES "VolunteerService" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "VolunteeredServiceParticipant_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "Meal" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "reason" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "phone" TEXT NOT NULL,
    "city" TEXT NOT NULL,
    "numOfAdults" INTEGER NOT NULL,
    "numberOfKids" INTEGER NOT NULL,
    "street" TEXT NOT NULL,
    "preferredTime" TEXT NOT NULL,
    "favourites" TEXT NOT NULL,
    "leastFavourites" TEXT NOT NULL,
    "allergies" TEXT NOT NULL,
    "instructions" TEXT NOT NULL,
    "userId" TEXT,
    "recipient" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "Meal_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "VolunteeredMeal" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "description" TEXT NOT NULL,
    "notes" TEXT NOT NULL,
    "userId" TEXT,
    "mealId" TEXT,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "date" TEXT NOT NULL DEFAULT '',
    CONSTRAINT "VolunteeredMeal_mealId_fkey" FOREIGN KEY ("mealId") REFERENCES "Meal" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "VolunteeredMeal_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "MealUpdates" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "message" TEXT NOT NULL,
    "mealId" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "MealUpdates_mealId_fkey" FOREIGN KEY ("mealId") REFERENCES "VolunteeredMeal" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "MealParticipant" (
    "userId" TEXT NOT NULL,
    "mealId" TEXT NOT NULL,
    "assignedAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "id" TEXT NOT NULL,

    PRIMARY KEY ("id", "userId", "mealId"),
    CONSTRAINT "MealParticipant_mealId_fkey" FOREIGN KEY ("mealId") REFERENCES "Meal" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "MealParticipant_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "Transaction" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "amount" INTEGER NOT NULL,
    "date" TEXT NOT NULL,
    "createdOn" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "userId" TEXT NOT NULL,
    "Description" TEXT NOT NULL,
    CONSTRAINT "Transaction_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE INDEX "Passage_messageId_idx" ON "Passage"("messageId");
CREATE UNIQUE INDEX "Service_type_key" ON "Service"("type");
CREATE INDEX "Note_userId_passageId_idx" ON "Note"("userId", "passageId");
CREATE INDEX "VolunteerService_userId_serviceId_idx" ON "VolunteerService"("userId", "serviceId");
CREATE INDEX "VolunteeredServiceDay_userId_serviceId_idx" ON "VolunteeredServiceDay"("userId", "serviceId");
CREATE INDEX "Meal_userId_idx" ON "Meal"("userId");
CREATE INDEX "VolunteeredMeal_userId_mealId_idx" ON "VolunteeredMeal"("userId", "mealId");
CREATE INDEX "MealUpdates_mealId_idx" ON "MealUpdates"("mealId");
CREATE INDEX "Transaction_userId_idx" ON "Transaction"("userId");
COMMIT;
