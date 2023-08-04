package model

type TransactionRequest struct {
	Amount   string `json:"amount"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type TransactionResponse struct {
	PublishableKey string `json:"publishableKey"`
	PaymentIntent  string `json:"paymentIntent"`
	Customer       string `json:"customer"`
	EphemeralKey   string `json:"ephemeralKey"`
}
