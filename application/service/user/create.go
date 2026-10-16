package user

import (
	"context"

	domainuser "Momentum/domain/user"
)

type CreateInput struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
}

func (service *Service) Create(ctx context.Context, input CreateInput) (Response, error) {
	userEntity, err := service.repository.Create(ctx, domainuser.CreateParams{
		ID:       input.ID,
		Email:    input.Email,
		Fullname: input.Fullname,
		Phone:    input.Phone,
	})
	if err != nil {
		return Response{}, err
	}

	service.cacheUser(ctx, userEntity)

	return toResponse(userEntity), nil
}
