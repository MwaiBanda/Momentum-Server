package user

import (
	"context"
	"errors"

	"Momentum/application/telemetry"
	domainuser "Momentum/domain/user"
)

func (service *Service) Delete(ctx context.Context, id string) (Response, error) {
	ctx, span := telemetry.Start(ctx, "application.user.delete")
	defer span.End()

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
