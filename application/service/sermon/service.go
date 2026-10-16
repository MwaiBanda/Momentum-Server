package sermon

import "Momentum/application/cache"

type Service struct {
	provider   Provider
	cacheStore cache.Store
}

func GetInstance(provider Provider, cacheStore cache.Store) *Service {
	return &Service{
		provider:   provider,
		cacheStore: cacheStore,
	}
}
