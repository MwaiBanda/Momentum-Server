package model

type TransactionRequest struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
}

type TransactionResponse struct {
	ID          string       `json:"id"`
	Amount      int          `json:"amount"`
	Date        string       `json:"date"`
	CreatedOn   string       `json:"createdOn"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
}

type TransactionHookRequest struct {
	Amount      int          `json:"amount"`
	Description string       `json:"description"`
	Fullname	string       `json:"fullname"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
}
