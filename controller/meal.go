package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

// PostMeal godoc
//
//	@Summary		Post a Meal
//	@Description	Used to post a meal
//	@tags			Meals
//	@Param			data	body	model.MealRequest	true	"Post a meal"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.MealRequest
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/meals [post]
func (controller *Controller) PostMeal(context *fiber.Ctx) error {
	meal := new(model.MealRequest)
	if err := context.BodyParser(meal); err != nil {
		log.Panic(err.Error())
	}
	_, err := controller.prisma.Meal.CreateOne(
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
		db.Meal.UserID.Set(meal.UserId),
	).Exec(controller.context)

	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(meal)
}

// PostMealParticipant godoc
//
//	@Summary		Post a meal participate
//	@Description	Used to post a meal participate
//	@tags			Meals
//	@Param			data	body	model.ParticipantRequest	true	"Post participant information"
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
//	@Param			data	body	model.VolunteeredRequest	true	"Post participant information"
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
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.MealResponse
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/meals [get]
func (controller *Controller) GetMeals(context *fiber.Ctx) error {
	var meal []model.MealResponse
	res, err := controller.prisma.Meal.FindMany().With(db.Meal.User.Fetch()).With(db.Meal.Participants.Fetch().With(db.UsersOnMeals.User.Fetch())).With(db.Meal.Meals.Fetch()).Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &meal); err != nil {
		fmt.Println(err)
	}
	return context.JSON(meal)
}
