package notification

import "context"

func (service *Service) Send(ctx context.Context, input Input) (Response, error) {
	if service == nil || service.sender == nil {
		return Response{}, ErrUnavailable
	}

	if err := service.sender.Send(ctx, input); err != nil {
		return Response{}, err
	}

	return Response{Title: input.Title, Body: input.Body, Topic: input.Topic}, nil
}
