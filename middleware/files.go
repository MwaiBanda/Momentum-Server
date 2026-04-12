package middleware

import (
	"Momentum/constants"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func configureFiles(app *fiber.App) {
	app.Use("/assets", static.New("cms/dist/assets", static.Config{
		Browse: true,
	}))

	for _, route := range []string{
		constants.DashboardRoute,
		constants.DashboardMessageRoute,
		constants.DashboardNotificationsRoute,
		constants.DashboardPaymentsRoute,
		constants.DashboardUsersRoute,
	} {
		app.Use(route, func(c fiber.Ctx) error {
			c.Response().Header.Set("Cache-Control", "no-store, no-cache, must-revalidate")
			c.Response().Header.Set("Pragma", "no-cache")
			c.Response().Header.Set("Expires", "0")
			return c.Next()
		})

		app.Use(route, static.New("cms/dist", static.Config{
			Browse: true,
		}))
	}
}
