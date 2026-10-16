package transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	domaintransaction "Momentum/domain/transaction"
	domainuser "Momentum/domain/user"
)

type fakeRepository struct {
	transactions []domaintransaction.Transaction
	created      domaintransaction.Transaction
	createParams domaintransaction.CreateParams
	deleted      domaintransaction.Transaction
	deleteErr    error
}

func (repository *fakeRepository) List(context.Context) ([]domaintransaction.Transaction, error) {
	return repository.transactions, nil
}

func (repository *fakeRepository) ListByUser(context.Context, string) ([]domaintransaction.Transaction, error) {
	return repository.transactions, nil
}

func (repository *fakeRepository) Create(_ context.Context, params domaintransaction.CreateParams) (domaintransaction.Transaction, error) {
	repository.createParams = params
	return repository.created, nil
}

func (repository *fakeRepository) Delete(context.Context, string) (domaintransaction.Transaction, error) {
	return repository.deleted, repository.deleteErr
}

type fakeMailer struct {
	input MailInput
	err   error
}

func (mailer *fakeMailer) Send(_ context.Context, input MailInput) error {
	mailer.input = input
	return mailer.err
}

func TestCreateMapsResponseAndSendsMail(t *testing.T) {
	createdOn := time.Date(2026, time.October, 16, 0, 0, 0, 0, time.UTC)
	repository := &fakeRepository{created: domaintransaction.Transaction{
		ID:          "transaction-1",
		Amount:      100,
		Date:        "2026-10-16",
		CreatedOn:   createdOn,
		Description: "Meals, Outreach",
		User: domainuser.User{
			ID:       "user-1",
			Fullname: "Mwai Banda",
			Email:    "mwai@example.com",
		},
	}}
	mailer := &fakeMailer{}
	service := GetInstance(repository, mailer)

	response, err := service.Create(context.Background(), Input{Amount: 100})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if response.ID != "transaction-1" || response.User.ID != "user-1" {
		t.Fatalf("response = %#v, want mapped transaction", response)
	}
	if mailer.input.Email != "mwai@example.com" {
		t.Fatalf("mail input = %#v, want created user details", mailer.input)
	}
	if repository.createParams.ID == "" || repository.createParams.UserID == "" {
		t.Fatalf("Create() params = %#v, want generated transaction and user IDs", repository.createParams)
	}
	if mailer.input.Description != "Meals<br> Outreach" {
		t.Fatalf("mail description = %q, want comma replacement", mailer.input.Description)
	}
}

func TestDeleteMapsNotFound(t *testing.T) {
	service := GetInstance(&fakeRepository{deleteErr: domaintransaction.ErrNotFound}, nil)

	_, err := service.Delete(context.Background(), "transaction-1")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Delete() error = %v, want %v", err, ErrNotFound)
	}
}
