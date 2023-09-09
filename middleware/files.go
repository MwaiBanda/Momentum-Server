package middleware

import (
	"Momentum/constants"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func configureFiles(app *fiber.App) {
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:   rice.MustFindBox("../cms/dist/assets").HTTPBox(),
		Browse: true,
	}))

	for _, route := range []string{
		constants.DashboardServicesRoute,
		constants.DashboardHomeRoute,
		constants.DashboardRoute,
	} {
		app.Use(route, filesystem.New(filesystem.Config{
			Root:   rice.MustFindBox("../cms/dist").HTTPBox(),
			Browse: true,
		}))
	}
}
