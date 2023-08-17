package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

// GetTransactionsByUserId godoc
//
//	@Summary		Show a list of transactions belonging to a user
//	@Description	get a list of transactions belonging to a user
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			userId			path		string	true	"provide user Id"
//	@Success		200				{array}		model.Transaction
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
//	@Router			/api/v1/transactions/{userId} [get]
func (controller *Controller) GetTransactionsByUserId(context *fiber.Ctx) error {
	var transactions []model.Transaction
	res, err := controller.prisma.Transaction.FindMany(
		db.Transaction.UserID.Equals(context.Params("userId")),
	).Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &transactions); err != nil {
		fmt.Println(err)
	}

	return context.JSON(transactions)
}

// PostTransaction godoc
//
//	@Summary		Post a transaction
//	@Description	Used to post a transaction
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{object}	model.Transaction
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
//	@Router			/api/v1/transactions [post]
func (controller *Controller) PostTransaction(context *fiber.Ctx) error {
	transaction := new(model.Transaction)
	if err := context.BodyParser(transaction); err != nil {
		log.Panic(err.Error())
	}
	_, err := controller.prisma.Transaction.CreateOne(
		db.Transaction.Amount.Set(transaction.Amount),
		db.Transaction.Date.Set(transaction.Date),
		db.Transaction.Description.Set(transaction.Description),
		db.Transaction.User.Link(db.User.ID.Equals(transaction.UserId))).Exec(controller.context)
	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(transaction)
}

// DeleteTransactionsById godoc
//
//	@Summary		Delete a transaction
//	@Description	delete a transaction by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			transactionId	path		string	true	"provide transaction Id"
//	@Success		200				{object}	model.Transaction
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
//	@Router			/api/v1/transactions/{transactionId} [delete]
func (controller *Controller) DeleteTransactionsById(context *fiber.Ctx) error {
	var transaction model.Transaction
	res, err := controller.prisma.Transaction.FindUnique(
		db.Transaction.ID.Equals(context.Params("transactionId")),
	).Delete().Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &transaction); err != nil {
		fmt.Println(err)
	}

	return context.JSON(transaction)
}
