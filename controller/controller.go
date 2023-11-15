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
)

type Controller struct {
	HttpClient                 *http.Client
	Prisma                     *db.PrismaClient
	Context                    context.Context
	Redis                      *redis.Client
	Firebase                   *firebase.App
	Messaging                  *messaging.Client
	CanSendNotificationChannel chan bool
}

func GetControllerInstance() *Controller {
	return &Controller{
		HttpClient: &http.Client{Timeout: time.Second * 10},
	}
}

func (controller *Controller) SetPrismaClient(client *db.PrismaClient) {
	controller.Prisma = client
}

func (controller *Controller) SetRedisClient(client *redis.Client) {
	controller.Redis = client
}

func (controller *Controller) SetFirebaseApp(app *firebase.App) {
	controller.CanSendNotificationChannel = make(chan bool)
	controller.Firebase = app
	messagingClient, err := app.Messaging(controller.Context)
	if err != nil {
		fmt.Println(err)
	}
	controller.Messaging = messagingClient
	controller.CanSendNotificationChannel <- true
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
