package payment

type Input struct {
	Amount   int    `json:"amount" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type Response struct {
	PublishableKey string `json:"publishableKey" binding:"required"`
	PaymentIntent  string `json:"paymentIntent" binding:"required"`
	Customer       string `json:"customer" binding:"required"`
	EphemeralKey   string `json:"ephemeralKey" binding:"required"`
}
