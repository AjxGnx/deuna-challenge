package model

import (
	"github.com/AjxGnx/deuna-challenge/internal/domain/entity"
)

type Transaction struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	Type          string
	Amount        float64
	CustomerID    uint
	MerchantID    uint
	Status        string
	PaymentSource string
	PaymentMethod string
	PaymentID     string
	RefundID      string
	ErrorCode     string
	ErrorType     string
	ErrorMessage  string
	Customer      Customer `gorm:"foreignKey:CustomerID"`
	Merchant      Merchant `gorm:"foreignKey:MerchantID"`
}

func (transaction Transaction) ToEntity() entity.Transaction {
	return entity.Transaction{
		ID:            transaction.ID,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Amount:        transaction.Amount,
		PaymentSource: transaction.PaymentSource,
		PaymentMethod: transaction.PaymentMethod,
		MerchantID:    transaction.MerchantID,
	}
}
