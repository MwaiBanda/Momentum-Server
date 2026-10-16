package volunteer_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	domainuser "Momentum/domain/user"
	domainservice "Momentum/domain/volunteer_service"
	"Momentum/infrastructure/persistence"
)

type detailRow struct {
	VolunteerID        string         `db:"volunteer_id"`
	Description        string         `db:"description"`
	Email              string         `db:"email"`
	Phone              string         `db:"phone"`
	Organizer          string         `db:"organizer"`
	VolunteerServiceID sql.NullString `db:"volunteer_service_id"`
	VolunteerUserID    sql.NullString `db:"volunteer_user_id"`
	VolunteerCreatedOn sql.NullTime   `db:"volunteer_created_on"`
	DayID              sql.NullString `db:"day_id"`
	DayNotes           sql.NullString `db:"day_notes"`
	DayDate            sql.NullString `db:"day_date"`
	DayUserID          sql.NullString `db:"day_user_id"`
	DayCreatedOn       sql.NullTime   `db:"day_created_on"`
	UserID             sql.NullString `db:"user_id"`
	UserEmail          sql.NullString `db:"user_email"`
	UserFullname       sql.NullString `db:"user_fullname"`
	UserPhone          sql.NullString `db:"user_phone"`
	UserCreatedOn      sql.NullTime   `db:"user_created_on"`
}

func (repository *Repository) GetByType(ctx context.Context, serviceType string) (domainservice.ServiceDetail, error) {
	var catalogService domainservice.Service
	serviceQuery := persistence.LoadFile("repository/volunteer_service/sql/get_service_by_type.sql")
	if err := repository.db.GetContext(ctx, &catalogService, serviceQuery, serviceType); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainservice.ServiceDetail{}, fmt.Errorf("get service type %q: %w", serviceType, domainservice.ErrServiceNotFound)
		}
		return domainservice.ServiceDetail{}, fmt.Errorf("get service type %q: %w", serviceType, err)
	}

	var rows []detailRow
	detailQuery := persistence.LoadFile("repository/volunteer_service/sql/get_by_service_id.sql")
	if err := repository.db.SelectContext(ctx, &rows, detailQuery, catalogService.ID); err != nil {
		return domainservice.ServiceDetail{}, fmt.Errorf("get volunteer services for service type %q: %w", serviceType, err)
	}

	servicesByID := make(map[string]int)
	detail := domainservice.ServiceDetail{
		Service:  catalogService,
		Services: make([]domainservice.VolunteerService, 0),
	}

	for _, row := range rows {
		index, exists := servicesByID[row.VolunteerID]
		if !exists {
			index = len(detail.Services)
			servicesByID[row.VolunteerID] = index
			detail.Services = append(detail.Services, domainservice.VolunteerService{
				ID:          row.VolunteerID,
				Description: row.Description,
				Email:       row.Email,
				Phone:       row.Phone,
				Organizer:   row.Organizer,
				ServiceID:   nullStringValue(row.VolunteerServiceID),
				UserID:      nullStringValue(row.VolunteerUserID),
				CreatedOn:   nullTimeValue(row.VolunteerCreatedOn),
				Days:        make([]domainservice.VolunteerServiceDay, 0),
			})
		}

		if !row.DayID.Valid {
			continue
		}

		var user *domainuser.User
		if row.UserID.Valid {
			user = &domainuser.User{
				ID:        row.UserID.String,
				Email:     row.UserEmail.String,
				Fullname:  row.UserFullname.String,
				Phone:     row.UserPhone.String,
				CreatedOn: nullTimeValue(row.UserCreatedOn),
			}
		}

		detail.Services[index].Days = append(detail.Services[index].Days, domainservice.VolunteerServiceDay{
			ID:        row.DayID.String,
			Notes:     row.DayNotes.String,
			Date:      row.DayDate.String,
			UserID:    nullStringValue(row.DayUserID),
			CreatedOn: nullTimeValue(row.DayCreatedOn),
			User:      user,
		})
	}

	return detail, nil
}

func nullStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	return value.String
}

func nullTimeValue(value sql.NullTime) (result time.Time) {
	if !value.Valid {
		return result
	}
	return value.Time
}
