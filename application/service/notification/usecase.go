package notification

import "context"

type Input struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

type Response struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

type Sender interface {
	Send(ctx context.Context, input Input) error
}

type UseCase interface {
	Send(ctx context.Context, input Input) (Response, error)
}
