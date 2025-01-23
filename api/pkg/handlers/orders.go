package handlers

import (
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
)

type purchase struct {
	service types.PurchaseService
}

func (a *purchase) HandleTransactionHook(w http.ResponseWriter, r *http.Request) error {
	body := new(types.Body)

	if err := DecodeBody(r.Body, &body); err != nil {
		return err
	}

	switch body.EventType {
	case "transaction.created":
		{
			a.service.CreateTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction created...")
		}
	case "transaction.ready":
		{
			a.service.ReadyTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction ready...")
		}
	case "transaction.completed":
		{
			a.service.CompleteTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction completed...")
		}
	case "transaction.payment_failed":
		{
			a.service.FailedTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction failed...")
		}
	}
	return writeJson(w, http.StatusOK, "waiting...")
}

func (a *purchase) GetMyOrders(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := uint(ctx.Value("userID").(int))
	orders, err := a.service.GetOrdersByUserID(userID)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, orders)
}

func (a *purchase) GetOrderStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	txnId := r.PathValue("txnId")
	paddle := services.NewPaddlePayment()
	if err := paddle.Init(); err != nil {
		return err
	}

	transactionStatus, err := a.service.GetOrderStatus(txnId)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, transactionStatus)
}

func (a *purchase) GetInvoice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	txnId := r.PathValue("txnId")
	paddle := services.NewPaddlePayment()
	if err := paddle.Init(); err != nil {
		return err
	}

	invoiceURL := paddle.GetInvoice(txnId)
	return writeJson(w, http.StatusOK, *invoiceURL)
}

func newPurchase(service types.PurchaseService) types.PurchaseHandler {
	handler := new(purchase)
	handler.service = service
	return handler
}
