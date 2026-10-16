package volunteer_service

import (
	"errors"
	"log"

	applicationservice "Momentum/application/service/volunteer_service"

	"github.com/gofiber/fiber/v3"
)

// GetByType godoc
//
//	@Summary		Get Service By Type
//	@Description	get a service by its type
//	@Accept			json
//	@Produce		json
//	@tags			Services
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Param			type			path		string	true	"provide service type"
//	@Success		200				{object}	volunteer_service.ServiceDetailResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/services/{type} [get]
func (handler *Handler) GetByType(ctx fiber.Ctx) error {
	if handler.services == nil {
		log.Printf("[GetServiceByType] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.services.GetByType(ctx.Context(), ctx.Params("type"))
	if err != nil {
		switch {
		case errors.Is(err, applicationservice.ErrInvalidInput):
			return fiber.ErrBadRequest
		case errors.Is(err, applicationservice.ErrServiceNotFound):
			return fiber.ErrNotFound
		default:
			log.Printf("[GetServiceByType] %v", err)
			return fiber.ErrInternalServerError
		}
	}

	return ctx.JSON(response)
}
