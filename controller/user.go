package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// GetUserById godoc
//
//	@Summary		Get user information by Id
//	@Description	get a user information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			userId			path		string	true	"provide user Id"
//	@Success		200				{array}		model.UserResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/users/{userId} [get]
func (controller *Controller) GetUserById(context *fiber.Ctx) error {
	userId := context.Params("userId")
	userResponse := new(model.UserResponse)
	cachedUser, err := controller.Redis.Get(controller.Context, constants.UserKeyPrefix+userId).Result()
	if err == redis.Nil {
		res, err := controller.Prisma.User.FindFirst(db.UserWhereParam(
			db.User.ID.Equals(userId),
		)).Exec(controller.Context)
		if err != nil {
			fmt.Println(err)
		}
		result, _ := json.MarshalIndent(res, "", "  ")
		if err := json.Unmarshal(result, userResponse); err != nil {
			fmt.Println(err)
		}
		controller.Redis.Set(controller.Context, constants.UserKeyPrefix+userId, string(result), time.Hour*24)
	} else if err != nil {
		log.Println("[GetUserById]", err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedUser), userResponse); err != nil {
			log.Println("[GetAllSermons]", err.Error())
		}
	}

	return context.JSON(userResponse)
}

// PostUser godoc
//
//	@Summary		Post a user
//	@Description	Used to post a user
//	@tags			Users
//	@Param			Authorization	header	string				true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.UserRequest	true	"Post a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.UserRequest
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/users [post]
func (controller *Controller) PostUser(context *fiber.Ctx) error {
	userRequest := new(model.UserRequest)
	if err := context.BodyParser(userRequest); err != nil {
		log.Panic(err.Error())
	}
	res, err := controller.Prisma.User.CreateOne(
		db.User.ID.Set(userRequest.Id),
		db.User.Email.Set(userRequest.Email),
		db.User.Fullname.Set(userRequest.Email),
		db.User.Phone.Set(userRequest.Phone),
	).Exec(controller.Context)
	if err != nil {
		log.Panic(err.Error())
	}
	result, _ := json.MarshalIndent(res, "", "  ")
	controller.Redis.Set(controller.Context, constants.UserKeyPrefix+userRequest.Id, string(result), time.Hour*24)
	return context.JSON(userRequest)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Used to update a user
//	@tags			Users
//	@Param			Authorization	header	string				true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.UserRequest	true	"Post a user"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.UserRequest
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/users [put]
func (controller *Controller) UpdateUser(context *fiber.Ctx) error {
	userRequest := new(model.UserRequest)
	userResponse := new(model.UserResponse)

	if err := context.BodyParser(userRequest); err != nil {
		log.Panic(err.Error())
	}
	onUpdateUser := func() {
		res, err := controller.Prisma.User.FindUnique(
			db.User.ID.Equals(userRequest.Id),
		).Update(
			db.User.Email.Set(userRequest.Email),
			db.User.Fullname.Set(userRequest.Fullname),
			db.User.Phone.Set(userRequest.Phone),
		).Exec(controller.Context)
		if err != nil {
			log.Panic(err.Error())
		}
		result, _ := json.MarshalIndent(res, "", "  ")
		controller.Redis.Set(controller.Context, constants.UserKeyPrefix+userRequest.Id, string(result), time.Hour*24)
		if err := json.Unmarshal(result, &userResponse); err != nil {
			fmt.Println(err)
		}
	}

	cachedUser, err := controller.Redis.Get(controller.Context, constants.UserKeyPrefix+userRequest.Id).Result()
	if err := json.Unmarshal([]byte(cachedUser), userResponse); err != nil {
		log.Println("[GetAllSermons]", err.Error())
	}
	if err == redis.Nil {
		onUpdateUser()
	} else if err != nil {
		log.Panic("[UpdateUser]", err.Error())
	} else if func() bool {
		return userResponse.Phone != userRequest.Phone ||
			userResponse.Email != userRequest.Email ||
			userResponse.Fullname != userRequest.Fullname
	}() {
		log.Println("[Cache[Miss]]")
		onUpdateUser()
	}

	return context.JSON(userResponse)

}

// DeleteUserById godoc
//
//	@Summary		Delete user information by Id
//	@Description	Delete a user's information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Users
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			userId			path		string	true	"provide user Id"
//	@Success		200				{array}		model.UserResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/users/{userId} [delete]
func (controller *Controller) DeleteUserById(context *fiber.Ctx) error {
	var user model.UserResponse
	res, err := controller.Prisma.User.FindUnique(
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
