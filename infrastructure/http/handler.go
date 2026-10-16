package http

import userHandler "Momentum/infrastructure/http/user"

type Handler struct {
	User *userHandler.Handler
}

func GetInstance(user *userHandler.Handler) *Handler {
	return &Handler{
		User: user,
	}
}
