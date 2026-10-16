package http

import (
	messageHandler "Momentum/infrastructure/http/message"
	notificationHandler "Momentum/infrastructure/http/notification"
	userHandler "Momentum/infrastructure/http/user"
	volunteerServiceHandler "Momentum/infrastructure/http/volunteer_service"
)

type Handler struct {
	User             *userHandler.Handler
	VolunteerService *volunteerServiceHandler.Handler
	Message          *messageHandler.Handler
	Notification     *notificationHandler.Handler
}

func GetInstance(
	user *userHandler.Handler,
	volunteerService *volunteerServiceHandler.Handler,
	message *messageHandler.Handler,
	notification *notificationHandler.Handler,
) *Handler {
	return &Handler{
		User:             user,
		VolunteerService: volunteerService,
		Message:          message,
		Notification:     notification,
	}
}
