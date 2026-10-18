package transaction

import (
	"context"
	"log"
	"strings"

	"Momentum/application/telemetry"
	domaintransaction "Momentum/domain/transaction"

	"github.com/lucsky/cuid"
)

func (service *Service) Create(ctx context.Context, input Input) (Response, error) {
	ctx, span := telemetry.Start(ctx, "application.transaction.create")
	defer span.End()

	if service == nil || service.repository == nil || input.Amount < 0 {
		return Response{}, ErrInvalidInput
	}

	transactionID := input.ID
	if transactionID == "" {
		transactionID = cuid.New()
	}
	userID := input.UserID
	if userID == "" {
		userID = cuid.New()
	}

	params := domaintransaction.CreateParams{
		ID:          transactionID,
		Amount:      input.Amount,
		Date:        input.Date,
		Description: input.Description,
		UserID:      userID,
		Fullname:    pointerValue(input.Fullname),
		Email:       pointerValue(input.Email),
		Phone:       pointerValue(input.Phone),
	}
	created, err := service.repository.Create(ctx, params)
	if err != nil {
		return Response{}, err
	}

	if service.mailer != nil {
		if err := service.mailer.Send(ctx, MailInput{
			Amount:      created.Amount,
			Description: strings.ReplaceAll(created.Description, ",", "<br>"),
			Fullname:    created.User.Fullname,
			Email:       created.User.Email,
			Phone:       created.User.Phone,
		}); err != nil {
			log.Printf("[transaction-mail] send receipt for %q: %v", created.ID, err)
		}
	}

	return toResponse(created), nil
}

func pointerValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
