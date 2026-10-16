package volunteer_service

import (
	"log"

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
//	@Failure		400				{object} http.ErrorResponse
//	@Failure		404				{object} http.ErrorResponse
//	@Failure		500				{object} http.ErrorResponse
//	@Router			/api/v1/services/{type} [get]
func (handler *Handler) GetByType(ctx fiber.Ctx) error {
	if handler.services == nil {
		log.Printf("[GetServiceByType] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.services.GetByType(ctx.Context(), ctx.Params("type"))
	if err != nil {
		return handler.mapError("GetServiceByType", err)
	}

	return ctx.JSON(response)
}
