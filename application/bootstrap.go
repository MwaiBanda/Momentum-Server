package application

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
	httpHandler "Momentum/infrastructure/http"
	eventHTTP "Momentum/infrastructure/http/event"
	mealHTTP "Momentum/infrastructure/http/meal"
	messageHTTP "Momentum/infrastructure/http/message"
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
	"Momentum/infrastructure/observability"
	database "Momentum/infrastructure/persistence/db"
	redisCache "Momentum/infrastructure/persistence/redis"
	mealRepository "Momentum/infrastructure/persistence/repository/meal"
	messageRepository "Momentum/infrastructure/persistence/repository/message"
	transactionRepository "Momentum/infrastructure/persistence/repository/transaction"
	userRepository "Momentum/infrastructure/persistence/repository/user"
	volunteerServiceRepository "Momentum/infrastructure/persistence/repository/volunteer_service"

	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	"github.com/redis/go-redis/extra/redisotel/v9"
)

func (app *Application) initialize(ctx context.Context) error {
	configuration, err := appconfig.LoadConfig()
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}
	app.config = configuration

	telemetryProviders, err := observability.Setup(
		ctx,
		configuration.Telemetry.ServiceName,
		configuration.Environment,
		configuration.Telemetry.OTLPEndpoint,
	)
	if err != nil {
		return fmt.Errorf("initialize OpenTelemetry: %w", err)
	}
	app.shutdownTelemetry = telemetryProviders.Shutdown

	app.database, err = database.Open(ctx, configuration.Database.URL)
	if err != nil {
		return fmt.Errorf("open SQL database: %w", err)
	}
	if err := database.Migrate(ctx, app.database); err != nil {
		return fmt.Errorf("migrate SQL database: %w", err)
	}

	app.redisClient, err = redisCache.Open(configuration.Redis.URL)
	if err != nil {
		return fmt.Errorf("open Redis client: %w", err)
	}
	if err := redisotel.InstrumentTracing(app.redisClient); err != nil {
		return fmt.Errorf("instrument Redis tracing: %w", err)
	}
	if err := redisotel.InstrumentMetrics(app.redisClient); err != nil {
		return fmt.Errorf("instrument Redis metrics: %w", err)
	}

	var messagingClient *messaging.Client
	if configuration.Firebase.ProjectID != "" {
		_, messagingClient, err = firebaseClient.Open(
			ctx,
			app.redisClient,
			configuration.Firebase.ProjectID,
		)
		if err != nil {
			log.Printf("initialize Firebase: %v", err)
		}
	} else {
		log.Println("Firebase disabled: FIREBASE_PROJECT_ID is not configured")
	}

	cacheStore := redisCache.NewCache(app.redisClient)
	outboundHTTP := httpIntegration.New(nil)
	users := userService.GetInstance(userRepository.GetInstance(app.database), cacheStore)
	volunteerServices := volunteerService.GetInstance(
		volunteerServiceRepository.GetInstance(app.database),
		cacheStore,
	)
	messages := messageService.GetInstance(
		messageRepository.GetInstance(app.database),
		cacheStore,
	)
	meals := mealService.GetInstance(
		mealRepository.GetInstance(app.database),
		cacheStore,
	)
	notifications := notificationService.GetInstance(firebaseClient.NewSender(messagingClient))
	transactions := transactionService.GetInstance(
		transactionRepository.GetInstance(app.database),
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
	handlers := httpHandler.GetInstance(
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

	app.configureHTTP(handlers)
	return nil
}
