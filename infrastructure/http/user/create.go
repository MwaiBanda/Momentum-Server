package user

import (
	"log"

	userService "Momentum/application/service/user"
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a user
//	@Description	Used to post a user
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body		user.CreateInput	true	"Post a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	user.Response
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/users [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	input, err := request.BindAndValidate[userService.CreateInput](ctx)
	if err != nil {
		return err
	}

	response, err := handler.users.Create(ctx.Context(), input)
	if err != nil {
		log.Printf("[CreateUser] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
