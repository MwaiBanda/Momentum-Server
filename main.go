package main

import (
	"Momentum/controller"
	_ "Momentum/docs"
	_ "Momentum/httputil"
	"Momentum/prisma/db"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/meals", controllerInstance.PostMeal)
	v1.Get("/meals", controllerInstance.GetMeals)
	v1.Post("/meals/participant", controllerInstance.PostMealParticipant)
	v1.Post("/meals/meal", controllerInstance.PostMealForMeal)

	v1.Post("/users", controllerInstance.PostUser)
	v1.Get("/users/:userId", controllerInstance.GetUserById)

	v1.Post("/payments", controllerInstance.PostPayment)

	v1.Get("/sermons", controllerInstance.GetSermon)

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
