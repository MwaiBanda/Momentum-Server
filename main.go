package main

import (
	"Momentum/constants"
	"Momentum/controller"
	_ "Momentum/docs"
	_ "Momentum/httputil"
	"Momentum/prisma/db"
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/swagger"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
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
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	controllerInstance := controller.GetControllerInstance()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, key string) (bool, error) {
			hashedAPIKey := sha256.Sum256([]byte(os.Getenv("API_KEY")))
			hashedKey := sha256.Sum256([]byte(key))

			if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:   rice.MustFindBox("./cms/dist/assets").HTTPBox(),
		Browse: true,
	}))

	for _, route := range []string {
		constants.DashboardServicesRoute,
		constants.DashboardHomeRoute,
		constants.DashboardRoute,
	} {
		app.Use(route, filesystem.New(filesystem.Config{
			Root:   rice.MustFindBox("./cms/dist").HTTPBox(),
			Browse: true,
	    }))
	}

	api := app.Group("/api")
	v1 := api.Group("/v1", authMiddleware)

	v1.Post("/meals", controllerInstance.PostMeal)
	v1.Get("/meals", controllerInstance.GetAllMeals)
	v1.Post("/meals/meal", controllerInstance.PostVolunteeredMealForMeal)

	v1.Post("/users", controllerInstance.PostUser)
	v1.Get("/users/:userId", controllerInstance.GetUserById)
	v1.Put("/users", controllerInstance.UpdateUser)
	v1.Delete("/users/:userId", controllerInstance.DeleteUserById)

	v1.Post("/payments", controllerInstance.PostPayment)

	v1.Get("/transactions/:userId", controllerInstance.GetTransactionsByUserId)
	v1.Delete("/transactions/:transactionId", controllerInstance.DeleteTransactionsById)

	v1.Get("/sermons", controllerInstance.GetAllSermons)

	app.Get("/*", swagger.HandlerDefault)

	app.Get("/*", swagger.New(swagger.Config{
		URL:          "https://services.momentumchurch.dev/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	backgroundContext := context.Background()
	controllerInstance.SetPrismaClient(client)
	controllerInstance.SetRedisClient(redis.NewClient(opt))
	controllerInstance.SetContext(backgroundContext)

	log.Fatal(app.Listen(":" + port))
}
