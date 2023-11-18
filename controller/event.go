package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
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
	url := constants.PlannnigCenterUrl + "?per_page=100&where[ends_at][gte]=" + time.Now().Add(-24*time.Hour).Format(time.RFC3339)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("New Request", err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("PLANNING_CENTER_TOKEN"))

	resp, err := controller.HttpClient.Do(req)
	if err != nil {
		log.Panic("Request", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Request", err)
	}
	eventResponse := new(model.EventReponse)

	if err := json.Unmarshal(data, &eventResponse); err != nil {
		log.Println(err)
	}

	return context.JSON(eventResponse)
}
