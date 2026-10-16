package user

import (
	"time"

	domainuser "Momentum/domain/user"
)

type Response struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	CreatedOn string `json:"createdOn"`
}

func toResponse(userEntity domainuser.User) Response {
	return Response{
		ID:        userEntity.ID,
		Email:     userEntity.Email,
		Fullname:  userEntity.Fullname,
		Phone:     userEntity.Phone,
		CreatedOn: userEntity.CreatedOn.Format(time.RFC3339Nano),
	}
}
