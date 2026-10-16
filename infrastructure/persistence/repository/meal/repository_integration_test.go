package meal

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	domainmeal "Momentum/domain/meal"
	"Momentum/infrastructure/persistence/db"

	"github.com/jmoiron/sqlx"
)

func openIntegrationDatabase(t *testing.T) *sqlx.DB {
	t.Helper()

	databaseURL := os.Getenv("MOMENTUM_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("MOMENTUM_TEST_DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	database, err := db.Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open integration database: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(database); err != nil {
			t.Errorf("close integration database: %v", err)
		}
	})

	if err := db.Migrate(ctx, database); err != nil {
		t.Fatalf("migrate integration database: %v", err)
	}

	return database
}

func TestRepositoryIntegrationCRUDAndAssignments(t *testing.T) {
	database := openIntegrationDatabase(t)
	repository := GetInstance(database)
	operationCtx, cancelOperation := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelOperation()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	mealID := "repository-test-meal-" + suffix
	deletableMealID := "repository-test-deletable-meal-" + suffix
	firstVolunteeredMealID := "repository-test-volunteered-meal-1-" + suffix
	secondVolunteeredMealID := "repository-test-volunteered-meal-2-" + suffix
	organizerID := "repository-test-organizer-" + suffix
	volunteerID := "repository-test-volunteer-" + suffix

	t.Cleanup(func() {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCleanup()

		queries := []struct {
			query string
			args  []any
		}{
			{query: `DELETE FROM "MealParticipant" WHERE "mealId" = $1`, args: []any{mealID}},
			{query: `DELETE FROM "VolunteeredMeal" WHERE id IN ($1, $2)`, args: []any{firstVolunteeredMealID, secondVolunteeredMealID}},
			{query: `DELETE FROM "Meal" WHERE id IN ($1, $2)`, args: []any{mealID, deletableMealID}},
			{query: `DELETE FROM "User" WHERE id IN ($1, $2)`, args: []any{organizerID, volunteerID}},
		}
		for _, cleanup := range queries {
			if _, err := database.ExecContext(cleanupCtx, cleanup.query, cleanup.args...); err != nil {
				t.Errorf("cleanup %q: %v", cleanup.query, err)
			}
		}
	})

	for _, user := range []struct {
		id       string
		email    string
		fullname string
	}{
		{id: organizerID, email: "organizer@example.com", fullname: "Meal Organizer"},
		{id: volunteerID, email: "volunteer@example.com", fullname: "Meal Volunteer"},
	} {
		if _, err := database.ExecContext(
			operationCtx,
			`INSERT INTO "User" (id, email, fullname, phone) VALUES ($1, $2, $3, $4)`,
			user.id,
			user.email,
			user.fullname,
			"555-0100",
		); err != nil {
			t.Fatalf("insert user %q: %v", user.id, err)
		}
	}

	created, err := repository.Create(operationCtx, domainmeal.CreateParams{
		ID:              mealID,
		Recipient:       "Meal Recipient",
		Reason:          "Community support",
		Email:           "recipient@example.com",
		Phone:           "555-0101",
		City:            "Indianapolis",
		NumOfAdults:     2,
		NumberOfKids:    1,
		Street:          "100 Test Street",
		PreferredTime:   "Evening",
		Favourites:      "Rice",
		LeastFavourites: "Spicy food",
		Allergies:       "None",
		Instructions:    "Call on arrival",
		UserID:          organizerID,
		Meals: []domainmeal.CreateVolunteeredMealParams{
			{ID: firstVolunteeredMealID, Date: "2026-10-20"},
			{ID: secondVolunteeredMealID, Date: "2026-10-21"},
		},
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if created.ID != mealID || created.User == nil || created.User.ID != organizerID || len(created.Meals) != 2 {
		t.Fatalf("Create() = %#v, want hydrated meal and two volunteered meals", created)
	}

	assigned, err := repository.AssignVolunteer(operationCtx, domainmeal.AssignVolunteerParams{
		MealID:            mealID,
		VolunteeredMealID: firstVolunteeredMealID,
		UserID:            volunteerID,
		Description:       "Delivered",
		Notes:             "Left at the front door",
	})
	if err != nil {
		t.Fatalf("AssignVolunteer() error = %v", err)
	}
	if len(assigned.Participants) != 1 ||
		assigned.Participants[0].User == nil ||
		assigned.Participants[0].User.ID != volunteerID ||
		assigned.Meals[0].Description != "Delivered" {
		t.Fatalf("AssignVolunteer() = %#v, want participant and volunteered meal update", assigned)
	}

	updatedVolunteer, err := repository.UpdateVolunteeredMeal(operationCtx, domainmeal.UpdateVolunteeredMealParams{
		ID:          secondVolunteeredMealID,
		Description: "Scheduled",
		Notes:       "Bring cooler",
	})
	if err != nil {
		t.Fatalf("UpdateVolunteeredMeal() error = %v", err)
	}
	if updatedVolunteer.ID != secondVolunteeredMealID || updatedVolunteer.Description != "Scheduled" {
		t.Fatalf("UpdateVolunteeredMeal() = %#v, want updated volunteered meal", updatedVolunteer)
	}

	updated, err := repository.Update(operationCtx, domainmeal.UpdateParams{
		ID:        mealID,
		Recipient: "Updated Recipient",
		Reason:    "Updated reason",
		Email:     "updated@example.com",
		Phone:     "555-0102",
		City:      "Carmel",
	})
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	if updated.Recipient != "Updated Recipient" || updated.Reason != "Updated reason" {
		t.Fatalf("Update() = %#v, want updated meal fields", updated)
	}

	listed, err := repository.List(operationCtx)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if !containsMeal(listed, mealID) {
		t.Fatalf("List() did not return created meal: %#v", listed)
	}

	if _, err := repository.Create(operationCtx, domainmeal.CreateParams{
		ID:        deletableMealID,
		Recipient: "Delete Me",
		Reason:    "Cleanup",
		Email:     "delete@example.com",
		Phone:     "555-0103",
		City:      "Indianapolis",
	}); err != nil {
		t.Fatalf("Create() deletable meal error = %v", err)
	}
	deleted, err := repository.Delete(operationCtx, deletableMealID)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	if deleted.ID != deletableMealID {
		t.Fatalf("Delete() = %#v, want deleted meal", deleted)
	}

	_, err = repository.Delete(operationCtx, deletableMealID)
	if !errors.Is(err, domainmeal.ErrNotFound) {
		t.Fatalf("Delete() missing error = %v, want %v", err, domainmeal.ErrNotFound)
	}
}

func containsMeal(meals []domainmeal.Meal, id string) bool {
	for _, meal := range meals {
		if meal.ID == id {
			return true
		}
	}
	return false
}
