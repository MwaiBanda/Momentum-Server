package transaction

import (
	"time"

	domaintransaction "Momentum/domain/transaction"
)

type Input struct {
	ID          string  `json:"id"`
	Amount      int     `json:"amount"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Fullname    *string `json:"fullname"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	UserID      string  `json:"userId"`
}

type Response struct {
	ID          string       `json:"id"`
	Amount      int          `json:"amount"`
	Date        string       `json:"date"`
	CreatedOn   string       `json:"createdOn"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	CreatedOn string `json:"createdOn"`
}

type MailInput struct {
	Amount      int
	Description string
	Fullname    string
	Email       string
	Phone       string
}

func toResponse(transaction domaintransaction.Transaction) Response {
	return Response{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		CreatedOn:   transaction.CreatedOn.Format(time.RFC3339Nano),
		Description: transaction.Description,
		User: UserResponse{
			ID:        transaction.User.ID,
			Email:     transaction.User.Email,
			Fullname:  transaction.User.Fullname,
			Phone:     transaction.User.Phone,
			CreatedOn: transaction.User.CreatedOn.Format(time.RFC3339Nano),
		},
	}
}
