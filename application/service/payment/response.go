package payment

type Input struct {
	Amount   int    `json:"amount"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Response struct {
	PublishableKey string `json:"publishableKey"`
	PaymentIntent  string `json:"paymentIntent"`
	Customer       string `json:"customer"`
	EphemeralKey   string `json:"ephemeralKey"`
}
