package main

import (
	handlers "Momentum/controller"
	_ "Momentum/docs"
	"Momentum/middleware"
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
	log.Println("Starting Momentum API")
	controller := handlers.GetControllerInstance()
	app := fiber.New()
	port := func() string {
		if len(os.Getenv("PORT")) > 0 {
			return os.Getenv("PORT")
		} else {
			return "8085"
		}
	}()

	middleware.ConfigureAppMiddleWare(app)

	api := app.Group("/api", func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache")
		return c.Next()
	})

	api.Get("/metrics", middleware.Monitor())
	v1 := api.Group("/v1", middleware.BearerTokenAuthentication())

	v1.Post("/meals", controller.PostMeal)
	v1.Get("/meals", controller.GetAllMeals)
	v1.Delete("/meals/:mealId", controller.DeleteMealById)
	v1.Post("/meals/meal", controller.PostVolunteeredMealForMeal)
	v1.Put("/meals/meal", controller.UpdateVolunteeredMeal)

	v1.Post("/users", controller.PostUser)
	v1.Get("/users/:userId", controller.GetUserById)
	v1.Put("/users", controller.UpdateUser)
	v1.Delete("/users/:userId", controller.DeleteUserById)

	v1.Post("/payments", controller.PostPayment)

	v1.Get("/transactions", controller.GetAllTransactions)
	v1.Get("/transactions/:userId", controller.GetTransactionsByUserId)
	v1.Delete("/transactions/:transactionId", controller.DeleteTransactionsById)
	v1.Post("/transactions", controller.PostTransaction)

	v1.Get("/sermons", controller.GetAllSermons)
	
	v1.Get("/services", controller.GetAllServices)
	v1.Post("/services", controller.PostVolunteerService)
	v1.Get("/services/:type", controller.GetServiceByType)
	v1.Put("/services/day", controller.UpdateVolunteeredDay)

	v1.Post("/messages", controller.PostMessage)
	v1.Put("/messages", controller.UpdateMessage)
	v1.Get("/messages/unpublished", controller.GetUnpublishedMessages)
	v1.Get("/messages/:userId", controller.GetAllMessages)
	v1.Post("/messages/notes", controller.AddUserNoteToMessage)
	v1.Put("/messages/notes", controller.UpdateUserNote)
	v1.Delete("/messages/notes/:noteId", controller.DeleteNote)
	v1.Delete("/messages/:messageId", controller.DeleteMessage)

	v1.Post("/notifications", controller.PostNotification)

	v1.Get("/events", controller.GetAllEvents)

	app.Get("/*", swagger.HandlerDefault)
	app.Get("/*", swagger.New(swagger.Config{
		URL:          "https://services.momentumchurch.dev/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	defer controller.DisconectPrismaClient()

	log.Fatal(app.Listen(":" + port))
}
