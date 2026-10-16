package user

import "time"

type User struct {
	ID        string
	Email     string
	Fullname  string
	Phone     string
	CreatedOn time.Time `db:"createdOn"`
}
