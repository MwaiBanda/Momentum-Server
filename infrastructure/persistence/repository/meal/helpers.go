package meal

import (
	"database/sql"
	"time"

	domainmeal "Momentum/domain/meal"
	domainuser "Momentum/domain/user"
)

type aggregateRow struct {
	MealID               string         `db:"meal_id"`
	Recipient            string         `db:"recipient"`
	Reason               string         `db:"reason"`
	Email                string         `db:"email"`
	Phone                string         `db:"phone"`
	City                 string         `db:"city"`
	NumOfAdults          int            `db:"num_of_adults"`
	NumberOfKids         int            `db:"number_of_kids"`
	Street               string         `db:"street"`
	PreferredTime        string         `db:"preferred_time"`
	Favourites           string         `db:"favourites"`
	LeastFavourites      string         `db:"least_favourites"`
	Allergies            string         `db:"allergies"`
	Instructions         string         `db:"instructions"`
	MealUserID           sql.NullString `db:"meal_user_id"`
	CreatedOn            time.Time      `db:"created_on"`
	MealUserEmail        sql.NullString `db:"meal_user_email"`
	MealUserFullname     sql.NullString `db:"meal_user_fullname"`
	MealUserPhone        sql.NullString `db:"meal_user_phone"`
	MealUserCreatedOn    sql.NullTime   `db:"meal_user_created_on"`
	ParticipantID        sql.NullString `db:"participant_id"`
	ParticipantUserID    sql.NullString `db:"participant_user_id"`
	ParticipantAssigned  sql.NullTime   `db:"participant_assigned_at"`
	ParticipantEmail     sql.NullString `db:"participant_email"`
	ParticipantFullname  sql.NullString `db:"participant_fullname"`
	ParticipantPhone     sql.NullString `db:"participant_phone"`
	ParticipantCreated   sql.NullTime   `db:"participant_created_on"`
	VolunteerID          sql.NullString `db:"volunteer_id"`
	VolunteerDescription sql.NullString `db:"volunteer_description"`
	VolunteerNotes       sql.NullString `db:"volunteer_notes"`
	VolunteerUserID      sql.NullString `db:"volunteer_user_id"`
	VolunteerMealID      sql.NullString `db:"volunteer_meal_id"`
	VolunteerCreatedOn   sql.NullTime   `db:"volunteer_created_on"`
	VolunteerDate        sql.NullString `db:"volunteer_date"`
	VolunteerEmail       sql.NullString `db:"volunteer_email"`
	VolunteerFullname    sql.NullString `db:"volunteer_fullname"`
	VolunteerPhone       sql.NullString `db:"volunteer_phone"`
	VolunteerUserCreated sql.NullTime   `db:"volunteer_user_created_on"`
	UpdateID             sql.NullString `db:"update_id"`
	UpdateMessage        sql.NullString `db:"update_message"`
	UpdateCreatedOn      sql.NullTime   `db:"update_created_on"`
}

func aggregateRows(rows []aggregateRow) []domainmeal.Meal {
	meals := make([]domainmeal.Meal, 0)
	mealIndexes := make(map[string]int)
	participantIndexes := make(map[string]map[string]int)
	volunteerIndexes := make(map[string]map[string]int)
	updateIndexes := make(map[string]map[string]int)

	for _, row := range rows {
		mealIndex, exists := mealIndexes[row.MealID]
		if !exists {
			mealIndex = len(meals)
			mealIndexes[row.MealID] = mealIndex
			participantIndexes[row.MealID] = make(map[string]int)
			volunteerIndexes[row.MealID] = make(map[string]int)
			meals = append(meals, domainmeal.Meal{
				ID:              row.MealID,
				Recipient:       row.Recipient,
				Reason:          row.Reason,
				Email:           row.Email,
				Phone:           row.Phone,
				City:            row.City,
				NumOfAdults:     row.NumOfAdults,
				NumberOfKids:    row.NumberOfKids,
				Street:          row.Street,
				PreferredTime:   row.PreferredTime,
				Favourites:      row.Favourites,
				LeastFavourites: row.LeastFavourites,
				Allergies:       row.Allergies,
				Instructions:    row.Instructions,
				UserID:          nullStringValue(row.MealUserID),
				CreatedOn:       row.CreatedOn,
				User:            userFromRow(row.MealUserID, row.MealUserEmail, row.MealUserFullname, row.MealUserPhone, row.MealUserCreatedOn),
				Participants:    make([]domainmeal.Participant, 0),
				Meals:           make([]domainmeal.VolunteeredMeal, 0),
			})
		}

		if row.ParticipantID.Valid {
			if _, exists := participantIndexes[row.MealID][row.ParticipantID.String]; !exists {
				participantIndexes[row.MealID][row.ParticipantID.String] = len(meals[mealIndex].Participants)
				meals[mealIndex].Participants = append(meals[mealIndex].Participants, domainmeal.Participant{
					ID:         row.ParticipantID.String,
					UserID:     nullStringValue(row.ParticipantUserID),
					MealID:     row.MealID,
					AssignedAt: nullTimeValue(row.ParticipantAssigned),
					User:       userFromRow(row.ParticipantUserID, row.ParticipantEmail, row.ParticipantFullname, row.ParticipantPhone, row.ParticipantCreated),
				})
			}
		}

		if !row.VolunteerID.Valid {
			continue
		}

		volunteerIndex, exists := volunteerIndexes[row.MealID][row.VolunteerID.String]
		if !exists {
			volunteerIndex = len(meals[mealIndex].Meals)
			volunteerIndexes[row.MealID][row.VolunteerID.String] = volunteerIndex
			updateIndexes[row.VolunteerID.String] = make(map[string]int)
			meals[mealIndex].Meals = append(meals[mealIndex].Meals, domainmeal.VolunteeredMeal{
				ID:          row.VolunteerID.String,
				Description: nullStringValue(row.VolunteerDescription),
				Notes:       nullStringValue(row.VolunteerNotes),
				UserID:      nullStringValue(row.VolunteerUserID),
				MealID:      nullStringValue(row.VolunteerMealID),
				CreatedOn:   nullTimeValue(row.VolunteerCreatedOn),
				Date:        nullStringValue(row.VolunteerDate),
				User:        userFromRow(row.VolunteerUserID, row.VolunteerEmail, row.VolunteerFullname, row.VolunteerPhone, row.VolunteerUserCreated),
				Updates:     make([]domainmeal.Update, 0),
			})
		}

		if row.UpdateID.Valid {
			if _, exists := updateIndexes[row.VolunteerID.String][row.UpdateID.String]; !exists {
				updateIndexes[row.VolunteerID.String][row.UpdateID.String] = len(meals[mealIndex].Meals[volunteerIndex].Updates)
				meals[mealIndex].Meals[volunteerIndex].Updates = append(meals[mealIndex].Meals[volunteerIndex].Updates, domainmeal.Update{
					ID:        row.UpdateID.String,
					Message:   row.UpdateMessage.String,
					MealID:    row.VolunteerID.String,
					CreatedOn: nullTimeValue(row.UpdateCreatedOn),
				})
			}
		}
	}

	return meals
}

func userFromRow(id, email, fullname, phone sql.NullString, createdOn sql.NullTime) *domainuser.User {
	if !id.Valid {
		return nil
	}
	return &domainuser.User{
		ID:        id.String,
		Email:     email.String,
		Fullname:  fullname.String,
		Phone:     phone.String,
		CreatedOn: nullTimeValue(createdOn),
	}
}

func nullStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	return value.String
}

func nullTimeValue(value sql.NullTime) time.Time {
	if !value.Valid {
		return time.Time{}
	}
	return value.Time
}
