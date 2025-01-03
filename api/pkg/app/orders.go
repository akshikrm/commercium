package app

import (
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
)

type OrdersApi struct {
	service            types.PurchaseServicer
	transactionService types.TransactionServicer
}

func (a *OrdersApi) HandleTransactionHook(w http.ResponseWriter, r *http.Request) error {
	body := new(types.Body)

	if err := DecodeBody(r.Body, &body); err != nil {
		return err
	}

	switch body.EventType {
	case "transaction.created":
		{
			a.transactionService.CreateTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction created...")
		}
	case "transaction.ready":
		{
			a.transactionService.ReadyTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction ready...")
		}
	case "transaction.completed":
		{
			a.transactionService.CompleteTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction completed...")
		}
	case "transaction.payment_failed":
		{
			a.transactionService.FailedTransaction(&body.Data)
			return writeJson(w, http.StatusOK, "transaction failed...")
		}
	}
	return writeJson(w, http.StatusOK, "waiting...")
}

func (a *OrdersApi) GetMyOrders(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := uint(ctx.Value("userID").(int))
	orders, err := a.service.GetOrdersByUserID(userID)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, orders)
}

func (a *OrdersApi) GetOrderStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	txnId := r.PathValue("txnId")
	paddle := new(services.PaddlePayment)
	if err := paddle.Init(); err != nil {
		return err
	}

	transactionStatus, err := a.transactionService.GetOrderStatus(txnId)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, transactionStatus)
}

func (a *OrdersApi) GetInvoice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	txnId := r.PathValue("txnId")
	paddle := new(services.PaddlePayment)
	if err := paddle.Init(); err != nil {
		return err
	}

	invoiceURL := paddle.GetInvoice(txnId)
	return writeJson(w, http.StatusOK, *invoiceURL)
}

func NewOrdersApi(transactionService types.TransactionServicer, purchaseService types.PurchaseServicer) *OrdersApi {

	return &OrdersApi{
		service:            purchaseService,
		transactionService: transactionService,
	}
}
