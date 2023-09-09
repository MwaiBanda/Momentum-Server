package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lucsky/cuid"
	"github.com/redis/go-redis/v9"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
	"log"
	"time"
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
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/meals [post]
func (controller *Controller) PostMeal(context *fiber.Ctx) error {
	mealRequest := new(model.MealRequest)
	mealID := cuid.New()

	if err := context.BodyParser(mealRequest); err != nil {
		log.Panic(err.Error())
	}
	go controller.redis.Expire(controller.context, "meal", time.Second*0)

	var transactions []transaction.Param
	transactions = append(transactions, controller.prisma.Meal.CreateOne(
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
	).Tx())
	for _, meal := range mealRequest.Meals {
		transactions = append(transactions, controller.prisma.VolunteeredMeal.CreateOne(
			db.VolunteeredMeal.Description.Set(""),
			db.VolunteeredMeal.Notes.Set(""),
			db.VolunteeredMeal.Date.Set(meal.Date),
			db.VolunteeredMeal.MealID.Set(mealID),
		).Tx())
	}

	err := controller.prisma.Prisma.Transaction(transactions...).Exec(controller.context)
	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(mealRequest)
}

// PostVolunteeredMealForMeal godoc
//
//	@Summary		Post a meal participant
//	@Description	Used to post a meal participant
//	@tags			Meals
//	@Param			Authorization	header	string							true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.VolunteeredMealRequest	true	"Post participant information"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.VolunteeredMealRequest
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/meals/meal [post]
func (controller *Controller) PostVolunteeredMealForMeal(context *fiber.Ctx) error {
	meal := new(model.VolunteeredMealRequest)
	if err := context.BodyParser(meal); err != nil {
		log.Panic(err.Error())
	}
	go controller.redis.Expire(controller.context, "meal", time.Second*0)
	err := controller.prisma.Prisma.Transaction(
		controller.prisma.VolunteeredMeal.FindUnique(db.VolunteeredMeal.ID.Equals(meal.VolunteeredMeal.Id)).Update(
			db.VolunteeredMeal.Description.Set(meal.VolunteeredMeal.Description),
			db.VolunteeredMeal.Notes.Set(meal.VolunteeredMeal.Notes),
			db.VolunteeredMeal.UserID.Set(meal.VolunteeredMeal.User.Id),
		).Tx(),
		controller.prisma.MealParticipant.UpsertOne(
			db.MealParticipant.IDUserIDMealID(db.MealParticipant.ID.Equals(meal.VolunteeredMeal.Id), db.MealParticipant.UserID.Equals(meal.MealId), db.MealParticipant.MealID.Equals(meal.VolunteeredMeal.User.Id)),
		).Create(
			db.MealParticipant.Meal.Link(db.Meal.ID.Equals(meal.MealId)),
			db.MealParticipant.User.Link(db.User.ID.Equals(meal.VolunteeredMeal.User.Id)),
		).Update().Tx(),
	).Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	return context.JSON(meal)
}

// GetAllMeals godoc
//
//	@Summary		Get all Meals
//	@Description	Used to get all meals
//	@tags			Meals
//
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.MealResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/meals [get]
func (controller *Controller) GetAllMeals(context *fiber.Ctx) error {
	mealResponse := new([]model.MealResponse)
	cachedMeal, err := controller.redis.Get(controller.context, constants.MealsKey).Result()
	if err == redis.Nil {
		res, err := controller.prisma.Meal.FindMany().With(
			db.Meal.User.Fetch(),
		).With(
			db.Meal.Participants.Fetch().With(
				db.MealParticipant.User.Fetch(),
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
		controller.redis.Set(controller.context, constants.MealsKey, string(result), time.Hour*6)
		if err := json.Unmarshal(result, &mealResponse); err != nil {
			fmt.Println(err)
		}
	} else if err != nil {
		log.Println(err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedMeal), &mealResponse); err != nil {
			fmt.Println(err)
		}
	}

	return context.JSON(mealResponse)
}
