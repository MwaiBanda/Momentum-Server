package notification

import "context"

type Input struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}

type Sender interface {
	Send(ctx context.Context, input Input) error
}

type UseCase interface {
	Send(ctx context.Context, input Input) (Input, error)
}
