package meal

import (
	"log"

	applicationmeal "Momentum/application/service/meal"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a meal
//	@Description	Create a meal and its requested volunteer dates
//	@tags			Meals
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	meal.Input	true	"Post a meal"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	meal.Input
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/meals [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	if handler.meals == nil {
		log.Printf("[CreateMeal] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationmeal.Input
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.meals.Create(ctx.Context(), input)
	if err != nil {
		return handler.mapError("CreateMeal", err)
	}
	return ctx.JSON(response)
}
