package handlers

import (
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
	"time"
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

//    status: "delivered" | "pending" | "in-transit"

const (
	pending   types.ShippingStatus = "pending"
	delivered types.ShippingStatus = "delivered"
	inTransit types.ShippingStatus = "in-transit"
)

func (a *purchase) GetShippingInformation(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	shippingInformation := []*types.ShippingInformation{
		{
			ID:       1,
			Status:   delivered,
			Amount:   4999,
			Quantity: 2,
			User: struct {
				ID    uint32 `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    101,
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			Product: struct {
				ID   uint32 `json:"id"`
				Name string `json:"name"`
			}{
				ID:   501,
				Name: "Wireless Mouse",
			},
			CreatedAt: time.Now().AddDate(0, 0, -1),
		},
		{
			ID:       2,
			Status:   inTransit,
			Amount:   29999,
			Quantity: 1,
			User: struct {
				ID    uint32 `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    102,
				Name:  "Jane Smith",
				Email: "jane.smith@example.com",
			},
			Product: struct {
				ID   uint32 `json:"id"`
				Name string `json:"name"`
			}{
				ID:   502,
				Name: "Gaming Keyboard",
			},
			CreatedAt: time.Now().AddDate(0, 0, -3),
		},
		{
			ID:       3,
			Status:   pending,
			Amount:   1999,
			Quantity: 3,
			User: struct {
				ID    uint32 `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    103,
				Name:  "Alice Johnson",
				Email: "alice.johnson@example.com",
			},
			Product: struct {
				ID   uint32 `json:"id"`
				Name string `json:"name"`
			}{
				ID:   503,
				Name: "USB-C Cable",
			},
			CreatedAt: time.Now().AddDate(0, 0, -7),
		},
	}
	return writeJson(w, http.StatusOK, &shippingInformation)
}

func newPurchase(service types.PurchaseService) types.PurchaseHandler {
	handler := new(purchase)
	handler.service = service
	return handler
}
