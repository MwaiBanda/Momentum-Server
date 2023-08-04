package main

import (
	"Momentum/controller"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"log"
	"os"

	_ "Momentum/docs"
	_ "Momentum/httputil"
	"Momentum/prisma/db"
	"github.com/gofiber/fiber/v2"
)

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
		port = "8080"
	}

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/post", controllerInstance.GetPost)
	v1.Get("/posts", controllerInstance.GetPosts)
	v1.Post("/payment", controllerInstance.PostPayment)
	app.Get("/*", swagger.HandlerDefault)

	app.Get("/*", swagger.New(swagger.Config{
		URL:          "http://services.momentumchurch.dev/doc.json",
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

	backgroundContext := context.Background()
	controllerInstance.SetClient(client)
	controllerInstance.SetContext(backgroundContext)

	log.Fatal(app.Listen(":" + port))
}
