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
	"github.com/lucsky/cuid"
	"github.com/redis/go-redis/v9"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

// GetAllServices godoc
//
//	@Summary		Show a list of volunteer services available
//	@Description	get a list of volunteer services available
//	@Accept			json
//	@Produce		json
//	@tags			Services
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{array}		model.Service
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/services [get]
func (controller *Controller) GetAllServices(context *fiber.Ctx) error {
	var services []model.Service
	cachedTabs, err := controller.Redis.Get(controller.Context, constants.TabsKey).Result()
	if err == redis.Nil {
		res, err := controller.PrismaClient.Service.FindMany().Exec(controller.Context)
		if err != nil {
			fmt.Println(err)
		}
	
		result, _ := json.MarshalIndent(res, "", "  ")
		if err := json.Unmarshal(result, &services); err != nil {
			fmt.Println(err)
		}
		controller.Redis.Set(controller.Context, constants.TabsKey, string(result), time.Hour*24)
	} else if err != nil {
		log.Println(err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedTabs), &services); err != nil {
			fmt.Println(err)
		}
	}
	return context.JSON(services)
}

// GetServiceByType godoc
//
//	@Summary		Get Service By Type
//	@Description	delete a Service by providing an Type
//	@Accept			json
//	@Produce		json
//	@tags			Services
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			type			path		string	true	"provide service type"
//	@Success		200				{object}	model.ServiceResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/services/{type} [get]
func (controller *Controller) GetServiceByType(context *fiber.Ctx) error {
	var service model.ServiceResponse
	res, err := controller.PrismaClient.Service.FindUnique(
		db.Service.Type.Equals(context.Params("type")),
	).With(
		db.Service.Services.Fetch().With(
			db.VolunteerService.Days.Fetch().With(
				db.VolunteeredServiceDay.User.Fetch(),
			),
		),
	).Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &service); err != nil {
		fmt.Println(err)
	}

	return context.JSON(service)
}


// PostVolunteerService godoc
//
//	@Summary		Post a voluteer service
//	@Description	Used to post a voluteer service
//	@tags			Services
//	@Param			Authorization	header	string							true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.VolunteerServiceRequest	true	"Post a meal"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.MealResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/services [post]
func (controller *Controller) PostVolunteerService(context *fiber.Ctx) error {
	volunteerRequest := new(model.VolunteerServiceRequest)
	volunteerServiceId := cuid.New()
	if err := context.BodyParser(volunteerRequest); err != nil {
		log.Panic(err.Error())
	}

	var transactions []transaction.Param
	transactions = append(transactions, controller.PrismaClient.VolunteerService.CreateOne(
		db.VolunteerService.Description.Set(volunteerRequest.Description),
		db.VolunteerService.Email.Set(volunteerRequest.Email),
		db.VolunteerService.Phone.Set(volunteerRequest.Phone),
		db.VolunteerService.Organizer.Set(volunteerRequest.Organizer),
		db.VolunteerService.ID.Set(volunteerServiceId),
		db.VolunteerService.ServiceID.Set(volunteerRequest.ServiceId),
	).Tx())
	for _, day := range volunteerRequest.Days {
		transactions = append(transactions, controller.PrismaClient.VolunteeredServiceDay.CreateOne(
			db.VolunteeredServiceDay.Notes.Set(""),
			db.VolunteeredServiceDay.Date.Set(day.Date),
			db.VolunteeredServiceDay.ServiceID.Set(volunteerServiceId),
		).Tx())
	}

	err := controller.PrismaClient.Prisma.Transaction(transactions...).Exec(controller.Context)
	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(volunteerRequest)
}


// UpdateVolunteeredMeal godoc
//
//	@Summary		Update a volunteered service day
//	@Description	Update a volunteered service day 
//	@Accept			json
//	@Produce		json
//	@tags			Services
//	@Param			Authorization	header		string						true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body		model.VolunteerServiceDay	true	"Post a day"
//	@Success		200				{object}	model.VolunteerServiceDay
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/services/day [put]
func (controller *Controller) UpdateVolunteeredDay(context *fiber.Ctx) error {
	day := new(model.VolunteerServiceDay)

	if err := context.BodyParser(day); err != nil {
		log.Panic(err.Error())
	}
	err := controller.PrismaClient.Prisma.Transaction(
		controller.PrismaClient.VolunteeredServiceDay.FindUnique(db.VolunteeredServiceDay.ID.Equals(day.Id)).Update(
			db.VolunteeredServiceDay.Notes.Set(day.Notes),
			db.VolunteeredServiceDay.UserID.Set(day.User.Id),
		).Tx(),
	).Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	return context.JSON(day)
}