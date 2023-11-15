package controller

import (
	"Momentum/constants"
	"Momentum/prisma/db"
	b64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type Controller struct {
	HttpClient                 *http.Client
	PrismaClient               *db.PrismaClient
	Context                    context.Context
	Redis                      *redis.Client
	Firebase                   *firebase.App
	Messaging                  *messaging.Client
	CanSendNotificationChannel chan bool
}

func GetControllerInstance() *Controller {
	return &Controller{
		HttpClient:                 &http.Client{Timeout: time.Second * 10},
		CanSendNotificationChannel: make(chan bool),
	}
}

func (controller *Controller) SetPrismaClient(client *db.PrismaClient) {
	controller.PrismaClient = client
}

func (controller *Controller) SetRedisClient(client *redis.Client) {
	controller.Redis = client
}

func (controller *Controller) InitPrismaClient() {
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}
	controller.SetPrismaClient(prismaClient)
}
func (controller *Controller) InitRedisClient() {
	opt, err := redis.ParseURL(constants.RedisURL)
	if err != nil {
		log.Fatalln(err)
	}
	redisClient := redis.NewClient(opt)
	controller.SetRedisClient(redisClient)
	log.Println("Connected to Redis")
}

func (controller *Controller) InitFirebaseApp() {
	if controller.Redis == nil {
		controller.InitRedisClient()
	}
	firebaseToken, err := controller.Redis.Get(controller.Context, constants.FirebaseServiceTokenKey).Result()
	if err != nil {
		fmt.Println(err)
	}
	creds := option.WithCredentialsJSON([]byte(firebaseToken))
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	firebaseApp, err := firebase.NewApp(controller.Context, config, creds)
	if err != nil {
		fmt.Println("[Firebase]", err)
	}
	controller.setFirebaseApp(firebaseApp)
}

func (controller *Controller) setFirebaseApp(app *firebase.App) {
	controller.Firebase = app
	messagingClient, err := app.Messaging(controller.Context)
	if err != nil {
		fmt.Println(err)
	}
	controller.Messaging = messagingClient
	log.Println("Connected to Firebase")
}

func (controller *Controller) SetContext(context context.Context) {
	controller.Context = context
}

type StripeResponse struct {
	Endpoint string
	ReadData func() ([]byte, error)
}

func (controller *Controller) StripeRequest(waitGroup *sync.WaitGroup, channel chan<- StripeResponse, endpoint string) {
	defer waitGroup.Done()
	req, err := http.NewRequest("POST", "https://api.stripe.com/v1"+endpoint, nil)
	if err != nil {
		log.Fatalln("New Request", err)
	}

	req.Header.Add("Authorization", "Basic "+b64.URLEncoding.EncodeToString([]byte(os.Getenv("STRIPE_SECRET_KEY"))))

	switch getBasePath(endpoint) {
	case constants.StripeEphemeralKeyRoute:
		req.Header.Add("Stripe-Version", "2020-08-27")
	default:
		break
	}

	resp, err := controller.HttpClient.Do(req)
	if err != nil {
		log.Panic("Request", err)
	}

	channel <- StripeResponse{
		Endpoint: getBasePath(endpoint),
		ReadData: func() ([]byte, error) {
			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		},
	}
}

func getBasePath(s string) string {
	return strings.SplitAfter(s, "?")[0]
}
