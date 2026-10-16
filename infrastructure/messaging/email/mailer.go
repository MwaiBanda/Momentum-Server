package email

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	applicationtransaction "Momentum/application/service/transaction"
)

//go:embed transaction.html
var transactionTemplate string

type Mailer struct {
	password string
}

func NewTransactionMailer(password string) *Mailer {
	if password == "" {
		return nil
	}
	return &Mailer{password: password}
}

func (mailer *Mailer) Send(ctx context.Context, input applicationtransaction.MailInput) error {
	if mailer == nil || mailer.password == "" {
		return fmt.Errorf("transaction mailer is not configured")
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	const (
		from     = "bandamwai@gmail.com"
		smtpHost = "smtp.gmail.com"
		smtpPort = "587"
	)
	recipients := []string{from, "charlie@momentumindiana.org"}

	messageTemplate, err := template.New("transaction").Parse(transactionTemplate)
	if err != nil {
		return fmt.Errorf("parse transaction email template: %w", err)
	}

	var body bytes.Buffer
	if err := messageTemplate.Execute(&body, input); err != nil {
		return fmt.Errorf("render transaction email template: %w", err)
	}

	headers := strings.Join([]string{
		"From: " + from,
		"To: " + strings.Join(recipients, ", "),
		"Subject: New Payment",
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=\"UTF-8\"",
		"",
		"",
	}, "\r\n")

	auth := smtp.PlainAuth("", from, mailer.password, smtpHost)
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, []byte(headers+body.String())); err != nil {
		return fmt.Errorf("send transaction email: %w", err)
	}
	return nil
}
