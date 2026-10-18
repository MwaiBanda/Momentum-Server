package notification

import (
	"context"
	"errors"
	"testing"
)

type fakeSender struct {
	input Input
	err   error
}

func (sender *fakeSender) Send(_ context.Context, input Input) error {
	sender.input = input
	return sender.err
}

func TestSendForwardsInputAndReturnsResponse(t *testing.T) {
	sender := &fakeSender{}
	service := GetInstance(sender)
	input := Input{Title: "Title", Body: "Body", Topic: "all-users"}

	response, err := service.Send(context.Background(), input)
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	wantResponse := Response{Title: input.Title, Body: input.Body, Topic: input.Topic}
	if response != wantResponse {
		t.Fatalf("Send() response = %#v, want %#v", response, wantResponse)
	}
	if sender.input != input {
		t.Fatalf("sender input = %#v, want %#v", sender.input, input)
	}
}

func TestSendReturnsSenderError(t *testing.T) {
	wantErr := errors.New("firebase unavailable")
	service := GetInstance(&fakeSender{err: wantErr})

	_, err := service.Send(context.Background(), Input{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("Send() error = %v, want %v", err, wantErr)
	}
}

func TestSendReturnsUnavailableWithoutSender(t *testing.T) {
	_, err := GetInstance(nil).Send(context.Background(), Input{})
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("Send() error = %v, want %v", err, ErrUnavailable)
	}
}
