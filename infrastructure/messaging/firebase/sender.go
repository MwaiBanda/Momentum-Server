package firebase

import (
	"context"
	"fmt"

	applicationnotification "Momentum/application/service/notification"

	"firebase.google.com/go/v4/messaging"
)

type Sender struct {
	client *messaging.Client
}

func NewSender(client *messaging.Client) *Sender {
	return &Sender{client: client}
}

func (sender *Sender) Send(ctx context.Context, input applicationnotification.Input) error {
	if sender == nil || sender.client == nil {
		return applicationnotification.ErrUnavailable
	}

	_, err := sender.client.Send(ctx, &messaging.Message{
		Notification: &messaging.Notification{
			Title: input.Title,
			Body:  input.Body,
		},
		Topic: input.Topic,
	})
	if err != nil {
		return fmt.Errorf("send Firebase notification: %w", err)
	}
	return nil
}
