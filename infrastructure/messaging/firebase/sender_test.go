package firebase

import (
	"context"
	"errors"
	"testing"

	applicationnotification "Momentum/application/service/notification"
)

func TestSenderReturnsUnavailableWithoutClient(t *testing.T) {
	err := NewSender(nil).Send(context.Background(), applicationnotification.Input{})
	if !errors.Is(err, applicationnotification.ErrUnavailable) {
		t.Fatalf("Send() error = %v, want %v", err, applicationnotification.ErrUnavailable)
	}
}
