package transaction

import (
	"database/sql"
	"time"

	domaintransaction "Momentum/domain/transaction"
	domainuser "Momentum/domain/user"
)

type transactionRow struct {
	ID            string       `db:"id"`
	Amount        int          `db:"amount"`
	Date          string       `db:"date"`
	CreatedOn     sql.NullTime `db:"created_on"`
	Description   string       `db:"description"`
	UserID        string       `db:"user_id"`
	UserEmail     string       `db:"user_email"`
	UserFullname  string       `db:"user_fullname"`
	UserPhone     string       `db:"user_phone"`
	UserCreatedOn sql.NullTime `db:"user_created_on"`
}

func (row transactionRow) entity() domaintransaction.Transaction {
	return domaintransaction.Transaction{
		ID:          row.ID,
		Amount:      row.Amount,
		Date:        row.Date,
		CreatedOn:   nullTimeValue(row.CreatedOn),
		Description: row.Description,
		User: domainuser.User{
			ID:        row.UserID,
			Email:     row.UserEmail,
			Fullname:  row.UserFullname,
			Phone:     row.UserPhone,
			CreatedOn: nullTimeValue(row.UserCreatedOn),
		},
	}
}

func nullTimeValue(value sql.NullTime) time.Time {
	if !value.Valid {
		return time.Time{}
	}
	return value.Time
}
