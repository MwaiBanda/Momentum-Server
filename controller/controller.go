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

	"firebase.google.com/go/v4/messaging"
	firebase "firebase.google.com/go/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type Controller struct {
	httpClient *http.Client
	prisma     *db.PrismaClient
	context    context.Context
	redis      *redis.Client
	firebase   *firebase.App
	messaging *messaging.Client
}

func GetControllerInstance() *Controller {
	return &Controller{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
}

func (controller *Controller) SetPrismaClient(client *db.PrismaClient) {
	controller.prisma = client
}

func (controller *Controller) SetRedisClient(client *redis.Client) {
	controller.redis = client
}

func (controller *Controller) SetFirebaseApp(app *firebase.App) {
	controller.firebase = app
	messagingClient, err := app.Messaging(controller.context)
	if err != nil {
		fmt.Println(err)
	}
	controller.messaging = messagingClient

}

func (controller *Controller) SetContext(context context.Context) {
	controller.context = context
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

	resp, err := controller.httpClient.Do(req)
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
