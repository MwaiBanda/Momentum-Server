package volunteer_service

import (
	"log"

	applicationservice "Momentum/application/service/volunteer_service"
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a volunteer service
//	@Description	Used to post a volunteer service
//	@tags			Services
//	@Param			Authorization	header	string							true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	volunteer_service.CreateInput	true	"Post a volunteer service"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	volunteer_service.CreateResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/services [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	if handler.services == nil {
		log.Printf("[CreateVolunteerService] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	input, err := request.BindAndValidate[applicationservice.CreateInput](ctx)
	if err != nil {
		return err
	}

	response, err := handler.services.Create(ctx.Context(), input)
	if err != nil {
		return handler.mapError("CreateVolunteerService", err)
	}

	return ctx.JSON(response)
}
