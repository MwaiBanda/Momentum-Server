package model

type PaymentRequest struct {
	Amount   int    `json:"amount"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type PaymentResponse struct {
	PublishableKey string `json:"publishableKey"`
	PaymentIntent  string `json:"paymentIntent"`
	Customer       string `json:"customer"`
	EphemeralKey   string `json:"ephemeralKey"`
}
