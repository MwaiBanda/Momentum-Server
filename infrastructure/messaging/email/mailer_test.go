package email

import (
	"context"
	"testing"

	applicationtransaction "Momentum/application/service/transaction"
)

func TestNewTransactionMailerRequiresPassword(t *testing.T) {
	if mailer := NewTransactionMailer(""); mailer != nil {
		t.Fatal("NewTransactionMailer() returned a mailer without a password")
	}
}

func TestMailerReturnsConfigurationErrorWhenUnavailable(t *testing.T) {
	err := (*Mailer)(nil).Send(context.Background(), applicationtransaction.MailInput{})
	if err == nil {
		t.Fatal("Send() error = nil, want configuration error")
	}
}
