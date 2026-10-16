package user

import (
	"context"
	"database/sql"
	"errors"
)

func (service *Service) Delete(ctx context.Context, id string) (Response, error) {
	userEntity, err := service.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}, ErrNotFound
		}
		return Response{}, err
	}

	service.deleteCachedUser(ctx, userEntity.ID)
	return toResponse(userEntity), nil
}
