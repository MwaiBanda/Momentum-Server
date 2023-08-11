package model

type Meal struct {
	Id              string `json:"id"`
	Recipient       string `json:"recipient"`
	Reason          string `json:"reason"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	City            string `json:"city"`
	NumOfAdults     int    `json:"num_of_adults"`
	NumberOfKids    int    `json:"number_of_kids"`
	Street          string `json:"street"`
	PreferredTime   string `json:"preferred_time"`
	Favourites      string `json:"favourites"`
	LeastFavourites string `json:"least_favourites"`
	Allergies       string `json:"allergies"`
	Instructions    string `json:"instructions"`
	UserId          string `json:"user_id"`
	User            User   `json:"user"`
}
