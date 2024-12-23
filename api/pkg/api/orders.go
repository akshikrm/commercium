package api

import (
	"akshidas/e-com/pkg/db"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/storage"
	"akshidas/e-com/pkg/types"
	"context"
	"log"
	"net/http"
)

type PurchaseServeicer interface {
	PlaceOrder(uint) error
	GetOrdersByUserID(uint) ([]*types.OrderList, error)
	GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, error)
}

type TransactionStorage interface {
	NewTransaction(*types.NewTransaction) error
	TransactionReady(*types.TransactionReady) error
	UpdateStatus(string, string) error
	TransactionCompleted(*types.TransactionCompleted) error
}

type PurchaseApi struct {
	service      PurchaseServeicer
	transactions TransactionStorage
}

func (a *PurchaseApi) TransactionComplete(w http.ResponseWriter, r *http.Request) error {
	body := new(types.Body)

	if err := DecodeBody(r.Body, &body); err != nil {
		return err
	}

	switch body.EventType {
	case "transaction.created":
		{
			transaction := types.NewTransaction{
				TransactionID: body.Data.ID,
				Status:        body.Data.Status,
				CreatedAt:     body.Data.CreatedAt,
			}
			a.transactions.NewTransaction(&transaction)
			log.Printf("transaction %s", transaction.Status)
			return writeJson(w, http.StatusOK, "transaction created...")
		}
	case "transaction.ready":
		{
			transaction := types.TransactionReady{
				TransactionID: body.Data.ID,
				Status:        body.Data.Status,
				CustomerID:    body.Data.CustomerID,
			}
			a.transactions.TransactionReady(&transaction)
			log.Printf("transaction %s", transaction.Status)
			return writeJson(w, http.StatusOK, "transaction ready...")
		}
	case "transaction.completed":
		{
			transaction := types.TransactionCompleted{
				TransactionID: body.Data.ID,
				Status:        body.Data.Status,
				InvoiceNumber: body.Data.InvoiceNumber,
			}
			a.transactions.TransactionCompleted(&transaction)
			log.Printf("transaction %s", body.Data.Status)
			return writeJson(w, http.StatusOK, "transaction completed...")
		}
	case "transaction.payment_failed":
		{
			a.transactions.UpdateStatus(body.Data.ID, "failed")
			log.Printf("transaction %s", "failed")
			return writeJson(w, http.StatusOK, "transaction failed...")

		}
	}
	return writeJson(w, http.StatusOK, "waiting...")
}

func (a *PurchaseApi) GetMyOrders(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := uint(ctx.Value("userID").(int))
	orders, err := a.service.GetOrdersByUserID(userID)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, orders)
}

func (a *PurchaseApi) GetPurchasesByOrderID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	orderID, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	purchases, err := a.service.GetPurchaseByOrderID(uint(orderID))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, purchases)
}

func (a *PurchaseApi) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := uint(ctx.Value("userID").(int))
	err := a.service.PlaceOrder(userID)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "order placed")
}

func NewPurchaseApi(database *db.Storage) *PurchaseApi {
	purchaseStorage := storage.NewOrdersStorage(database.DB)
	cartStorage := storage.NewCartStorage(database.DB)
	transactionStorage := storage.NewTransactionsStorage(database.DB)
	purchaseService := services.NewOrderService(purchaseStorage, cartStorage)
	return &PurchaseApi{service: purchaseService, transactions: transactionStorage}
}
