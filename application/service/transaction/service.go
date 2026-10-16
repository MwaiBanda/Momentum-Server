package transaction

import (
	domaintransaction "Momentum/domain/transaction"
)

type Service struct {
	repository domaintransaction.Repository
	mailer     Mailer
}

func GetInstance(repository domaintransaction.Repository, mailer Mailer) *Service {
	return &Service{
		repository: repository,
		mailer:     mailer,
	}
}
