package user

import (
	"time"

	domainuser "Momentum/domain/user"
)

type Response struct {
	ID        string `json:"id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Fullname  string `json:"fullname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	CreatedOn string `json:"createdOn" binding:"required"`
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
