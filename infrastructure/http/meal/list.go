package meal

import "github.com/gofiber/fiber/v3"

// List godoc
//
//	@Summary		Get all meals
//	@Description	Get all meals and their volunteer assignments
//	@tags			Meals
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	meal.Response
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/meals [get]
func (handler *Handler) List(ctx fiber.Ctx) error {
	if handler.meals == nil {
		return handler.mapError("ListMeals", fiber.ErrInternalServerError)
	}

	response, err := handler.meals.List(ctx.Context())
	if err != nil {
		return handler.mapError("ListMeals", err)
	}
	return ctx.JSON(response)
}
