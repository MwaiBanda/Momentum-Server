package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

// DeleteUserById godoc
//
//	@Summary		Delete user information by Id
//	@Description	Delete a user's information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(******
//	@Param			userId			path		string	true	"provide user Id"
//	@Success		200				{array}		model.UserResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/users/{userId} [delete]
func (controller *Controller) DeleteUserById(context fiber.Ctx) error {
	var user model.UserResponse
	res, err := controller.PrismaClient.User.FindUnique(
		db.User.ID.Equals(context.Params("userId")),
	).Delete().Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &user); err != nil {
		fmt.Println(err)
	}

	return context.JSON(user)
}
