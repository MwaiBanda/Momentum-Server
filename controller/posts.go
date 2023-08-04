package controller

import (
	_ "Momentum/httputil"
	"Momentum/model"

	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

// GetPost godoc
//
//	@Summary		Show a post
//	@Description	get a hello post
//	@tags			Posts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Post
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/post [get]
func (controller *Controller) GetPost(context *fiber.Ctx) error {
	demo := model.Post{Title: "GetPost World"}
	_, err := json.Marshal(demo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return context.JSON(demo)
}

// GetPosts godoc
//
//	@Summary		Show a list of posts
//	@Description	get a list of posts
//	@Accept			json
//	@Produce		json
//	@tags			Posts
//	@Success		200	{array}		model.Post
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/posts [get]
func (controller *Controller) GetPosts(fiberContext *fiber.Ctx) error {
	var posts []model.Post
	post, err := controller.prismaClient.Post.FindMany().Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(post, "", "  ")
	if err := json.Unmarshal(result, &posts); err != nil {
		fmt.Println(err)
	}

	return fiberContext.JSON(posts)
}
