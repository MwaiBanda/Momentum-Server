package main

import (
	appconfig "Momentum/application/config"
	eventService "Momentum/application/service/event"
	mealService "Momentum/application/service/meal"
	messageService "Momentum/application/service/message"
	notificationService "Momentum/application/service/notification"
	paymentService "Momentum/application/service/payment"
	sermonService "Momentum/application/service/sermon"
	transactionService "Momentum/application/service/transaction"
	userService "Momentum/application/service/user"
	volunteerService "Momentum/application/service/volunteer_service"
	_ "Momentum/docs"
	httpHandler "Momentum/infrastructure/http"
	eventHTTP "Momentum/infrastructure/http/event"
	mealHTTP "Momentum/infrastructure/http/meal"
	messageHTTP "Momentum/infrastructure/http/message"
	httpMiddleware "Momentum/infrastructure/http/middleware"
	notificationHTTP "Momentum/infrastructure/http/notification"
	paymentHTTP "Momentum/infrastructure/http/payment"
	sermonHTTP "Momentum/infrastructure/http/sermon"
	transactionHTTP "Momentum/infrastructure/http/transaction"
	userHTTP "Momentum/infrastructure/http/user"
	volunteerServiceHTTP "Momentum/infrastructure/http/volunteer_service"
	httpIntegration "Momentum/infrastructure/integration/httpclient"
	planningcenter "Momentum/infrastructure/integration/planningcenter"
	sermoncloud "Momentum/infrastructure/integration/sermoncloud"
	stripeIntegration "Momentum/infrastructure/integration/stripe"
	emailMessaging "Momentum/infrastructure/messaging/email"
	firebaseClient "Momentum/infrastructure/messaging/firebase"
	database "Momentum/infrastructure/persistence/db"
	redisCache "Momentum/infrastructure/persistence/redis"
	mealRepository "Momentum/infrastructure/persistence/repository/meal"
	messageRepository "Momentum/infrastructure/persistence/repository/message"
	transactionRepository "Momentum/infrastructure/persistence/repository/transaction"
	userRepository "Momentum/infrastructure/persistence/repository/user"
	volunteerServiceRepository "Momentum/infrastructure/persistence/repository/volunteer_service"
	"context"
	"log"
	"time"

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
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Bearer token required for protected API routes.

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

	redisClient, err := redisCache.Open(configuration.Redis.URL)
	if err != nil {
		log.Fatalf("open Redis client: %v", err)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("close Redis client: %v", err)
		}
	}()

	var messagingClient *messaging.Client
	if configuration.Firebase.ProjectID != "" {
		_, messagingClient, err = firebaseClient.Open(
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
	outboundHTTP := httpIntegration.New(nil)
	users := userService.GetInstance(userRepository.GetInstance(sqlDatabase), cacheStore)
	volunteerServices := volunteerService.GetInstance(
		volunteerServiceRepository.GetInstance(sqlDatabase),
		cacheStore,
	)
	messages := messageService.GetInstance(
		messageRepository.GetInstance(sqlDatabase),
		cacheStore,
	)
	meals := mealService.GetInstance(
		mealRepository.GetInstance(sqlDatabase),
		cacheStore,
	)
	notifications := notificationService.GetInstance(firebaseClient.NewSender(messagingClient))
	transactions := transactionService.GetInstance(
		transactionRepository.GetInstance(sqlDatabase),
		emailMessaging.NewTransactionMailer(configuration.Email.AppPassword),
	)
	payments := paymentService.GetInstance(
		stripeIntegration.NewClient(
			configuration.Stripe.SecretKey,
			configuration.Stripe.Version,
		),
		configuration.Stripe.PublishableKey,
	)
	events := eventService.GetInstance(
		planningcenter.NewClient(
			outboundHTTP,
			cacheStore,
			configuration.PlanningCenter.URL,
			configuration.PlanningCenter.TokenURL,
		),
		cacheStore,
	)
	sermons := sermonService.GetInstance(
		sermoncloud.NewWithTransport(outboundHTTP, configuration.SermonCloud.URL),
		cacheStore,
	)
	httpHandlers := httpHandler.GetInstance(
		userHTTP.GetInstance(users),
		volunteerServiceHTTP.GetInstance(volunteerServices),
		mealHTTP.GetInstance(meals),
		messageHTTP.GetInstance(messages),
		notificationHTTP.GetInstance(notifications),
		eventHTTP.GetInstance(events),
		sermonHTTP.GetInstance(sermons),
		paymentHTTP.GetInstance(payments),
		transactionHTTP.GetInstance(transactions),
	)
	app := fiber.New(fiber.Config{ErrorHandler: httpHandler.ErrorHandler})
	httpMiddleware.ConfigureAppMiddleware(app)

	api := app.Group("/api", func(c fiber.Ctx) error {
		c.Response().Header.Set("Cache-Control", "no-cache")
		return c.Next()
	})

	api.Get("/metrics", httpMiddleware.Monitor())
	v1 := api.Group("/v1", httpMiddleware.BearerTokenAuthentication(configuration.Security.APIKey))

	v1.Post("/meals", httpHandlers.Meal.Create)
	v1.Get("/meals", httpHandlers.Meal.List)
	v1.Put("/meals", httpHandlers.Meal.Update)
	v1.Delete("/meals/:mealId", httpHandlers.Meal.Delete)
	v1.Post("/meals/meal", httpHandlers.Meal.AssignVolunteer)
	v1.Put("/meals/meal", httpHandlers.Meal.UpdateVolunteeredMeal)

	v1.Post("/users", httpHandlers.User.Create)
	v1.Get("/users/:userId", httpHandlers.User.GetByID)
	v1.Put("/users", httpHandlers.User.Update)
	v1.Delete("/users/:userId", httpHandlers.User.Delete)

	v1.Post("/payments", httpHandlers.Payment.Create)

	v1.Get("/transactions", httpHandlers.Transaction.List)
	v1.Get("/transactions/:userId", httpHandlers.Transaction.ListByUser)
	v1.Delete("/transactions/:transactionId", httpHandlers.Transaction.Delete)
	v1.Post("/transactions", httpHandlers.Transaction.Create)

	v1.Get("/sermons", httpHandlers.Sermon.List)

	v1.Get("/services", httpHandlers.VolunteerService.ListServices)
	v1.Post("/services", httpHandlers.VolunteerService.Create)
	v1.Get("/services/:type", httpHandlers.VolunteerService.GetByType)
	v1.Put("/services/day", httpHandlers.VolunteerService.UpdateDay)

	v1.Post("/messages", httpHandlers.Message.Create)
	v1.Put("/messages", httpHandlers.Message.Update)
	v1.Get("/messages/admin", httpHandlers.Message.ListAdmin)
	v1.Get("/messages/unpublished", httpHandlers.Message.ListUnpublished)
	v1.Get("/messages/:userId", httpHandlers.Message.ListPublishedByUser)
	v1.Post("/messages/notes", httpHandlers.Message.CreateNote)
	v1.Put("/messages/notes", httpHandlers.Message.UpdateNote)
	v1.Delete("/messages/notes/:noteId", httpHandlers.Message.DeleteNote)
	v1.Delete("/messages/:messageId", httpHandlers.Message.Delete)

	v1.Post("/notifications", httpHandlers.Notification.Send)

	v1.Get("/events", httpHandlers.Event.List)

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/swagger/index.html")
	})
	app.Get("/swagger/*", swaggo.New(swaggo.Config{
		URL:          "/swagger/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	log.Fatal(app.Listen(":" + configuration.Port))
}
