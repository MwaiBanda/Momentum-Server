package model

type UserResponse struct {
	Id        string         `json:"id"`
	Email     string         `json:"email"`
	Fullname  string         `json:"fullname"`
	Phone     string         `json:"phone"`
	CreatedOn string         `json:"createdOn"`
	Meals     []MealResponse `json:"meals"`
}

type UserRequest struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
}
