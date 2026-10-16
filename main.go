package main

import (
	userService "Momentum/application/service/user"
	"Momentum/constants"
	handlers "Momentum/controller"
	_ "Momentum/docs"
	httpHandler "Momentum/infrastructure/http"
	userHTTP "Momentum/infrastructure/http/user"
	database "Momentum/infrastructure/persistence/db"
	redisCache "Momentum/infrastructure/persistence/redis"
	userRepository "Momentum/infrastructure/persistence/repository/user"
	"Momentum/middleware"
	prisma "Momentum/prisma/db"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/option"
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

func openPrismaClient() (*prisma.PrismaClient, error) {
	prismaClient := prisma.NewClient()

	var err error
	for attempt := 1; attempt <= 5; attempt++ {
		err = prismaClient.Prisma.Connect()
		if err == nil {
			log.Println("Connected to Prisma")
			return prismaClient, nil
		}
		log.Printf("Prisma connect attempt %d/5 failed: %v", attempt, err)
		time.Sleep(time.Duration(attempt) * time.Second)
	}

	_ = prismaClient.Prisma.Disconnect()
	return nil, fmt.Errorf("connect to Prisma after retries: %w", err)
}

func openRedisClient() (*redis.Client, error) {
	options, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, fmt.Errorf("parse Redis URL: %w", err)
	}

	client := redis.NewClient(options)
	log.Println("Connected to Redis")
	return client, nil
}

func initializeFirebase(
	ctx context.Context,
	redisClient *redis.Client,
) (*firebase.App, *messaging.Client, error) {
	firebaseToken, err := redisClient.Get(ctx, constants.FirebaseServiceTokenKey).Result()
	if err != nil {
		return nil, nil, fmt.Errorf("read Firebase service token from Redis: %w", err)
	}

	credentials := option.WithCredentialsJSON([]byte(firebaseToken))
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	firebaseApp, err := firebase.NewApp(ctx, config, credentials)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize Firebase: %w", err)
	}

	messagingClient, err := firebaseApp.Messaging(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize Firebase messaging: %w", err)
	}

	log.Println("Connected to Firebase")
	return firebaseApp, messagingClient, nil
}

func main() {
	log.Println("Starting Momentum API")
	if err := godotenv.Load(); err != nil && os.Getenv("ENVIRONMENT") == constants.Development {
		log.Fatalf("load environment: %v", err)
	}

	startupContext, cancelStartup := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelStartup()

	sqlDatabase, err := database.Open(startupContext, os.Getenv("DATABASE_URL"))
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

	prismaClient, err := openPrismaClient()
	if err != nil {
		log.Fatalf("open Prisma client: %v", err)
	}
	defer func() {
		if err := prismaClient.Prisma.Disconnect(); err != nil {
			log.Printf("close Prisma client: %v", err)
		}
	}()

	redisClient, err := openRedisClient()
	if err != nil {
		log.Fatalf("open Redis client: %v", err)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("close Redis client: %v", err)
		}
	}()

	firebaseApp, messagingClient, err := initializeFirebase(startupContext, redisClient)
	if err != nil {
		log.Printf("initialize Firebase: %v", err)
	}

	users := userService.GetInstance(
		userRepository.GetInstance(sqlDatabase),
		redisCache.NewCache(redisClient),
	)
	controller := handlers.NewController(handlers.Dependencies{
		PrismaClient: prismaClient,
		Context:      context.Background(),
		Redis:        redisClient,
		Firebase:     firebaseApp,
		Messaging:    messagingClient,
	})
	httpHandlers := httpHandler.GetInstance(userHTTP.GetInstance(users))
	app := fiber.New()
	port := func() string {
		if len(os.Getenv("PORT")) > 0 {
			return os.Getenv("PORT")
		} else {
			return "8080"
		}
	}()

	middleware.ConfigureAppMiddleWare(app)

	api := app.Group("/api", func(c fiber.Ctx) error {
		c.Response().Header.Set("Cache-Control", "no-cache")
		return c.Next()
	})

	api.Get("/metrics", middleware.Monitor())
	v1 := api.Group("/v1", middleware.BearerTokenAuthentication())

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

	v1.Get("/services", controller.GetAllServices)
	v1.Post("/services", controller.PostVolunteerService)
	v1.Get("/services/:type", controller.GetServiceByType)
	v1.Put("/services/day", controller.UpdateVolunteeredDay)

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

	log.Fatal(app.Listen(":" + port))
}
