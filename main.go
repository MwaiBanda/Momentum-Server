package main

import (
	appconfig "Momentum/application/config"
	userService "Momentum/application/service/user"
	volunteerService "Momentum/application/service/volunteer_service"
	handlers "Momentum/controller"
	_ "Momentum/docs"
	httpHandler "Momentum/infrastructure/http"
	userHTTP "Momentum/infrastructure/http/user"
	volunteerServiceHTTP "Momentum/infrastructure/http/volunteer_service"
	firebaseClient "Momentum/infrastructure/messaging/firebase"
	database "Momentum/infrastructure/persistence/db"
	prismaClientFactory "Momentum/infrastructure/persistence/prisma"
	redisCache "Momentum/infrastructure/persistence/redis"
	userRepository "Momentum/infrastructure/persistence/repository/user"
	volunteerServiceRepository "Momentum/infrastructure/persistence/repository/volunteer_service"
	"Momentum/middleware"
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

/* PROD host			services.momentumchurch.dev */
/* DEV host				localhost:8085 */

// @title			Momentum API
// @version		v1
// @description	This is a backend service to process payments and provide some CRUD functionality
// @termsOfService	https://momentumindiana.org/statement-of-faith/
// @contact.name	Mwai Banda
// @contact.email	bandamwai@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			services.momentumchurch.dev
// @BasePath		/

func main() {
	log.Println("Starting Momentum API")
	configuration, err := appconfig.LoadConfig()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	startupContext, cancelStartup := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelStartup()

	sqlDatabase, err := database.Open(startupContext, configuration.Database.URL)
	if err != nil {
		log.Fatalf("open SQL database: %v", err)
	}
	defer func() {
		if err := database.Close(sqlDatabase); err != nil {
			log.Printf("close SQL database: %v", err)
		}
	}()

	if err := database.Migrate(startupContext, sqlDatabase); err != nil {
		log.Fatalf("migrate SQL database: %v", err)
	}

	prismaClient, err := prismaClientFactory.Open()
	if err != nil {
		log.Fatalf("open Prisma client: %v", err)
	}
	defer func() {
		if err := prismaClient.Prisma.Disconnect(); err != nil {
			log.Printf("close Prisma client: %v", err)
		}
	}()

	redisClient, err := redisCache.Open(configuration.Redis.URL)
	if err != nil {
		log.Fatalf("open Redis client: %v", err)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("close Redis client: %v", err)
		}
	}()

	var firebaseApp *firebase.App
	var messagingClient *messaging.Client
	if configuration.Firebase.ProjectID != "" {
		firebaseApp, messagingClient, err = firebaseClient.Open(
			startupContext,
			redisClient,
			configuration.Firebase.ProjectID,
		)
		if err != nil {
			log.Printf("initialize Firebase: %v", err)
		}
	} else {
		log.Println("Firebase disabled: FIREBASE_PROJECT_ID is not configured")
	}

	cacheStore := redisCache.NewCache(redisClient)
	users := userService.GetInstance(userRepository.GetInstance(sqlDatabase), cacheStore)
	volunteerServices := volunteerService.GetInstance(
		volunteerServiceRepository.GetInstance(sqlDatabase),
		cacheStore,
	)
	controller := handlers.NewController(handlers.Dependencies{
		Configuration: *configuration,
		PrismaClient:  prismaClient,
		Context:       context.Background(),
		Redis:         redisClient,
		Firebase:      firebaseApp,
		Messaging:     messagingClient,
	})
	httpHandlers := httpHandler.GetInstance(
		userHTTP.GetInstance(users),
		volunteerServiceHTTP.GetInstance(volunteerServices),
	)
	app := fiber.New()
	middleware.ConfigureAppMiddleWare(app)

	api := app.Group("/api", func(c fiber.Ctx) error {
		c.Response().Header.Set("Cache-Control", "no-cache")
		return c.Next()
	})

	api.Get("/metrics", middleware.Monitor())
	v1 := api.Group("/v1", middleware.BearerTokenAuthentication(configuration.Security.APIKey))

	v1.Post("/meals", controller.PostMeal)
	v1.Get("/meals", controller.GetAllMeals)
	v1.Delete("/meals/:mealId", controller.DeleteMealById)
	v1.Post("/meals/meal", controller.PostVolunteeredMealForMeal)
	v1.Put("/meals/meal", controller.UpdateVolunteeredMeal)

	v1.Post("/users", httpHandlers.User.Create)
	v1.Get("/users/:userId", httpHandlers.User.GetByID)
	v1.Put("/users", httpHandlers.User.Update)
	v1.Delete("/users/:userId", httpHandlers.User.Delete)

	v1.Post("/payments", controller.PostPayment)

	v1.Get("/transactions", controller.GetAllTransactions)
	v1.Get("/transactions/:userId", controller.GetTransactionsByUserId)
	v1.Delete("/transactions/:transactionId", controller.DeleteTransactionsById)
	v1.Post("/transactions", controller.PostTransaction)

	v1.Get("/sermons", controller.GetAllSermons)

	v1.Get("/services", httpHandlers.VolunteerService.ListServices)
	v1.Post("/services", httpHandlers.VolunteerService.Create)
	v1.Get("/services/:type", httpHandlers.VolunteerService.GetByType)
	v1.Put("/services/day", httpHandlers.VolunteerService.UpdateDay)

	v1.Post("/messages", controller.PostMessage)
	v1.Put("/messages", controller.UpdateMessage)
	v1.Get("/messages/unpublished", controller.GetUnpublishedMessages)
	v1.Get("/messages/:userId", controller.GetAllMessages)
	v1.Post("/messages/notes", controller.AddUserNoteToMessage)
	v1.Put("/messages/notes", controller.UpdateUserNote)
	v1.Delete("/messages/notes/:noteId", controller.DeleteNote)
	v1.Delete("/messages/:messageId", controller.DeleteMessage)

	v1.Post("/notifications", controller.PostNotification)

	v1.Get("/events", controller.GetAllEvents)

	app.Get("/swagger/*", swaggo.HandlerDefault)
	app.Get("/swagger/*", swaggo.New(swaggo.Config{
		URL:          "https://services.momentumchurch.dev/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	log.Fatal(app.Listen(":" + configuration.Port))
}
