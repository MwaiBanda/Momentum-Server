package middleware

import (
	"Momentum/constants"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func configureFiles(app *fiber.App) {
	app.Use("/assets", static.New("../cms/dist/assets", static.Config{
		Browse: true,
	}))

	for _, route := range []string{
		constants.DashboardRoute,
		constants.DashboardMessageRoute,
		constants.DashboardNotificationsRoute,
		constants.DashboardPaymentsRoute,
		constants.DashboardUsersRoute,
	} {
		app.Use(route, static.New("../cms/dist", static.Config{
			Browse: true,
		}))
	}
}
