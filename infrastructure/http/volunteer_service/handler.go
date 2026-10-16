package volunteer_service

import (
	applicationservice "Momentum/application/service/volunteer_service"
)

type Handler struct {
	services applicationservice.UseCase
}

func GetInstance(services applicationservice.UseCase) *Handler {
	return &Handler{services: services}
}
