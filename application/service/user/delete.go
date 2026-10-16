package user

import (
	"context"
	"errors"

	domainuser "Momentum/domain/user"
)

func (service *Service) Delete(ctx context.Context, id string) (Response, error) {
	userEntity, err := service.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domainuser.ErrNotFound) {
			return Response{}, ErrNotFound
		}
		return Response{}, err
	}

	service.deleteCachedUser(ctx, userEntity.ID)
	return toResponse(userEntity), nil
}
