package model

type Post struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Published   bool   `json:"published"`
	Description string `json:"desc"`
}
