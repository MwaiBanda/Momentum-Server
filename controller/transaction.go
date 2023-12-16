package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
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
//	@Success		200				{array}		model.TransactionResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/transactions/{userId} [get]
func (controller *Controller) GetTransactionsByUserId(context *fiber.Ctx) error {
	var transactions []model.TransactionResponse
	res, err := controller.PrismaClient.Transaction.FindMany(
		db.Transaction.UserID.Equals(context.Params("userId")),
	).With(db.Transaction.User.Fetch()).Exec(controller.Context)
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
//	@Success		200				{object}	model.TransactionRequest
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/transactions [post]
func (controller *Controller) PostTransaction(context *fiber.Ctx) error {
	transaction := new(model.TransactionRequest)
	transactionResponse := new(model.TransactionResponse)
	if err := context.BodyParser(transaction); err != nil {
		log.Panic(err.Error())
	}
	res, err := controller.PrismaClient.Transaction.CreateOne(
		db.Transaction.Amount.Set(transaction.Amount),
		db.Transaction.Date.Set(transaction.Date),
		db.Transaction.Description.Set(transaction.Description),
		db.Transaction.User.Link(db.User.ID.Equals(transaction.UserId)),
	).With(
		db.Transaction.User.Fetch(),
	).Exec(controller.Context)
	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, transactionResponse); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		log.Panic(err.Error())
	}

	controller.PostTransactionMail(model.TransactionMailRequest{
		Amount:      transactionResponse.Amount,
		Description: strings.ReplaceAll(transactionResponse.Description, ",", "<br>"),
		Fullname:    transactionResponse.User.Fullname,
		Email:       transactionResponse.User.Email,
		Phone:       transactionResponse.User.Phone,
	})

	return context.JSON(transactionResponse)
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
//	@Success		200				{object}	model.TransactionResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/transactions/{transactionId} [delete]
func (controller *Controller) DeleteTransactionsById(context *fiber.Ctx) error {

	var transaction model.TransactionResponse
	res, err := controller.PrismaClient.Transaction.FindUnique(
		db.Transaction.ID.Equals(context.Params("transactionId")),
	).Delete().Exec(controller.Context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &transaction); err != nil {
		fmt.Println(err)
	}

	return context.JSON(transaction)
}
