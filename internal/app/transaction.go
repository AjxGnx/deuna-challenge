package app

import (
	"fmt"

	"github.com/AjxGnx/deuna-challenge/internal/domain/dto"
	"github.com/AjxGnx/deuna-challenge/internal/domain/entity"
	"github.com/AjxGnx/deuna-challenge/internal/domain/model"
	"github.com/AjxGnx/deuna-challenge/internal/enums"
	"github.com/AjxGnx/deuna-challenge/internal/infra/adapters/db/repository"
	"github.com/AjxGnx/deuna-challenge/internal/infra/adapters/stripe"
	"github.com/AjxGnx/deuna-challenge/pkg"
	stripeClient "github.com/stripe/stripe-go/v78"
)

const ConvertCentToDollar = 100

type Transaction interface {
	CreatePaymentIntentWithStripe(transactionDTO dto.Transaction) (entity.Transaction, error)
	CreateRefundWithStripe(transactionID uint) (entity.Transaction, error)
	GetByID(transactionID uint) (entity.Transaction, error)
	Update(updateTransaction dto.UpdateTransaction) (entity.Transaction, error)
}

type transactionApp struct {
	stripe stripe.Stripe
	repo   repository.Transaction
}

func NewTransactionApp(stripe stripe.Stripe, repo repository.Transaction) Transaction {
	return transactionApp{
		stripe,
		repo,
	}
}

func (app transactionApp) CreatePaymentIntentWithStripe(transactionDTO dto.Transaction) (entity.Transaction, error) {
	paymentIntent, err := app.stripe.CreatePaymentIntent(stripeClient.PaymentIntentParams{
		Amount:   stripeClient.Int64(int64(transactionDTO.Amount * ConvertCentToDollar)),
		Currency: stripeClient.String(string(stripeClient.CurrencyUSD)),
	})

	if err != nil {
		return entity.Transaction{}, err
	}

	transaction, err := app.repo.Create(model.Transaction{
		Amount:        transactionDTO.Amount,
		CustomerID:    transactionDTO.CustomerID,
		MerchantID:    transactionDTO.MerchantID,
		Status:        string(paymentIntent.Status),
		PaymentSource: enums.Stripe,
		Type:          enums.Payment,
		PaymentID:     paymentIntent.ID,
	})
	if err != nil {
		return entity.Transaction{}, err
	}

	transaction.ClientSecret = paymentIntent.ClientSecret

	return transaction, nil

}

func (app transactionApp) CreateRefundWithStripe(transactionID uint) (entity.Transaction, error) {
	transaction, err := app.repo.GetByID(transactionID)
	if err != nil {
		return entity.Transaction{}, err
	}

	if transaction.Status == enums.Refunded {
		return entity.Transaction{}, fmt.Errorf("transaction %v has already been refunded", transactionID)
	}

	refund, err := app.stripe.CreateRefund(stripeClient.RefundParams{
		Amount:        stripeClient.Int64(int64(transaction.Amount * ConvertCentToDollar)),
		PaymentIntent: stripeClient.String(transaction.PaymentID),
	})
	if err != nil {
		return entity.Transaction{}, err
	}

	if refund.Status == enums.Succeeded {
		_, err = app.repo.Update(model.Transaction{
			ID:     transactionID,
			Status: enums.Refunded,
		})
		if err != nil {
			return entity.Transaction{}, err
		}
	}

	return app.repo.Create(model.Transaction{
		Amount:        transaction.Amount,
		CustomerID:    transaction.CustomerID,
		MerchantID:    transaction.MerchantID,
		Status:        string(refund.Status),
		PaymentSource: enums.Stripe,
		Type:          enums.Refund,
		PaymentMethod: transaction.PaymentMethod,
		PaymentID:     transaction.PaymentID,
		RefundID:      refund.ID,
	})
}

func (app transactionApp) GetByID(transactionID uint) (entity.Transaction, error) {
	transaction, err := app.repo.GetByID(transactionID)
	return transaction.ToEntity(), err
}

func (app transactionApp) Update(updateTransaction dto.UpdateTransaction) (entity.Transaction, error) {
	return app.repo.Update(model.Transaction{
		ID:            updateTransaction.ID,
		Status:        updateTransaction.Status,
		PaymentMethod: pkg.GetPaymentMethod(updateTransaction.PaymentMethod),
		ErrorCode:     updateTransaction.ErrorCode,
		ErrorType:     updateTransaction.ErrorType,
		ErrorMessage:  updateTransaction.ErrorMessage,
	})
}
