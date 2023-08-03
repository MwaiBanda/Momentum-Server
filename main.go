package main

import (
	"Momentum/prisma/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
	desc      string `json:"desc"`
}

func main() {
	app := fiber.New()
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	app.Get("/", func(context *fiber.Ctx) error {
		demo := Post{Title: "Hello World"}
		_, err := json.Marshal(demo)
		if err != nil {
			log.Fatal(err.Error())
		}
		return context.JSON(demo)
	})

	app.Get("/posts", func(context *fiber.Ctx) error {
		var posts []Post
		post, err := client.Post.FindMany().Exec(ctx)
		if err != nil {
			fmt.Println(err)
		}

		result, _ := json.MarshalIndent(post, "", "  ")
		if err := json.Unmarshal(result, &posts); err != nil {
			fmt.Println(err)
		}

		return context.JSON(posts)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
