package payment

type Service struct {
	provider       Provider
	publishableKey string
}

func GetInstance(provider Provider, publishableKey string) *Service {
	return &Service{
		provider:       provider,
		publishableKey: publishableKey,
	}
}
