package controller

import (
	appconfig "Momentum/application/config"
	"Momentum/prisma/db"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type Controller struct {
	Configuration appconfig.Config
	PrismaClient  *db.PrismaClient
	Context       context.Context
	Redis         *redis.Client
}

type Dependencies struct {
	Configuration appconfig.Config
	PrismaClient  *db.PrismaClient
	Context       context.Context
	Redis         *redis.Client
}

func NewController(dependencies Dependencies) *Controller {
	controllerContext := dependencies.Context
	if controllerContext == nil {
		controllerContext = context.Background()
	}

	return &Controller{
		Configuration: dependencies.Configuration,
		PrismaClient:  dependencies.PrismaClient,
		Context:       controllerContext,
		Redis:         dependencies.Redis,
	}
}
