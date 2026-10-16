package http

import (
	userHandler "Momentum/infrastructure/http/user"
	volunteerServiceHandler "Momentum/infrastructure/http/volunteer_service"
)

type Handler struct {
	User             *userHandler.Handler
	VolunteerService *volunteerServiceHandler.Handler
}

func GetInstance(
	user *userHandler.Handler,
	volunteerService *volunteerServiceHandler.Handler,
) *Handler {
	return &Handler{
		User:             user,
		VolunteerService: volunteerService,
	}
}
