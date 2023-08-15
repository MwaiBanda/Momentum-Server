package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lucsky/cuid"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
	"log"
)

// PostMeal godoc
//
//	@Summary		Post a Meal
//	@Description	Used to post a meal
//	@tags			Meals
//	@Param			Authorization	header	string				true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.MealRequest	true	"Post a meal"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.MealResponse
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/meals [post]
func (controller *Controller) PostMeal(context *fiber.Ctx) error {
	mealRequest := new(model.MealRequest)
	mealID := cuid.New()

	if err := context.BodyParser(mealRequest); err != nil {
		log.Panic(err.Error())
	}
	var volunteered []transaction.Param
	for _, meal := range mealRequest.Meals {
		volunteered = append(volunteered, controller.prisma.VolunteeredMeal.CreateOne(
			db.VolunteeredMeal.Description.Set(""),
			db.VolunteeredMeal.Notes.Set(""),
			db.VolunteeredMeal.Date.Set(meal.Date),
			db.VolunteeredMeal.MealID.Set(mealID),
		).Tx())
	}

	err := controller.prisma.Prisma.Transaction(volunteered...).Exec(controller.context)
	if err != nil {
		log.Panic(err.Error())
	}
	mealResponse := new(model.MealResponse)

	res, _ := controller.prisma.Meal.CreateOne(
		db.Meal.Reason.Set(mealRequest.Reason),
		db.Meal.Email.Set(mealRequest.Email),
		db.Meal.Phone.Set(mealRequest.Phone),
		db.Meal.City.Set(mealRequest.City),
		db.Meal.NumOfAdults.Set(mealRequest.NumOfAdults),
		db.Meal.NumberOfKids.Set(mealRequest.NumberOfKids),
		db.Meal.Street.Set(mealRequest.Street),
		db.Meal.PreferredTime.Set(mealRequest.PreferredTime),
		db.Meal.Favourites.Set(mealRequest.Favourites),
		db.Meal.LeastFavourites.Set(mealRequest.LeastFavourites),
		db.Meal.Allergies.Set(mealRequest.Allergies),
		db.Meal.Instructions.Set(mealRequest.Instructions),
		db.Meal.Recipient.Set(mealRequest.Recipient),
		db.Meal.UserID.Set(mealRequest.UserId),
		db.Meal.ID.Set(mealID),
		db.Meal.Meals.Link(db.VolunteeredMeal.MealID.Equals(mealID)),
	).Exec(controller.context)

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &mealResponse); err != nil {
		fmt.Println(err)
	}

	return context.JSON(mealResponse)
}

// PostMealParticipant godoc
//
//	@Summary		Post a meal participate
//	@Description	Used to post a meal participate
//	@tags			Meals
//	@Param			Authorization	header	string						true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.ParticipantRequest	true	"Post participant information"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.ParticipantRequest
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/meals/participant [post]
func (controller *Controller) PostMealParticipant(context *fiber.Ctx) error {
	meal := new(model.ParticipantRequest)
	if err := context.BodyParser(meal); err != nil {
		log.Panic(err.Error())
	}
	_, err := controller.prisma.UsersOnMeals.CreateOne(
		db.UsersOnMeals.Meal.Link(db.Meal.ID.Equals(meal.MealId)),
		db.UsersOnMeals.User.Link(db.User.ID.Equals(meal.UserId)),
	).Exec(controller.context)

	if err != nil {
		log.Panic(err.Error())
	}
	return context.JSON(meal)
}

// PostMealForMeal godoc
//
//	@Summary		Post a meal participate
//	@Description	Used to post a meal participate
//	@tags			Meals
//	@Param			Authorization	header	string						true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.VolunteeredRequest	true	"Post participant information"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.VolunteeredRequest
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/meals/meal [post]
func (controller *Controller) PostMealForMeal(context *fiber.Ctx) error {
	meal := new(model.VolunteeredRequest)
	if err := context.BodyParser(meal); err != nil {
		log.Panic(err.Error())
	}
	_, err := controller.prisma.VolunteeredMeal.CreateOne(
		db.VolunteeredMeal.Description.Set(meal.Description),
		db.VolunteeredMeal.Notes.Set(meal.Notes),
		db.VolunteeredMeal.MealID.Set(meal.MealId),
		db.VolunteeredMeal.UserID.Set(meal.UserId),
	).Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}
	return context.JSON(meal)
}

// GetMeals godoc
//
//	@Summary		Get all Meals
//	@Description	Used to get all meals
//	@tags			Meals
//
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.MealResponse
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/meals [get]
func (controller *Controller) GetMeals(context *fiber.Ctx) error {
	var meal []model.MealResponse
	res, err := controller.prisma.Meal.FindMany().With(
		db.Meal.User.Fetch(),
	).With(
		db.Meal.Participants.Fetch().With(
			db.UsersOnMeals.User.Fetch(),
		),
	).With(
		db.Meal.Meals.Fetch().With(
			db.VolunteeredMeal.User.Fetch(),
		),
	).Exec(controller.context)

	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &meal); err != nil {
		fmt.Println(err)
	}
	return context.JSON(meal)
}
