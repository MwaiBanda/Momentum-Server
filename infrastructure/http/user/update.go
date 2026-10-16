package user

import (
	"log"

	userService "Momentum/application/service/user"

	"github.com/gofiber/fiber/v3"
)

// Update godoc
//
//	@Summary		Update a user
//	@Description	Used to update a user
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Param			data			body		user.UpdateInput	true	"Update a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	user.Response
//	@Failure		400	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/users [put]
func (handler *Handler) Update(ctx fiber.Ctx) error {
	var input userService.UpdateInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.users.Update(ctx.Context(), input)
	if err != nil {
		log.Printf("[UpdateUser] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
