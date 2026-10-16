package user

import (
	"Momentum/application/cache"
	domainuser "Momentum/domain/user"
)

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
