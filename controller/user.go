package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// GetUser godoc
//
//	@Summary		Get user information by Id
//	@Description	get a user information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Success		200	{array}		model.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/user [get]
func (controller *Controller) GetUser(context *fiber.Ctx) error {
	var user model.User
	res, err := controller.prisma.User.FindFirst(db.UserWhereParam(
		db.User.ID.Equals(context.Params("userId")),
	)).Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &user); err != nil {
		fmt.Println(err)
	}

	return context.JSON(user)
}
