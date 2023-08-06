package controller

import (
	"Momentum/prisma/db"
	b64 "encoding/base64"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Controller struct {
	httpClient *http.Client
	prisma     *db.PrismaClient
	context    context.Context
	redis      *redis.Client
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
func (controller *Controller) StripeRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("POST", "https://api.stripe.com/v1"+endpoint, nil)
	if err != nil {
		log.Fatalln("New Request", err)
	}

	req.Header.Add("Authorization", "Basic "+b64.URLEncoding.EncodeToString([]byte(os.Getenv("STRIPE_SECRET_KEY"))))
	switch strings.Split(endpoint, "?")[0] {
	case "/ephemeral_keys":
		req.Header.Add("Stripe-Version", "2020-08-27")
	default:
		break
	}
	resp, err := controller.httpClient.Do(req)
	if err != nil {
		log.Fatalln("Request", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (controller *Controller) SetContext(context context.Context) {
	controller.context = context
}
