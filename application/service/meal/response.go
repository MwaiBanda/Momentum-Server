package meal

import (
	"time"

	domainmeal "Momentum/domain/meal"
	domainuser "Momentum/domain/user"
)

type Input struct {
	ID              string                   `json:"id"`
	Recipient       string                   `json:"recipient" binding:"required"`
	Reason          string                   `json:"reason" binding:"required"`
	Email           string                   `json:"email" binding:"required"`
	Phone           string                   `json:"phone" binding:"required"`
	City            string                   `json:"city" binding:"required"`
	NumOfAdults     int                      `json:"num_of_adults" binding:"required"`
	NumberOfKids    int                      `json:"number_of_kids" binding:"required"`
	Street          string                   `json:"street" binding:"required"`
	PreferredTime   string                   `json:"preferred_time" binding:"required"`
	Favourites      string                   `json:"favourites" binding:"required"`
	LeastFavourites string                   `json:"least_favourites" binding:"required"`
	Allergies       string                   `json:"allergies" binding:"required"`
	Instructions    string                   `json:"instructions" binding:"required"`
	Meals           []MealVolunteeredRequest `json:"meals" binding:"required"`
	UserID          string                   `json:"user_id" binding:"required"`
}

type CreateResponse struct {
	ID              string                   `json:"id" binding:"required"`
	Recipient       string                   `json:"recipient" binding:"required"`
	Reason          string                   `json:"reason" binding:"required"`
	Email           string                   `json:"email" binding:"required"`
	Phone           string                   `json:"phone" binding:"required"`
	City            string                   `json:"city" binding:"required"`
	NumOfAdults     int                      `json:"num_of_adults" binding:"required"`
	NumberOfKids    int                      `json:"number_of_kids" binding:"required"`
	Street          string                   `json:"street" binding:"required"`
	PreferredTime   string                   `json:"preferred_time" binding:"required"`
	Favourites      string                   `json:"favourites" binding:"required"`
	LeastFavourites string                   `json:"least_favourites" binding:"required"`
	Allergies       string                   `json:"allergies" binding:"required"`
	Instructions    string                   `json:"instructions" binding:"required"`
	Meals           []MealVolunteeredRequest `json:"meals" binding:"required"`
	UserID          string                   `json:"user_id" binding:"required"`
}

type UpdateResponse struct {
	ID              string                   `json:"id" binding:"required"`
	Recipient       string                   `json:"recipient" binding:"required"`
	Reason          string                   `json:"reason" binding:"required"`
	Email           string                   `json:"email" binding:"required"`
	Phone           string                   `json:"phone" binding:"required"`
	City            string                   `json:"city" binding:"required"`
	NumOfAdults     int                      `json:"num_of_adults" binding:"required"`
	NumberOfKids    int                      `json:"number_of_kids" binding:"required"`
	Street          string                   `json:"street" binding:"required"`
	PreferredTime   string                   `json:"preferred_time" binding:"required"`
	Favourites      string                   `json:"favourites" binding:"required"`
	LeastFavourites string                   `json:"least_favourites" binding:"required"`
	Allergies       string                   `json:"allergies" binding:"required"`
	Instructions    string                   `json:"instructions" binding:"required"`
	Meals           []MealVolunteeredRequest `json:"meals" binding:"required"`
	UserID          string                   `json:"user_id" binding:"required"`
}

type MealVolunteeredRequest struct {
	Description string `json:"description"`
	Notes       string `json:"notes"`
	MealID      string `json:"meal_id"`
	UserID      string `json:"user_id"`
	Date        string `json:"date" binding:"required"`
}

type Response struct {
	ID              string                    `json:"id" binding:"required"`
	Recipient       string                    `json:"recipient" binding:"required"`
	Reason          string                    `json:"reason" binding:"required"`
	Email           string                    `json:"email" binding:"required"`
	Phone           string                    `json:"phone" binding:"required"`
	City            string                    `json:"city" binding:"required"`
	NumOfAdults     int                       `json:"num_of_adults" binding:"required"`
	NumberOfKids    int                       `json:"number_of_kids" binding:"required"`
	Street          string                    `json:"street" binding:"required"`
	PreferredTime   string                    `json:"preferred_time" binding:"required"`
	Favourites      string                    `json:"favourites" binding:"required"`
	LeastFavourites string                    `json:"least_favourites" binding:"required"`
	Allergies       string                    `json:"allergies" binding:"required"`
	Instructions    string                    `json:"instructions" binding:"required"`
	Participants    []ParticipantResponse     `json:"participants" binding:"required"`
	Meals           []VolunteeredMealResponse `json:"meals" binding:"required"`
	User            UserResponse              `json:"user" binding:"required"`
}

type ParticipantResponse struct {
	ID   string       `json:"id" binding:"required"`
	User UserResponse `json:"user" binding:"required"`
}

type VolunteeredMealResponse struct {
	ID          string                `json:"id" binding:"required"`
	CreatedOn   string                `json:"created_on" binding:"required"`
	Description string                `json:"description" binding:"required"`
	Notes       string                `json:"notes" binding:"required"`
	Date        string                `json:"date" binding:"required"`
	User        UserResponse          `json:"user" binding:"required"`
	Updates     []UpdateEntryResponse `json:"updates" binding:"required"`
}

type UpdateEntryResponse struct {
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Fullname  string `json:"fullname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	CreatedOn string `json:"createdOn" binding:"required"`
}

type UserInput struct {
	ID string `json:"id" binding:"required"`
}

type VolunteerAssignmentInput struct {
	MealID          string               `json:"meal_id" binding:"required"`
	VolunteeredMeal VolunteeredMealInput `json:"volunteered_meal" binding:"required"`
}

type VolunteerAssignmentResponse struct {
	MealID          string                  `json:"meal_id" binding:"required"`
	VolunteeredMeal VolunteeredMealResponse `json:"volunteered_meal" binding:"required"`
}

type VolunteeredMealInput struct {
	ID          string                `json:"id" binding:"required"`
	CreatedOn   string                `json:"created_on"`
	Description string                `json:"description"`
	Notes       string                `json:"notes"`
	Date        string                `json:"date"`
	User        UserInput             `json:"user" binding:"required"`
	Updates     []UpdateEntryResponse `json:"updates"`
}

type UpdateVolunteeredMealInput struct {
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Notes       string `json:"notes" binding:"required"`
}

type UpdateVolunteeredMealResponse struct {
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Notes       string `json:"notes" binding:"required"`
}

func toResponse(entity domainmeal.Meal) Response {
	participants := make([]ParticipantResponse, 0, len(entity.Participants))
	for _, participant := range entity.Participants {
		participants = append(participants, ParticipantResponse{
			ID:   participant.ID,
			User: toUserResponse(participant.User),
		})
	}

	volunteeredMeals := make([]VolunteeredMealResponse, 0, len(entity.Meals))
	for _, volunteeredMeal := range entity.Meals {
		updates := make([]UpdateEntryResponse, 0, len(volunteeredMeal.Updates))
		for _, update := range volunteeredMeal.Updates {
			updates = append(updates, UpdateEntryResponse{
				ID:      update.ID,
				Message: update.Message,
			})
		}
		volunteeredMeals = append(volunteeredMeals, VolunteeredMealResponse{
			ID:          volunteeredMeal.ID,
			CreatedOn:   formatTime(volunteeredMeal.CreatedOn),
			Description: volunteeredMeal.Description,
			Notes:       volunteeredMeal.Notes,
			Date:        volunteeredMeal.Date,
			User:        toUserResponse(volunteeredMeal.User),
			Updates:     updates,
		})
	}

	return Response{
		ID:              entity.ID,
		Recipient:       entity.Recipient,
		Reason:          entity.Reason,
		Email:           entity.Email,
		Phone:           entity.Phone,
		City:            entity.City,
		NumOfAdults:     entity.NumOfAdults,
		NumberOfKids:    entity.NumberOfKids,
		Street:          entity.Street,
		PreferredTime:   entity.PreferredTime,
		Favourites:      entity.Favourites,
		LeastFavourites: entity.LeastFavourites,
		Allergies:       entity.Allergies,
		Instructions:    entity.Instructions,
		Participants:    participants,
		Meals:           volunteeredMeals,
		User:            toUserResponse(entity.User),
	}
}

func toUserResponse(user *domainuser.User) UserResponse {
	if user == nil {
		return UserResponse{}
	}

	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Fullname:  user.Fullname,
		Phone:     user.Phone,
		CreatedOn: formatTime(user.CreatedOn),
	}
}

func formatTime(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.Format(time.RFC3339Nano)
}
