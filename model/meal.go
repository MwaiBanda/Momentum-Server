package model

type MealRequest struct {
	Id			  	string                   `json:"id"`
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
	UserId          string                   `json:"user_id"`
}

type MealResponse struct {
	Id              string            `json:"id"`
	Recipient       string            `json:"recipient"`
	Reason          string            `json:"reason"`
	Email           string            `json:"email"`
	Phone           string            `json:"phone"`
	City            string            `json:"city"`
	NumOfAdults     int               `json:"num_of_adults"`
	NumberOfKids    int               `json:"number_of_kids"`
	Street          string            `json:"street"`
	PreferredTime   string            `json:"preferred_time"`
	Favourites      string            `json:"favourites"`
	LeastFavourites string            `json:"least_favourites"`
	Allergies       string            `json:"allergies"`
	Instructions    string            `json:"instructions"`
	Participants    []Participant     `json:"participants"`
	Meals           []VolunteeredMeal `json:"meals"`
	User            UserResponse      `json:"user"`
}

type MealVolunteeredRequest struct {
	Description string `json:"description"`
	Notes       string `json:"notes"`
	MealId      string `json:"meal_id"`
	UserId      string `json:"user_id"`
	Date        string `json:"date"`
}

type VolunteeredMealRequest struct {
	MealId          string          `json:"meal_id"`
	VolunteeredMeal VolunteeredMeal `json:"volunteered_meal"`
}

type VolunteeredMeal struct {
	Id          string        `json:"id"`
	CreatedOn   string        `json:"created_on"`
	Description string        `json:"description"`
	Notes       string        `json:"notes"`
	Date        string        `json:"date"`
	User        UserResponse  `json:"user"`
	Updates     []MealUpdates `json:"updates"`
}

type MealUpdates struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}
