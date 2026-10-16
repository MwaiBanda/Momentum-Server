package notification

type Service struct {
	sender Sender
}

func GetInstance(sender Sender) *Service {
	return &Service{sender: sender}
}
