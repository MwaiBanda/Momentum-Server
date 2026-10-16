package meal

import (
	"Momentum/application/cache"
	domainmeal "Momentum/domain/meal"
)

type Service struct {
	repository domainmeal.Repository
	cacheStore cache.Store
}

func GetInstance(repository domainmeal.Repository, cacheStore cache.Store) *Service {
	return &Service{
		repository: repository,
		cacheStore: cacheStore,
	}
}
