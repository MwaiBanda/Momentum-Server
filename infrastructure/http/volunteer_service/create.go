package volunteer_service

import (
	"errors"
	"log"

	applicationservice "Momentum/application/service/volunteer_service"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a volunteer service
//	@Description	Used to post a volunteer service
//	@tags			Services
//	@Param			Authorization	header	string							true	"Provide a bearer token"	example(******
//	@Param			data			body	volunteer_service.CreateInput	true	"Post a volunteer service"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	volunteer_service.CreateResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/services [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	if handler.services == nil {
		log.Printf("[CreateVolunteerService] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationservice.CreateInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.services.Create(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, applicationservice.ErrInvalidInput) {
			return fiber.ErrBadRequest
		}
		log.Printf("[CreateVolunteerService] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
