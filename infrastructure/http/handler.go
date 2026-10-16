package http

import (
	messageHandler "Momentum/infrastructure/http/message"
	userHandler "Momentum/infrastructure/http/user"
	volunteerServiceHandler "Momentum/infrastructure/http/volunteer_service"
)

type Handler struct {
	User             *userHandler.Handler
	VolunteerService *volunteerServiceHandler.Handler
	Message          *messageHandler.Handler
}

func GetInstance(
	user *userHandler.Handler,
	volunteerService *volunteerServiceHandler.Handler,
	message *messageHandler.Handler,
) *Handler {
	return &Handler{
		User:             user,
		VolunteerService: volunteerService,
		Message:          message,
	}
}
