package volunteer_service

import (
	"Momentum/application/cache"
	domainservice "Momentum/domain/volunteer_service"
)

type Service struct {
	repository domainservice.Repository
	cacheStore cache.Store
}

func GetInstance(repository domainservice.Repository, cacheStore cache.Store) *Service {
	return &Service{
		repository: repository,
		cacheStore: cacheStore,
	}
}
