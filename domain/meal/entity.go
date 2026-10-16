package meal

import (
	"time"

	domainuser "Momentum/domain/user"
)

type Meal struct {
	ID              string
	Reason          string
	Email           string
	Phone           string
	City            string
	NumOfAdults     int
	NumberOfKids    int
	Street          string
	PreferredTime   string
	Favourites      string
	LeastFavourites string
	Allergies       string
	Instructions    string
	UserID          string
	Recipient       string
	CreatedOn       time.Time
	User            *domainuser.User
	Participants    []Participant
	Meals           []VolunteeredMeal
}

type Participant struct {
	ID         string
	UserID     string
	MealID     string
	AssignedAt time.Time
	User       *domainuser.User
}

type VolunteeredMeal struct {
	ID          string
	Description string
	Notes       string
	UserID      string
	MealID      string
	CreatedOn   time.Time
	Date        string
	User        *domainuser.User
	Updates     []Update
}

type Update struct {
	ID        string
	Message   string
	MealID    string
	CreatedOn time.Time
}

type CreateParams struct {
	ID              string
	Recipient       string
	Reason          string
	Email           string
	Phone           string
	City            string
	NumOfAdults     int
	NumberOfKids    int
	Street          string
	PreferredTime   string
	Favourites      string
	LeastFavourites string
	Allergies       string
	Instructions    string
	UserID          string
	Meals           []CreateVolunteeredMealParams
}

type CreateVolunteeredMealParams struct {
	ID   string
	Date string
}

type UpdateParams struct {
	ID              string
	Recipient       string
	Reason          string
	Email           string
	Phone           string
	City            string
	NumOfAdults     int
	NumberOfKids    int
	Street          string
	PreferredTime   string
	Favourites      string
	LeastFavourites string
	Allergies       string
	Instructions    string
}

type AssignVolunteerParams struct {
	MealID            string
	VolunteeredMealID string
	UserID            string
	Description       string
	Notes             string
}

type UpdateVolunteeredMealParams struct {
	ID          string
	Description string
	Notes       string
}
