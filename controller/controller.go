package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"Momentum/prisma/db"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/joho/godotenv"
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
	controller := &Controller{
		HttpClient:                 &http.Client{Timeout: time.Second * 10},
		CanSendNotificationChannel: make(chan bool),
	}
	err := godotenv.Load()
	if os.Getenv("ENVIRONMENT") == constants.Development {
		if err != nil {
			log.Fatalln("Error loading .env file")
		}
	}

	controller.InitDependencies()
	return controller
}

func (controller *Controller) InitDependencies() {
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)
	go controller.SetContext(context.Background(), waitGroup)
	waitGroup.Add(1)
	go controller.InitRedisClient(waitGroup)
	waitGroup.Wait()
	waitGroup.Add(1)
	go controller.InitPrismaClient(waitGroup)
	waitGroup.Add(1)
	go controller.InitFirebaseApp(waitGroup)
	waitGroup.Wait()
}

func (controller *Controller) SetPrismaClient(client *db.PrismaClient) {
	controller.PrismaClient = client
}

func (controller *Controller) SetRedisClient(client *redis.Client) {
	controller.Redis = client
}

func (controller *Controller) InitPrismaClient(wg ...*sync.WaitGroup) {
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}
	if len(wg) > 0 {
		defer wg[0].Done()
	}
	controller.SetPrismaClient(prismaClient)
	log.Println("Connected to Prisma")
}

func (controller *Controller) DisconectPrismaClient() {
	if err := controller.PrismaClient.Prisma.Disconnect(); err != nil {
		panic(err)
	}
}

func (controller *Controller) InitRedisClient(wg ...*sync.WaitGroup) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	redisClient := redis.NewClient(opt)
	if len(wg) > 0 {
		defer wg[0].Done()
	}
	controller.SetRedisClient(redisClient)
	log.Println("Connected to Redis")
}

func (controller *Controller) InitFirebaseApp(wg ...*sync.WaitGroup) {
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
	if len(wg) > 0 {
		defer wg[0].Done()
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

func (controller *Controller) SetContext(context context.Context, wg *sync.WaitGroup) {
	controller.Context = context
	wg.Done()
}

type StripeResponse struct {
	Endpoint string
	ReadData func() ([]byte, error)
}

func (controller *Controller) GetPlanningCenterToken() Token {
	var tokenResponse TokenResponse
	var cachedToken, err = controller.Redis.Get(controller.Context, constants.TokenKey).Result()
	if err == redis.Nil {
		resp, err := controller.HttpClient.Post(os.Getenv("PLANNING_CENTER_TOKEN_URL"), "application/json", nil)
		if err != nil {
			log.Println("GetPlanningCenterToken", err)
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetPlanningCenterToken", err)
		}
		if err := json.Unmarshal(data, &tokenResponse); err != nil {
			log.Println("GetPlanningCenterToken", err)
		}
		if err = controller.Redis.Set(controller.Context, constants.TokenKey, string(data), 1*time.Hour+30*time.Minute).Err(); err != nil {
			log.Println("[GetPlanningCenterToken]", err.Error())
		}
	} else if err != nil {
		log.Println("GetPlanningCenterToken", err)
	} else {
		err := json.Unmarshal([]byte(cachedToken), &tokenResponse)
		if err != nil {
			log.Println("GetPlanningCenterToken", err)
		}
	}

	return tokenResponse.Data
}

type TokenResponse struct {
	Data Token `json:"data"`
}
type Token struct {
	Type       string          `json:"type"`
	Attributes TokenAttributes `json:"attributes"`
}
type TokenAttributes struct {
	Token string `json:"token"`
}

func (controller *Controller) StripeRequest(waitGroup *sync.WaitGroup, channel chan<- StripeResponse, endpoint string) {
	defer waitGroup.Done()
	req, err := http.NewRequest("POST", "https://api.stripe.com/v1"+endpoint, nil)
	if err != nil {
		log.Fatalln("StripeRequest", err)
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
		log.Panic("StripeRequest", err)
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

func (controller *Controller) EventRequest() EventReponse {
	url := os.Getenv("PLANNING_CENTER_URL") + "?per_page=100&where[ends_at][gte]=" + time.Now().Add(-4*time.Hour).Format(time.RFC3339)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("New Request", err)
	}

	token := controller.GetPlanningCenterToken()
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", token.Type+" "+token.Attributes.Token)

	resp, err := controller.HttpClient.Do(req)
	if err != nil {
		log.Panic("Request", err)
	}

	return EventReponse{
		ReadData: func() ([]byte, error) {
			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		},
	}
}

type EventReponse struct {
	ReadData func() ([]byte, error)
}

func (controller *Controller) PostTransactionMail(mail model.TransactionMailRequest) {
	from := "bandamwai@gmail.com"
	password := os.Getenv("EMAIL_APP_PASSWORD")

	to := []string{
		"bandamwai@gmail.com",
		"charlie@momentumindiana.org",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("./controller/transaction.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: New Payment 📲 \n%s\n\n", mimeHeaders)))

	t.Execute(&body, mail)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")

}

func (controller *Controller) AddUserCacheSession(key string) {
	v, err := controller.Redis.Get(controller.Context, constants.UserCachedSessionKey).Result()
	if err == redis.Nil {
		str, err := json.Marshal([]string{key})
		if err != nil {
			fmt.Println(err)
		}
		if err := controller.Redis.Set(controller.Context, constants.UserCachedSessionKey, string(str), time.Hour*24).Err(); err != nil {
			fmt.Println(err)
		}
	} else if err != nil {
		fmt.Println(err)
	} else {
		var sessions []string
		if err := json.Unmarshal([]byte(v), &sessions); err != nil {
			fmt.Println(err)
		}
		sessions = append(sessions, key)
		str, _ := json.Marshal(sessions)
		if err := controller.Redis.Set(controller.Context, constants.UserCachedSessionKey, string(str), time.Hour*24).Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func (controller *Controller) DeleteUserCachedSessions() {
	v, err := controller.Redis.Get(controller.Context, constants.UserCachedSessionKey).Result()
	if err == redis.Nil {
		fmt.Println("No sessions")
	} else if err != nil {
		fmt.Println(err)
	} else {
		var sessions []string
		if err := json.Unmarshal([]byte(v), &sessions); err != nil {
			fmt.Println(err)
		}
		sessions = append(sessions, constants.UserCachedSessionKey)
		if err := controller.Redis.Del(controller.Context, sessions...).Err(); err != nil {
			fmt.Println(err)
		}
	}
}
