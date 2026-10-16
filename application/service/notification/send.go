package notification

import "context"

func (service *Service) Send(ctx context.Context, input Input) (Input, error) {
	if service == nil || service.sender == nil {
		return Input{}, ErrUnavailable
	}

	if err := service.sender.Send(ctx, input); err != nil {
		return Input{}, err
	}

	return input, nil
}
