package controller

import (
	"Momentum/model"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllServices godoc
//
//	@Summary		Show a list of volunteer services available
//	@Description	get a list of volunteer services available
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{array}		model.Service
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/services [get]
func (controller *Controller) GetAllServices(context *fiber.Ctx) error {
	var services []model.Service
	res, err := controller.PrismaClient.Service.FindMany().Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &services); err != nil {
		fmt.Println(err)
	}
	return context.JSON(services)
}
