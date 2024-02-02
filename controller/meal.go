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
	go controller.Redis.Expire(controller.Context, constants.MealsKey, time.Second*0)

	var transactions []transaction.Param
	transactions = append(transactions, controller.PrismaClient.Meal.CreateOne(
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
		transactions = append(transactions, controller.PrismaClient.VolunteeredMeal.CreateOne(
			db.VolunteeredMeal.Description.Set(""),
			db.VolunteeredMeal.Notes.Set(""),
			db.VolunteeredMeal.Date.Set(meal.Date),
			db.VolunteeredMeal.MealID.Set(mealID),
		).Tx())
	}

	err := controller.PrismaClient.Prisma.Transaction(transactions...).Exec(controller.Context)
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
	go controller.Redis.Expire(controller.Context, constants.MealsKey, time.Second*0)
	err := controller.PrismaClient.Prisma.Transaction(
		controller.PrismaClient.VolunteeredMeal.FindUnique(db.VolunteeredMeal.ID.Equals(meal.VolunteeredMeal.Id)).Update(
			db.VolunteeredMeal.Description.Set(meal.VolunteeredMeal.Description),
			db.VolunteeredMeal.Notes.Set(meal.VolunteeredMeal.Notes),
			db.VolunteeredMeal.UserID.Set(meal.VolunteeredMeal.User.Id),
		).Tx(),
		controller.PrismaClient.MealParticipant.UpsertOne(
			db.MealParticipant.IDUserIDMealID(db.MealParticipant.ID.Equals(meal.VolunteeredMeal.Id), db.MealParticipant.UserID.Equals(meal.MealId), db.MealParticipant.MealID.Equals(meal.VolunteeredMeal.User.Id)),
		).Create(
			db.MealParticipant.Meal.Link(db.Meal.ID.Equals(meal.MealId)),
			db.MealParticipant.User.Link(db.User.ID.Equals(meal.VolunteeredMeal.User.Id)),
		).Update().Tx(),
	).Exec(controller.Context)
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
	cachedMeal, err := controller.Redis.Get(controller.Context, constants.MealsKey).Result()
	if err == redis.Nil {
		res, err := controller.PrismaClient.Meal.FindMany().With(
			db.Meal.User.Fetch(),
		).With(
			db.Meal.Participants.Fetch().With(
				db.MealParticipant.User.Fetch(),
			),
		).With(
			db.Meal.Meals.Fetch().With(
				db.VolunteeredMeal.User.Fetch(),
			),
		).Exec(controller.Context)

		if err != nil {
			fmt.Println(err)
		}

		result, _ := json.MarshalIndent(res, "", "  ")
		if result == nil {
			controller.Redis.Set(controller.Context, constants.MealsKey, string(result), time.Hour*24)
		}
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

// UpdateMeal godoc
//
//	@Summary		Update a meal
//	@Description	Update a meal
//	@Accept			json
//	@Produce		json
//	@tags			Meals
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{object}	model.MealRequest
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/meals [put]
func (controller *Controller) UpdateMeal(context *fiber.Ctx) error {
	meal := new(model.MealRequest)
	go controller.Redis.Expire(controller.Context, constants.MealsKey, time.Second*0)

	if err := context.BodyParser(meal); err != nil {
		log.Panic(err.Error())
	}
	err := controller.PrismaClient.Prisma.Transaction(
		controller.PrismaClient.Meal.FindUnique(db.Meal.ID.Equals(meal.Id)).Update(
			db.Meal.Reason.Set(meal.Reason),
			db.Meal.Email.Set(meal.Email),
			db.Meal.Phone.Set(meal.Phone),
			db.Meal.City.Set(meal.City),
			db.Meal.NumOfAdults.Set(meal.NumOfAdults),
			db.Meal.NumberOfKids.Set(meal.NumberOfKids),
			db.Meal.Street.Set(meal.Street),
			db.Meal.PreferredTime.Set(meal.PreferredTime),
			db.Meal.Favourites.Set(meal.Favourites),
			db.Meal.LeastFavourites.Set(meal.LeastFavourites),
			db.Meal.Allergies.Set(meal.Allergies),
			db.Meal.Instructions.Set(meal.Instructions),
			db.Meal.Recipient.Set(meal.Recipient),
		).Tx(),
	).Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	return context.JSON(meal)
}

// DeleteMealById godoc
//
//	@Summary		Delete a meal
//	@Description	delete a meal by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Meals
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			mealId			path		string	true	"provide meal Id"
//	@Success		200				{object}	model.MealResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/meals/{mealId} [delete]
func (controller *Controller) DeleteMealById(context *fiber.Ctx) error {
	var meal model.MealResponse
	res, err := controller.PrismaClient.Meal.FindUnique(
		db.Meal.ID.Equals(context.Params("mealId")),
	).Delete().Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &meal); err != nil {
		fmt.Println(err)
	}

	return context.JSON(meal)
}


// UpdateVolunteeredMeal godoc
//
//	@Summary		Update a volunteered meal
//	@Description	Update a volunteered meal 
//	@Accept			json
//	@Produce		json
//	@tags			Meals
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{object}	model.VolunteeredMeal
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/meals/meal [put]
func (controller *Controller) UpdateVolunteeredMeal(context *fiber.Ctx) error {
	meal := new(model.VolunteeredMeal)
	go controller.Redis.Expire(controller.Context, constants.MealsKey, time.Second*0)

	if err := context.BodyParser(meal); err != nil {
		log.Panic(err.Error())
	}
	err := controller.PrismaClient.Prisma.Transaction(
		controller.PrismaClient.VolunteeredMeal.FindUnique(db.VolunteeredMeal.ID.Equals(meal.Id)).Update(
			db.VolunteeredMeal.Description.Set(meal.Description),
			db.VolunteeredMeal.Notes.Set(meal.Notes),
		).Tx(),
	).Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	return context.JSON(meal)
}