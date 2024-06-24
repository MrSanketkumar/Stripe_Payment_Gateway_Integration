package models

type CreateIntentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type PaymentIntent struct {
	ID       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type Refund struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
}
