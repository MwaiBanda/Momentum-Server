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

type Service struct {
	sender Sender
}

func GetInstance(sender Sender) *Service {
	return &Service{sender: sender}
}

func (service *Service) Send(ctx context.Context, input Input) (Input, error) {
	if service == nil || service.sender == nil {
		return Input{}, ErrUnavailable
	}

	if err := service.sender.Send(ctx, input); err != nil {
		return Input{}, err
	}

	return input, nil
}
