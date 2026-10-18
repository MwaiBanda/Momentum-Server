package meal

import (
	"log"

	applicationmeal "Momentum/application/service/meal"

	"github.com/gofiber/fiber/v3"
)

// AssignVolunteer godoc
//
//	@Summary		Assign a volunteer to a meal
//	@Description	Update a volunteered meal and assign its volunteer
//	@tags			Meals
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	meal.VolunteerAssignmentInput	true	"Assign a volunteer"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	meal.VolunteerAssignmentInput
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/meals/meal [post]
func (handler *Handler) AssignVolunteer(ctx fiber.Ctx) error {
	if handler.meals == nil {
		log.Printf("[AssignMealVolunteer] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationmeal.VolunteerAssignmentInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.meals.AssignVolunteer(ctx.Context(), input)
	if err != nil {
		return handler.mapError("AssignMealVolunteer", err)
	}
	return ctx.JSON(response)
}
