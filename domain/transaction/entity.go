package transaction

import (
	"time"

	domainuser "Momentum/domain/user"
)

type Transaction struct {
	ID          string
	Amount      int
	Date        string
	CreatedOn   time.Time
	Description string
	User        domainuser.User
}

type CreateParams struct {
	ID          string
	Amount      int
	Date        string
	Description string
	UserID      string
	Fullname    string
	Email       string
	Phone       string
}
