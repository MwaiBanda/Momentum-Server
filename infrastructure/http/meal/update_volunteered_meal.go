package meal

import (
	"log"

	applicationmeal "Momentum/application/service/meal"

	"github.com/gofiber/fiber/v3"
)

// UpdateVolunteeredMeal godoc
//
//	@Summary		Update a volunteered meal
//	@Description	Update a volunteered meal's description and notes
//	@tags			Meals
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	meal.UpdateVolunteeredMealInput	true	"Update a volunteered meal"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	meal.UpdateVolunteeredMealInput
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/meals/meal [put]
func (handler *Handler) UpdateVolunteeredMeal(ctx fiber.Ctx) error {
	if handler.meals == nil {
		log.Printf("[UpdateVolunteeredMeal] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationmeal.UpdateVolunteeredMealInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.meals.UpdateVolunteeredMeal(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateVolunteeredMeal", err)
	}
	return ctx.JSON(response)
}
