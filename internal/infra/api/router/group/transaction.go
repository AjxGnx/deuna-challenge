package group

import (
	"github.com/AjxGnx/deuna-challenge/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

const transactionPath = "/transactions/"

type Transaction interface {
	Resource(c *echo.Group)
}

type transactionGroup struct {
	transactionHandler handler.Transaction
}

func NewTransactionGroup(transactionHandler handler.Transaction) Transaction {
	return transactionGroup{
		transactionHandler,
	}
}

func (routes transactionGroup) Resource(c *echo.Group) {
	groupPath := c.Group(transactionPath)
	groupPath.POST("stripe/intent", routes.transactionHandler.CreatePaymentIntentWithStripe)
	groupPath.POST("stripe/refund", routes.transactionHandler.CreateRefundWithStripe)
	groupPath.GET(":id", routes.transactionHandler.GetByID)
	groupPath.PUT(":id", routes.transactionHandler.Update)
}
