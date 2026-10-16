package user

import (
	"log"

	userService "Momentum/application/service/user"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a user
//	@Description	Used to post a user
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Param			data			body		user.CreateInput	true	"Post a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	user.Response
//	@Failure		400	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/users [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	var input userService.CreateInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.users.Create(ctx.Context(), input)
	if err != nil {
		log.Printf("[CreateUser] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
