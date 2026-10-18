package meal

import (
	"log"

	applicationmeal "Momentum/application/service/meal"
	"Momentum/infrastructure/http/request"

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
//	@Success		200	{object}	meal.UpdateVolunteeredMealResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/meals/meal [put]
func (handler *Handler) UpdateVolunteeredMeal(ctx fiber.Ctx) error {
	if handler.meals == nil {
		log.Printf("[UpdateVolunteeredMeal] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	input, err := request.BindAndValidate[applicationmeal.UpdateVolunteeredMealInput](ctx)
	if err != nil {
		return err
	}

	response, err := handler.meals.UpdateVolunteeredMeal(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateVolunteeredMeal", err)
	}
	return ctx.JSON(response)
}
