package message

import (
	"Momentum/application/cache"
	domainmessage "Momentum/domain/message"
)

type Service struct {
	repository domainmessage.Repository
	cacheStore cache.Store
}

func GetInstance(repository domainmessage.Repository, cacheStore cache.Store) *Service {
	return &Service{
		repository: repository,
		cacheStore: cacheStore,
	}
}
