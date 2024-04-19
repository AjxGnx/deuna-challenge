package handler

import (
	"net/http"
	"strconv"

	"github.com/AjxGnx/deuna-challenge/internal/app"
	"github.com/AjxGnx/deuna-challenge/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type Transaction interface {
	CreatePaymentIntentWithStripe(ctx echo.Context) error
	CreateRefundWithStripe(ctx echo.Context) error
	GetByID(ctx echo.Context) error
	Update(ctx echo.Context) error
}

type transactionHandler struct {
	app app.Transaction
}

func NewTransactionHandler(service app.Transaction) Transaction {
	return transactionHandler{
		service,
	}
}

// @Tags        Transactions
// @Summary     Create payment intent
// @Description Create a payment intent with stripe
// @Accept      json
// @Produce     json
// @Param       request body     dto.Transaction true "Request Body"
// @Success     200       {object} entity.Transaction
// @Failure     400  {object} dto.MessageError
// @Failure     500  {object} dto.MessageError
// @Router      /transactions/stripe/intent [post]
func (handler transactionHandler) CreatePaymentIntentWithStripe(ctx echo.Context) error {
	var transaction dto.Transaction

	if err := ctx.Bind(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := handler.app.CreatePaymentIntentWithStripe(transaction)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto.Message{
		Message: "your payment intent with stripe has been successfully created",
		Data:    response,
	})
}

// @Tags        Transactions
// @Summary     Create refund
// @Description Create a refund with stripe (to run this method the transaction needs to be succeeded in stripe)
// @Accept      json
// @Produce     json
// @Param       request body     dto.Refund true "Request Body"
// @Success     200       {object} entity.Transaction
// @Failure     400  {object} dto.MessageError
// @Failure     500  {object} dto.MessageError
// @Router      /transactions/stripe/refund [post]
func (handler transactionHandler) CreateRefundWithStripe(ctx echo.Context) error {
	var refund dto.Refund

	if err := ctx.Bind(&refund); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transaction, err := handler.app.CreateRefundWithStripe(refund.TransactionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto.Message{
		Message: "your refund with stripe has been successful",
		Data:    transaction,
	})

}

// @Tags        Transactions
// @Summary     Get by ID
// @Description Get transaction by ID
// @Accept      json
// @Produce     json
// @Param       id  path  int true "value of transaction to find"
// @Success     200       {object} entity.Transaction
// @Failure     500  {object} dto.MessageError
// @Router      /transactions/{id} [GET]
func (handler transactionHandler) GetByID(ctx echo.Context) error {
	transactionID, _ := strconv.Atoi(ctx.Param("id"))

	transaction, err := handler.app.GetByID(uint(transactionID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto.Message{
		Message: "transaction successfully loaded",
		Data:    transaction,
	})
}

// @Tags        Transactions
// @Summary     Update By ID
// @Description Update transaction by ID (to run this acction you need to know the state of a transaction "payment intent")
// @Accept      json
// @Produce     json
// @Param       request body     dto.UpdateTransaction true "Request Body"
// @Param       id  path  int true "value of transaction to update"
// @Success     200       {object} entity.Transaction
// @Failure     500  {object} dto.MessageError
// @Router      /transactions/{id} [PUT]
func (handler transactionHandler) Update(ctx echo.Context) error {
	var updateTransaction dto.UpdateTransaction

	transactionID, _ := strconv.Atoi(ctx.Param("id"))
	updateTransaction.ID = uint(transactionID)

	if err := ctx.Bind(&updateTransaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transaction, err := handler.app.Update(updateTransaction)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto.Message{
		Message: "transaction successfully updated",
		Data:    transaction,
	})
}
