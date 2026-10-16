package volunteer_service

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	domainservice "Momentum/domain/volunteer_service"
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

func TestRepositoryIntegrationVolunteerService(t *testing.T) {
	database := openIntegrationDatabase(t)
	repository := GetInstance(database)
	operationCtx, cancelOperation := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelOperation()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	catalogID := "repository-test-catalog-" + suffix
	volunteerID := "repository-test-volunteer-" + suffix
	firstDayID := "repository-test-day-1-" + suffix
	secondDayID := "repository-test-day-2-" + suffix
	userID := "repository-test-user-" + suffix

	t.Cleanup(func() {
		cleanupCtx, cancelCleanup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCleanup()

		queries := []struct {
			query string
			arg   string
		}{
			{query: `DELETE FROM "VolunteeredServiceDay" WHERE id IN ($1, $2)`, arg: firstDayID},
			{query: `DELETE FROM "VolunteerService" WHERE id = $1`, arg: volunteerID},
			{query: `DELETE FROM "Service" WHERE id = $1`, arg: catalogID},
			{query: `DELETE FROM "User" WHERE id = $1`, arg: userID},
		}
		for index, cleanup := range queries {
			var err error
			if index == 0 {
				_, err = database.ExecContext(cleanupCtx, cleanup.query, cleanup.arg, secondDayID)
			} else {
				_, err = database.ExecContext(cleanupCtx, cleanup.query, cleanup.arg)
			}
			if err != nil {
				t.Errorf("cleanup query %d: %v", index, err)
			}
		}
	})

	if _, err := database.ExecContext(
		operationCtx,
		`INSERT INTO "Service" (id, name, type, "Image") VALUES ($1, $2, $3, $4)`,
		catalogID,
		"Repository Test Service",
		"repository-test-"+suffix,
		"test.png",
	); err != nil {
		t.Fatalf("insert catalog service: %v", err)
	}
	if _, err := database.ExecContext(
		operationCtx,
		`INSERT INTO "User" (id, email, fullname, phone) VALUES ($1, $2, $3, $4)`,
		userID,
		"repository-test@example.com",
		"Repository Test User",
		"555-0100",
	); err != nil {
		t.Fatalf("insert assignment user: %v", err)
	}

	if err := repository.Create(operationCtx, domainservice.CreateParams{
		ID:          volunteerID,
		Description: "Repository test description",
		Email:       "organizer@example.com",
		Phone:       "555-0101",
		Organizer:   "Repository Test Organizer",
		ServiceID:   catalogID,
		Days: []domainservice.CreateDayParams{
			{ID: firstDayID, Date: "2026-04-01"},
			{ID: secondDayID, Date: "2026-04-08"},
		},
	}); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	serviceType := "repository-test-" + suffix
	detail, err := repository.GetByType(operationCtx, serviceType)
	if err != nil {
		t.Fatalf("GetByType() error = %v", err)
	}
	if detail.Service.ID != catalogID || len(detail.Services) != 1 {
		t.Fatalf("GetByType() detail = %#v, want catalog and one volunteer service", detail)
	}
	if len(detail.Services[0].Days) != 2 {
		t.Fatalf("GetByType() days = %d, want 2", len(detail.Services[0].Days))
	}

	updated, err := repository.UpdateDay(operationCtx, domainservice.UpdateDayParams{
		ID:     firstDayID,
		Notes:  "Assigned by repository test",
		UserID: userID,
	})
	if err != nil {
		t.Fatalf("UpdateDay() error = %v", err)
	}
	if updated.ID != firstDayID || updated.Notes != "Assigned by repository test" || updated.UserID != userID {
		t.Fatalf("UpdateDay() = %#v, want updated day", updated)
	}

	detail, err = repository.GetByType(operationCtx, serviceType)
	if err != nil {
		t.Fatalf("GetByType() after update error = %v", err)
	}
	assignedDay := detail.Services[0].Days[0]
	if assignedDay.User == nil || assignedDay.User.ID != userID {
		t.Fatalf("assigned day user = %#v, want %q", assignedDay.User, userID)
	}
}
