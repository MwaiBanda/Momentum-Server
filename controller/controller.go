package controller

import (
	"Momentum/prisma/db"
	"golang.org/x/net/context"
)

type Controller struct {
	client  *db.PrismaClient
	context context.Context
}

func GetControllerInstance() *Controller {
	return &Controller{}
}

func (controller *Controller) SetClient(client *db.PrismaClient) {
	controller.client = client
}

func (controller *Controller) SetContext(context context.Context) {
	controller.context = context
}
