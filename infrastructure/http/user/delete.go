package user

import (
	"errors"
	"log"

	userService "Momentum/application/service/user"

	"github.com/gofiber/fiber/v3"
)

// Delete godoc
//
//	@Summary		Delete user information by Id
//	@Description	Delete a user's information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Param			userId			path		string	true	"provide user Id"
//	@Success		200				{object}	user.Response
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/users/{userId} [delete]
func (handler *Handler) Delete(ctx fiber.Ctx) error {
	response, err := handler.users.Delete(ctx.Context(), ctx.Params("userId"))
	if err != nil {
		if errors.Is(err, userService.ErrNotFound) {
			return fiber.ErrNotFound
		}
		log.Printf("[DeleteUser] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
