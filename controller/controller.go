package controller

import (
	appconfig "Momentum/application/config"
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
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type Controller struct {
	Configuration appconfig.Config
	HttpClient    *http.Client
	PrismaClient  *db.PrismaClient
	Context       context.Context
	Redis         *redis.Client
}

type Dependencies struct {
	Configuration appconfig.Config
	HttpClient    *http.Client
	PrismaClient  *db.PrismaClient
	Context       context.Context
	Redis         *redis.Client
}

func NewController(dependencies Dependencies) *Controller {
	httpClient := dependencies.HttpClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Second * 10}
	}

	controllerContext := dependencies.Context
	if controllerContext == nil {
		controllerContext = context.Background()
	}

	return &Controller{
		Configuration: dependencies.Configuration,
		HttpClient:    httpClient,
		PrismaClient:  dependencies.PrismaClient,
		Context:       controllerContext,
		Redis:         dependencies.Redis,
	}
}

type StripeResponse struct {
	Endpoint string
	ReadData func() ([]byte, error)
}

func (controller *Controller) GetPlanningCenterToken() Token {
	var tokenResponse TokenResponse
	var cachedToken, err = controller.Redis.Get(controller.Context, constants.TokenKey).Result()
	if err == redis.Nil {
		req, reqErr := http.NewRequest(http.MethodPost, controller.Configuration.PlanningCenter.TokenURL, nil)
		if reqErr != nil {
			log.Panic("GetPlanningCenterToken", reqErr)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, reqErr := controller.HttpClient.Do(req)
		if reqErr != nil {
			log.Panic("GetPlanningCenterToken", reqErr)
		}
		defer resp.Body.Close()
		body, reqErr := io.ReadAll(resp.Body)
		log.Println("GetPlanningCenterToken[REQUEST]", resp.StatusCode, string(body), reqErr)
		if reqErr != nil {
			log.Panic("GetPlanningCenterToken", reqErr)
		}
		if err := json.Unmarshal(body, &tokenResponse); err != nil {
			log.Println("GetPlanningCenterToken", err)
		}
		if err := controller.Redis.Set(controller.Context, constants.TokenKey, string(body), 1*time.Hour+30*time.Minute).Err(); err != nil {
			log.Println("GetPlanningCenterToken[Redis]", err)
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

	req.Header.Add("Authorization", "Basic "+b64.URLEncoding.EncodeToString([]byte(controller.Configuration.Stripe.SecretKey)))

	switch getBasePath(endpoint) {
	case constants.StripeEphemeralKeyRoute:
		req.Header.Add("Stripe-Version", controller.Configuration.Stripe.Version)
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
	url := controller.Configuration.PlanningCenter.URL + "?per_page=100&where[ends_at][gte]=" + time.Now().Add(-4*time.Hour).Format(time.RFC3339)
	token := controller.GetPlanningCenterToken()
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return EventReponse{ReadData: func() ([]byte, []error) { return nil, []error{reqErr} }}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token.Type+" "+token.Attributes.Token)
	resp, reqErr := controller.HttpClient.Do(req)
	var body []byte
	if reqErr == nil {
		defer resp.Body.Close()
		body, reqErr = io.ReadAll(resp.Body)
	}
	statusCode := 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	return EventReponse{
		ReadData: func() ([]byte, []error) {
			log.Println("EventRequest[REQUEST]", statusCode, string(body), reqErr)
			if reqErr != nil {
				return body, []error{reqErr}
			}
			return body, nil
		},
	}
}

type EventReponse struct {
	ReadData func() ([]byte, []error)
}

func (controller *Controller) PostTransactionMail(mail model.TransactionMailRequest) {
	from := "bandamwai@gmail.com"
	password := controller.Configuration.Email.AppPassword

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
