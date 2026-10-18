package meal

import (
	"log"

	applicationmeal "Momentum/application/service/meal"
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
)

// Update godoc
//
//	@Summary		Update a meal
//	@Description	Update a meal
//	@tags			Meals
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	meal.Input	true	"Update a meal"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	meal.UpdateResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/meals [put]
func (handler *Handler) Update(ctx fiber.Ctx) error {
	if handler.meals == nil {
		log.Printf("[UpdateMeal] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	input, err := request.BindAndValidate[applicationmeal.Input](ctx)
	if err != nil {
		return err
	}

	response, err := handler.meals.Update(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateMeal", err)
	}
	return ctx.JSON(response)
}
