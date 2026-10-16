package user

import (
	"context"
	"log"

	"Momentum/application/cache"
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

	if service.cacheStore != nil {
		if err := cache.Set(ctx, service.cacheStore, userCacheKeyPrefix+userEntity.ID, userEntity, userCacheTTL); err != nil {
			log.Printf("[user-cache] set %q: %v", userEntity.ID, err)
		}
	}

	return toResponse(userEntity), nil
}
