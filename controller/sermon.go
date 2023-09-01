package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"time"
)

// GetAllSermons godoc
//
//	@Summary		Show a list of sermons
//	@Description	get a list of sermons
//	@Accept			json
//	@Produce		json
//	@tags			Sermons
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			page			query		number	false	"Provide a page number"
//	@Success		200				{array}		model.SermonResponse
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
//	@Router			/api/v1/sermons [get]
func (controller *Controller) GetAllSermons(context *fiber.Ctx) error {
	pageNumber := context.Query("page", "1")
	sermons := new(model.SermonResponse)
	var cachedSermons, err = controller.redis.Get(controller.context, constants.SermonKeyPrefix+pageNumber).Result()
	if err == redis.Nil {
		resp, err := controller.httpClient.Get("https://api.sermoncloud.com/momentum-church-1/sermons?page=" + pageNumber)
		if err != nil {
			log.Println("[GetAllSermons]", err.Error())
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, sermons); err != nil {
			log.Println("[GetAllSermons]", err.Error())
		}
		if err = controller.redis.Set(controller.context, constants.SermonKeyPrefix+pageNumber, string(body), time.Hour*6).Err(); err != nil {
			log.Println("[GetAllSermons]", err.Error())
		}
	} else if err != nil {
		log.Println("[GetAllSermons]", err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedSermons), sermons); err != nil {
			log.Println("[GetAllSermons]", err.Error())
		}
	}
	return context.JSON(sermons)
}
