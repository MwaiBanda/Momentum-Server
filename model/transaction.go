package model

type Transaction struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	CreatedOn   string `json:"createdOn"`
	UserId      string `json:"userId"`
	Description string `json:"description"`
}
