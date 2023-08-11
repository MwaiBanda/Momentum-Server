package model

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	CreatedOn string `json:"createdOn"`
	Meals     []Meal `json:"meals"`
}
