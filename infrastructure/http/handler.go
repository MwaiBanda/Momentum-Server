package http

import (
	eventHandler "Momentum/infrastructure/http/event"
	mealHandler "Momentum/infrastructure/http/meal"
	messageHandler "Momentum/infrastructure/http/message"
	notificationHandler "Momentum/infrastructure/http/notification"
	paymentHandler "Momentum/infrastructure/http/payment"
	sermonHandler "Momentum/infrastructure/http/sermon"
	transactionHandler "Momentum/infrastructure/http/transaction"
	userHandler "Momentum/infrastructure/http/user"
	volunteerServiceHandler "Momentum/infrastructure/http/volunteer_service"
)

type Handler struct {
	User             *userHandler.Handler
	VolunteerService *volunteerServiceHandler.Handler
	Meal             *mealHandler.Handler
	Message          *messageHandler.Handler
	Notification     *notificationHandler.Handler
	Event            *eventHandler.Handler
	Sermon           *sermonHandler.Handler
	Payment          *paymentHandler.Handler
	Transaction      *transactionHandler.Handler
}

func GetInstance(
	user *userHandler.Handler,
	volunteerService *volunteerServiceHandler.Handler,
	meal *mealHandler.Handler,
	message *messageHandler.Handler,
	notification *notificationHandler.Handler,
	event *eventHandler.Handler,
	sermon *sermonHandler.Handler,
	payment *paymentHandler.Handler,
	transaction *transactionHandler.Handler,
) *Handler {
	return &Handler{
		User:             user,
		VolunteerService: volunteerService,
		Meal:             meal,
		Message:          message,
		Notification:     notification,
		Event:            event,
		Sermon:           sermon,
		Payment:          payment,
		Transaction:      transaction,
	}
}
