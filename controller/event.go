package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// GetAllEvents godoc
//
//	@Summary		Get all Calender Events
//	@Description	Used to get all events scheduled on the church calender
//	@tags			Events
//
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.EventReponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/events [get]
func (controller *Controller) GetAllEvents(context *fiber.Ctx) error {
	eventResponse := new(model.EventReponse)
	cachedEvents, err := controller.Redis.Get(controller.Context, constants.EventsKey).Result()	
	if err == redis.Nil {
		data, err := controller.EventRequest().ReadData()
		if err != nil {
			log.Panic("Request", err)
		}
		if err := json.Unmarshal(data, &eventResponse); err != nil {
			log.Println(err)
		}
		controller.Redis.Set(controller.Context, constants.EventsKey, string(data), time.Hour*2)
	} else if err != nil {
		log.Println(err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedEvents), &eventResponse); err != nil {
			fmt.Println(err)
		}
	}
	return context.JSON(eventResponse)
}
