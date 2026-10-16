package user

import (
	"context"

	"Momentum/application/cache"
	domainuser "Momentum/domain/user"
)

type UseCase interface {
	GetByID(ctx context.Context, id string) (Response, error)
}

type Service struct {
	repository domainuser.Repository
	cacheStore cache.Store
}

func GetInstance(repository domainuser.Repository, cacheStore cache.Store) *Service {
	return &Service{
		repository: repository,
		cacheStore: cacheStore,
	}
}
