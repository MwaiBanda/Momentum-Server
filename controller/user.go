package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
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
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
//	@Router			/api/v1/users/{userId} [get]
func (controller *Controller) GetUserById(context *fiber.Ctx) error {
	userId := context.Params("userId")
	userResponse := new(model.UserResponse)
	cachedUser, err := controller.redis.Get(controller.context, "userResponse-"+userId).Result()
	if err == redis.Nil {
		res, err := controller.prisma.User.FindFirst(db.UserWhereParam(
			db.User.ID.Equals(userId),
		)).Exec(controller.context)
		if err != nil {
			fmt.Println(err)
		}
		result, _ := json.MarshalIndent(res, "", "  ")
		if err := json.Unmarshal(result, userResponse); err != nil {
			fmt.Println(err)
		}
		controller.redis.Set(controller.context, "user-"+userId, string(result), time.Hour*24)
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
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/users [post]
func (controller *Controller) PostUser(context *fiber.Ctx) error {
	userRequest := new(model.UserRequest)
	if err := context.BodyParser(userRequest); err != nil {
		log.Panic(err.Error())
	}
	res, err := controller.prisma.User.CreateOne(
		db.User.ID.Set(userRequest.Id),
		db.User.Email.Set(userRequest.Email),
		db.User.Fullname.Set(userRequest.Email),
		db.User.Phone.Set(userRequest.Phone),
	).Exec(controller.context)
	if err != nil {
		log.Panic(err.Error())
	}
	result, _ := json.MarshalIndent(res, "", "  ")
	controller.redis.Set(controller.context, "userRequest-"+userRequest.Id, string(result), time.Hour*24)
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
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/users [put]
func (controller *Controller) UpdateUser(context *fiber.Ctx) error {
	userRequest := new(model.UserRequest)
	userResponse := new(model.UserResponse)

	if err := context.BodyParser(userRequest); err != nil {
		log.Panic(err.Error())
	}

	cachedUser, err := controller.redis.Get(controller.context, "user-"+userRequest.Id).Result()
	if err != nil {
		log.Panic("[UpdateUser]", err.Error())
	}
	if err := json.Unmarshal([]byte(cachedUser), userResponse); err != nil {
		log.Println("[GetAllSermons]", err.Error())
	}

	if func() bool {
		return userResponse.Phone != userRequest.Phone ||
			userResponse.Email != userRequest.Email ||
			userResponse.Fullname != userRequest.Fullname
	}() {
		log.Println("[Cache[Miss]]")
		res, err := controller.prisma.User.FindUnique(
			db.User.ID.Equals(userRequest.Id),
		).Update(
			db.User.Email.Set(userRequest.Email),
			db.User.Fullname.Set(userRequest.Fullname),
			db.User.Phone.Set(userRequest.Phone),
		).Exec(controller.context)
		if err != nil {
			log.Panic(err.Error())
		}
		result, _ := json.MarshalIndent(res, "", "  ")
		controller.redis.Set(controller.context, "userRequest-"+userRequest.Id, string(result), time.Hour*24)
		if err := json.Unmarshal(result, &userResponse); err != nil {
			fmt.Println(err)
		}
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
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
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
