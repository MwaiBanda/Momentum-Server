package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

// GetUserById godoc
//
//	@Summary		Get user information by Id
//	@Description	get a user information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			userId	path		string	true	"provide user Id"
//	@Success		200		{array}		model.UserResponse
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/api/v1/users/{userId} [get]
func (controller *Controller) GetUserById(context *fiber.Ctx) error {
	var user model.UserResponse
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

// PostUser godoc
//
//	@Summary		Post a user
//	@Description	Used to post a user
//	@tags			Users
//	@Param			data	body	model.UserRequest	true	"Post a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.UserRequest
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/users [post]
func (controller *Controller) PostUser(context *fiber.Ctx) error {
	user := new(model.UserRequest)
	if err := context.BodyParser(user); err != nil {
		log.Panic(err.Error())
	}
	_, err := controller.prisma.User.CreateOne(
		db.User.ID.Set(user.Id),
		db.User.Email.Set(user.Email),
		db.User.Fullname.Set(user.Email),
		db.User.Phone.Set(user.Phone),
	).Exec(controller.context)
	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(user)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Used to update a user
//	@tags			Users
//	@Param			data	body	model.UserRequest	true	"Post a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.UserRequest
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/users [put]
func (controller *Controller) UpdateUser(context *fiber.Ctx) error {
	user := new(model.UserRequest)
	if err := context.BodyParser(user); err != nil {
		log.Panic(err.Error())
	}
	_, err := controller.prisma.User.UpsertOne(
		db.User.ID.Equals(user.Id),
	).Update(
		db.User.Email.Set(user.Email),
		db.User.Fullname.Set(user.Email),
		db.User.Phone.Set(user.Phone),
	).Exec(controller.context)
	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(user)
}

// DeleteUserById godoc
//
//	@Summary		Delete user information by Id
//	@Description	Delete a user's information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			userId	path		string	true	"provide user Id"
//	@Success		200		{array}		model.UserResponse
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/api/v1/users/{userId} [delete]
func (controller *Controller) DeleteUserById(context *fiber.Ctx) error {
	var user model.UserResponse
	res, err := controller.prisma.User.FindUnique(
		db.User.ID.Equals(context.Params("userId")),
	).Delete().Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &user); err != nil {
		fmt.Println(err)
	}

	return context.JSON(user)
}
