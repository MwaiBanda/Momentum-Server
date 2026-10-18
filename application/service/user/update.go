package user

import (
	"context"

	domainuser "Momentum/domain/user"
)

type UpdateInput struct {
	ID       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

func (service *Service) Update(ctx context.Context, input UpdateInput) (Response, error) {
	userEntity, err := service.repository.Update(ctx, domainuser.UpdateParams{
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
