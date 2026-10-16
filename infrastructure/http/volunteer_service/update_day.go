package volunteer_service

import (
	"log"

	applicationservice "Momentum/application/service/volunteer_service"

	"github.com/gofiber/fiber/v3"
)

// UpdateDay godoc
//
//	@Summary		Update a volunteered service day
//	@Description	Update a volunteered service day
//	@Accept			json
//	@Produce		json
//	@tags			Services
//	@Param			Authorization	header	string							true	"Provide a bearer token"	example(******
//	@Param			data			body	volunteer_service.UpdateDayInput	true	"Update a day"
//	@Success		200	{object}	volunteer_service.UpdateDayResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/services/day [put]
func (handler *Handler) UpdateDay(ctx fiber.Ctx) error {
	if handler.services == nil {
		log.Printf("[UpdateServiceDay] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationservice.UpdateDayInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.services.UpdateDay(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateServiceDay", err)
	}

	return ctx.JSON(response)
}
