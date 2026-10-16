package meal

import (
	"time"

	domainmeal "Momentum/domain/meal"
	domainuser "Momentum/domain/user"
)

type Input struct {
	ID              string                   `json:"id"`
	Recipient       string                   `json:"recipient"`
	Reason          string                   `json:"reason"`
	Email           string                   `json:"email"`
	Phone           string                   `json:"phone"`
	City            string                   `json:"city"`
	NumOfAdults     int                      `json:"num_of_adults"`
	NumberOfKids    int                      `json:"number_of_kids"`
	Street          string                   `json:"street"`
	PreferredTime   string                   `json:"preferred_time"`
	Favourites      string                   `json:"favourites"`
	LeastFavourites string                   `json:"least_favourites"`
	Allergies       string                   `json:"allergies"`
	Instructions    string                   `json:"instructions"`
	Meals           []MealVolunteeredRequest `json:"meals"`
	UserID          string                   `json:"user_id"`
}

type MealVolunteeredRequest struct {
	Description string `json:"description"`
	Notes       string `json:"notes"`
	MealID      string `json:"meal_id"`
	UserID      string `json:"user_id"`
	Date        string `json:"date"`
}

type Response struct {
	ID              string                    `json:"id"`
	Recipient       string                    `json:"recipient"`
	Reason          string                    `json:"reason"`
	Email           string                    `json:"email"`
	Phone           string                    `json:"phone"`
	City            string                    `json:"city"`
	NumOfAdults     int                       `json:"num_of_adults"`
	NumberOfKids    int                       `json:"number_of_kids"`
	Street          string                    `json:"street"`
	PreferredTime   string                    `json:"preferred_time"`
	Favourites      string                    `json:"favourites"`
	LeastFavourites string                    `json:"least_favourites"`
	Allergies       string                    `json:"allergies"`
	Instructions    string                    `json:"instructions"`
	Participants    []ParticipantResponse     `json:"participants"`
	Meals           []VolunteeredMealResponse `json:"meals"`
	User            UserResponse              `json:"user"`
}

type ParticipantResponse struct {
	ID   string       `json:"id"`
	User UserResponse `json:"user"`
}

type VolunteeredMealResponse struct {
	ID          string           `json:"id"`
	CreatedOn   string           `json:"created_on"`
	Description string           `json:"description"`
	Notes       string           `json:"notes"`
	Date        string           `json:"date"`
	User        UserResponse     `json:"user"`
	Updates     []UpdateResponse `json:"updates"`
}

type UpdateResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	CreatedOn string `json:"createdOn"`
}

type VolunteerAssignmentInput struct {
	MealID          string               `json:"meal_id"`
	VolunteeredMeal VolunteeredMealInput `json:"volunteered_meal"`
}

type VolunteeredMealInput struct {
	ID          string           `json:"id"`
	CreatedOn   string           `json:"created_on"`
	Description string           `json:"description"`
	Notes       string           `json:"notes"`
	Date        string           `json:"date"`
	User        UserResponse     `json:"user"`
	Updates     []UpdateResponse `json:"updates"`
}

type UpdateVolunteeredMealInput struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
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
		updates := make([]UpdateResponse, 0, len(volunteeredMeal.Updates))
		for _, update := range volunteeredMeal.Updates {
			updates = append(updates, UpdateResponse{
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
