package model

type Transaction struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	CreatedOn   string `json:"created_on"`
	UserId      string `json:"user_id"`
	Description string `json:"description"`
}
