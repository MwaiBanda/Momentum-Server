package sermon

import (
	"errors"
	"log"

	applicationsermon "Momentum/application/service/sermon"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	sermons applicationsermon.UseCase
}

func GetInstance(sermons applicationsermon.UseCase) *Handler {
	return &Handler{sermons: sermons}
}

// List godoc
//
//	@Summary		Show a list of sermons
//	@Description	Get a paginated list of sermons
//	@Accept			json
//	@Produce		json
//	@tags			Sermons
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			page			query	integer	false	"Provide a page number"	default(1)
//	@Success		200	{object}	sermon.Response
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/sermons [get]
func (handler *Handler) List(ctx fiber.Ctx) error {
	if handler.sermons == nil {
		log.Printf("[ListSermons] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.sermons.List(ctx.Context(), ctx.Query("page", "1"))
	if err != nil {
		if errors.Is(err, applicationsermon.ErrInvalidPage) {
			return fiber.ErrBadRequest
		}
		log.Printf("[ListSermons] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}
