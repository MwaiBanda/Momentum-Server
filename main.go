package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Hello struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()
	res, err := json.Marshal(Hello{Text: "Hello World"})
	if err != nil {
		log.Fatal(err.Error())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(res)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
