package dto

type Transaction struct {
	Amount     float64 `json:"amount"`
	MerchantID uint    `json:"merchant_id"`
	CustomerID uint    `json:"customer_id"`
}

type UpdateTransaction struct {
	ID            uint   `json:"id"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	ErrorCode     string `json:"error_code"`
	ErrorType     string `json:"error_type"`
	ErrorMessage  string `json:"error_message"`
}

type Refund struct {
	TransactionID uint `json:"transaction_id"`
}
