package stripe

import (
	"github.com/AjxGnx/deuna-challenge/config"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/payout"
	"github.com/stripe/stripe-go/v78/refund"
	"github.com/stripe/stripe-go/v78/token"
)

type Stripe interface {
	CreateRefund(params stripe.RefundParams) (*stripe.Refund, error)
	CreatePaymentIntent(params stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	CreatePayout(params stripe.PayoutParams) (*stripe.Payout, error)
	CreateAccountToken(params stripe.BankAccountParams) (*stripe.Token, error)
}

type stripeAdapter struct {
}

func (s stripeAdapter) CreateAccountToken(params stripe.BankAccountParams) (*stripe.Token, error) {
	bankAccountTokenParams := &stripe.TokenParams{
		BankAccount: &params,
	}
	return token.New(bankAccountTokenParams)
}

func NewStripe() Stripe {
	stripe.Key = config.Environments().StripeKey
	return stripeAdapter{}
}

func (s stripeAdapter) CreateRefund(params stripe.RefundParams) (*stripe.Refund, error) {
	return refund.New(&params)
}

func (s stripeAdapter) CreatePaymentIntent(params stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return paymentintent.New(&params)
}

func (s stripeAdapter) CreatePayout(params stripe.PayoutParams) (*stripe.Payout, error) {
	return payout.New(&params)
}
