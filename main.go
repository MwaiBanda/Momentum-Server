package main

import (
	handlers "Momentum/controller"
	_ "Momentum/docs"
	"Momentum/middleware"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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
	controller := handlers.GetControllerInstance()
	controller.SetContext(context.Background())
	go func ()  {
		controller.InitRedisClient()
		go controller.InitPrismaClient()
		go controller.InitFirebaseApp()
	}()
	app := fiber.New()
	port := func() string {
		if len(os.Getenv("PORT")) > 0 {
			return os.Getenv("PORT")
		} else {
			return "8083"
		}
	}()

	middleware.ConfigureAppMiddleWare(app)

	api := app.Group("/api")
	api.Get("/metrics", middleware.Monitor())
	v1 := api.Group("/v1", middleware.BearerTokenAuthentication())

	v1.Post("/meals", controller.PostMeal)
	v1.Get("/meals", controller.GetAllMeals)
	v1.Post("/meals/meal", controller.PostVolunteeredMealForMeal)

	v1.Post("/users", controller.PostUser)
	v1.Get("/users/:userId", controller.GetUserById)
	v1.Put("/users", controller.UpdateUser)
	v1.Delete("/users/:userId", controller.DeleteUserById)

	v1.Post("/payments", controller.PostPayment)

	v1.Get("/transactions/:userId", controller.GetTransactionsByUserId)
	v1.Delete("/transactions/:transactionId", controller.DeleteTransactionsById)

	v1.Get("/sermons", controller.GetAllSermons)

	v1.Get("/messages/:userId", controller.GetAllMessages)

	v1.Post("/notifications", controller.PostNotification)

	app.Get("/*", swagger.HandlerDefault)
	app.Get("/*", swagger.New(swagger.Config{
		URL:          "https://services.momentumchurch.dev/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	defer func() {
		if err := controller.PrismaClient.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	log.Fatal(app.Listen(":" + port))
}
