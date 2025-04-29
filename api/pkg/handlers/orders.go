package handlers

import (
	"akshidas/e-com/pkg/types"
	"context"
	"fmt"
	"net/http"
)

type purchase struct {
	service         types.PurchaseService
	cartService     types.CartServicer
	paymentProvider types.PaymentProvider
}

func (a *purchase) HandleTransactionHook(w http.ResponseWriter, r *http.Request) error {
	body := new(types.Body)

	if err := DecodeBody(r.Body, &body); err != nil {
		return err
	}

	fmt.Println(body.EventType, "event type")
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

func (a *purchase) GetAllOrders(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := ctx.Value("userID")
	role := ctx.Value("role")

	var orders = []*types.OrderList{}
	var err error = nil
	if role == "admin" {
		if orders, err = a.service.GetAllOrders(); err != nil {
			return err
		}
		return writeJson(w, http.StatusOK, orders)
	}

	if orders, err = a.service.GetOrdersByUserID(userID.(uint32)); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, orders)
}

func (a *purchase) GetOrderStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	txnId := r.PathValue("txnId")

	transactionStatus, err := a.service.GetOrderStatus(txnId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, transactionStatus)
}

func (a *purchase) GetInvoice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	txnId := r.PathValue("txnId")

	invoiceURL := a.paymentProvider.GetInvoice(txnId)
	return writeJson(w, http.StatusOK, *invoiceURL)
}

func (a *purchase) GetShippingInformation(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	shippingInformation, err := a.service.GetShippingInformation()
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, &shippingInformation)
}

func (a *purchase) UpdateShippingStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	orderID, err := parseId(r.PathValue("orderId"))
	if err != nil {
		return err
	}

	var shippingStatus types.ShippingStatus
	if err := DecodeBody(r.Body, &shippingStatus); err != nil {
		return err
	}

	if err := a.service.UpdateShippingStatus(uint(orderID), shippingStatus); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "updated shipping status")
}

func (a *purchase) CreateTransaction(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := ctx.Value("userID")
	txnID, err := a.service.NewTransaction(userID.(uint32))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, txnID)
}

func newPurchase(service types.PurchaseService, cartService types.CartServicer, paymentProvider types.PaymentProvider) types.PurchaseHandler {
	handler := new(purchase)
	handler.service = service
	handler.cartService = cartService
	handler.paymentProvider = paymentProvider

	return handler
}
