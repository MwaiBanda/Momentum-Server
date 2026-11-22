package application

import (
	httpHandler "Momentum/infrastructure/http"
	httpMiddleware "Momentum/infrastructure/http/middleware"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func (app *Application) registerRoutes(handlers *httpHandler.Handler) {
	api := app.server.Group("/api", func(c fiber.Ctx) error {
		c.Response().Header.Set("Cache-Control", "no-cache")
		return c.Next()
	})

	v1 := api.Group("/v1", httpMiddleware.BearerTokenAuthentication(app.config.Security.APIKey))

	v1.Post("/meals", handlers.Meal.Create)
	v1.Get("/meals", handlers.Meal.List)
	v1.Put("/meals", handlers.Meal.Update)
	v1.Delete("/meals/:mealId", handlers.Meal.Delete)
	v1.Post("/meals/meal", handlers.Meal.AssignVolunteer)
	v1.Put("/meals/meal", handlers.Meal.UpdateVolunteeredMeal)

	v1.Post("/users", handlers.User.Create)
	v1.Get("/users/:userId", handlers.User.GetByID)
	v1.Put("/users", handlers.User.Update)
	v1.Delete("/users/:userId", handlers.User.Delete)

	v1.Post("/payments", handlers.Payment.Create)

	v1.Get("/transactions", handlers.Transaction.List)
	v1.Get("/transactions/:userId", handlers.Transaction.ListByUser)
	v1.Delete("/transactions/:transactionId", handlers.Transaction.Delete)
	v1.Post("/transactions", handlers.Transaction.Create)

	v1.Get("/sermons", handlers.Sermon.List)

	v1.Get("/services", handlers.VolunteerService.ListServices)
	v1.Post("/services", handlers.VolunteerService.Create)
	v1.Get("/services/:type", handlers.VolunteerService.GetByType)
	v1.Put("/services/day", handlers.VolunteerService.UpdateDay)

	v1.Post("/messages", handlers.Message.Create)
	v1.Put("/messages", handlers.Message.Update)
	v1.Get("/messages/admin", handlers.Message.ListAdmin)
	v1.Get("/messages/unpublished", handlers.Message.ListUnpublished)
	v1.Get("/messages/:userId", handlers.Message.ListPublishedByUser)
	v1.Post("/messages/notes", handlers.Message.CreateNote)
	v1.Put("/messages/notes", handlers.Message.UpdateNote)
	v1.Delete("/messages/notes/:noteId", handlers.Message.DeleteNote)
	v1.Delete("/messages/:messageId", handlers.Message.Delete)

	v1.Post("/notifications", handlers.Notification.Send)

	v1.Get("/events", handlers.Event.List)

	app.server.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/swagger/index.html")
	})
	app.server.Get("/swagger/*", swaggo.New(swaggo.Config{
		URL:          "/swagger/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))
}
