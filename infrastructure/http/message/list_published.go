package message

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

const adminUserID = "admin"

// ListAdmin godoc
//
//	@Summary		Show all published messages for the CMS
//	@Description	get published messages without user-specific passage notes
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Success		200	{object}	message.MessageListResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages/admin [get]
func (handler *Handler) ListAdmin(ctx fiber.Ctx) error {
	return handler.listPublished(ctx, adminUserID, "ListAdminMessages")
}

// ListPublishedByUser godoc
//
//	@Summary		Show published messages available for a user
//	@Description	get published messages and that user's passage notes
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			userId			path	string	true	"provide user Id"
//	@Success		200	{object}	message.MessageListResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages/{userId} [get]
func (handler *Handler) ListPublishedByUser(ctx fiber.Ctx) error {
	return handler.listPublished(ctx, ctx.Params("userId"), "ListPublishedMessages")
}

func (handler *Handler) listPublished(ctx fiber.Ctx, userID string, operation string) error {
	if handler.messages == nil {
		log.Printf("[%s] handler dependencies are not initialized", operation)
		return fiber.ErrInternalServerError
	}

	response, err := handler.messages.ListPublishedByUser(ctx.Context(), userID)
	if err != nil {
		return handler.mapError(operation, err)
	}

	return ctx.JSON(response)
}
