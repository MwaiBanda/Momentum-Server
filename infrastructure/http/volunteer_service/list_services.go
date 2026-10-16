package volunteer_service

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// ListServices godoc
//
//	@Summary		Show a list of volunteer services available
//	@Description	get a list of volunteer services available
//	@Accept			json
//	@Produce		json
//	@tags			Services
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Success		200				{array}		volunteer_service.ServiceResponse
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/services [get]
func (handler *Handler) ListServices(ctx fiber.Ctx) error {
	if handler.services == nil {
		log.Printf("[ListServices] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	services, err := handler.services.ListServices(ctx.Context())
	if err != nil {
		log.Printf("[ListServices] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(services)
}
