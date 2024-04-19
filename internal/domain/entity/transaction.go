package entity

type Transaction struct {
	ID            uint    `json:"id"`
	Status        string  `json:"status"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	MerchantID    uint    `json:"merchant_id"`
	PaymentSource string  `json:"payment_source"`
	PaymentMethod string  `json:"payment_method"`
	ClientSecret  string  `json:"client_secret"`
}
