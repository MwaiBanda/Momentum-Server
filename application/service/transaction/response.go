package transaction

import (
	"time"

	domaintransaction "Momentum/domain/transaction"
)

type Input struct {
	ID          string  `json:"id"`
	Amount      int     `json:"amount" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Fullname    *string `json:"fullname"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	UserID      string  `json:"userId"`
}

type Response struct {
	ID          string       `json:"id" binding:"required"`
	Amount      int          `json:"amount" binding:"required"`
	Date        string       `json:"date" binding:"required"`
	CreatedOn   string       `json:"createdOn" binding:"required"`
	Description string       `json:"description" binding:"required"`
	User        UserResponse `json:"user" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Fullname  string `json:"fullname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	CreatedOn string `json:"createdOn" binding:"required"`
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
