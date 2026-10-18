package meal

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// Delete godoc
//
//	@Summary		Delete a meal
//	@Description	Delete a meal by ID
//	@tags			Meals
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			mealId			path	string	true	"provide meal ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	meal.Response
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/meals/{mealId} [delete]
func (handler *Handler) Delete(ctx fiber.Ctx) error {
	if handler.meals == nil {
		log.Printf("[DeleteMeal] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.meals.Delete(ctx.Context(), ctx.Params("mealId"))
	if err != nil {
		return handler.mapError("DeleteMeal", err)
	}
	return ctx.JSON(response)
}
