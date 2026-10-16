package http

import (
	eventHandler "Momentum/infrastructure/http/event"
	messageHandler "Momentum/infrastructure/http/message"
	notificationHandler "Momentum/infrastructure/http/notification"
	sermonHandler "Momentum/infrastructure/http/sermon"
	userHandler "Momentum/infrastructure/http/user"
	volunteerServiceHandler "Momentum/infrastructure/http/volunteer_service"
)

type Handler struct {
	User             *userHandler.Handler
	VolunteerService *volunteerServiceHandler.Handler
	Message          *messageHandler.Handler
	Notification     *notificationHandler.Handler
	Event            *eventHandler.Handler
	Sermon           *sermonHandler.Handler
}

func GetInstance(
	user *userHandler.Handler,
	volunteerService *volunteerServiceHandler.Handler,
	message *messageHandler.Handler,
	notification *notificationHandler.Handler,
	event *eventHandler.Handler,
	sermon *sermonHandler.Handler,
) *Handler {
	return &Handler{
		User:             user,
		VolunteerService: volunteerService,
		Message:          message,
		Notification:     notification,
		Event:            event,
		Sermon:           sermon,
	}
}
