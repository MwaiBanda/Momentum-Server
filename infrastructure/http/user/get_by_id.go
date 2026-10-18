package user

import (
	"errors"
	"log"

	userService "Momentum/application/service/user"

	"github.com/gofiber/fiber/v3"
)

// GetByID godoc
//
//	@Summary		Get user information by Id
//	@Description	get a user information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			userId			path		string	true	"provide user Id"
//	@Success		200				{object}	user.Response
//	@Failure		400				{object} http.ErrorResponse
//	@Failure		404				{object} http.ErrorResponse
//	@Failure		500				{object} http.ErrorResponse
//	@Router			/api/v1/users/{userId} [get]
func (handler *Handler) GetByID(ctx fiber.Ctx) error {
	if handler.users == nil {
		log.Printf("[GetByID] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	userID := ctx.Params("userId")
	response, err := handler.users.GetByID(ctx.Context(), userID)
	if err != nil {
		if errors.Is(err, userService.ErrNotFound) {
			return fiber.ErrNotFound
		}
		log.Printf("[GetByID] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
