package user

import (
	userService "Momentum/application/service/user"
)

type Handler struct {
	users userService.UseCase
}

func GetInstance(users userService.UseCase) *Handler {
	return &Handler{
		users: users,
	}
}
