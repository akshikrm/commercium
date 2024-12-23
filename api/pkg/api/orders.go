package api

import (
	"akshidas/e-com/pkg/db"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/storage"
	"akshidas/e-com/pkg/types"
	"context"
	"fmt"
	"log"
	"net/http"
)

type PurchaseServeicer interface {
	PlaceOrder(uint) error
	GetOrdersByUserID(uint) ([]*types.OrderList, error)
	GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, error)
	// GetAll(uint32, string) ([]*types.PurchaseList, error)
	// GetByOrderID(string) (*types.OrderView, error)
}

type PurchaseApi struct {
	service PurchaseServeicer
}

func (a *PurchaseApi) TransactionComplete(w http.ResponseWriter, r *http.Request) error {
	log.Println("transaction is being completed")
	body := new(types.Body)

	if err := DecodeBody(r.Body, &body); err != nil {
		return err
	}

	fmt.Println(body.EventID, body.EventType, body.NotificationID, body.Data.ID, body.Data.InvoiceID, body.Data.InvoiceNumber)
	return writeJson(w, http.StatusOK, "transaction completed...")

}

// func (a *PurchaseApi) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	id := uint32(ctx.Value("userID").(int))
// 	role := ctx.Value("role").(string)
//
// 	purchases, err := a.service.GetAll(id, role)
//
// 	if err != nil {
// 		return err
// 	}
// 	return writeJson(w, http.StatusOK, purchases)
// }

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

// func (a *PurchaseApi) GetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	id := r.PathValue("id")
//
// 	purchase, err := a.service.GetByOrderID(id)
// 	if err != nil {
// 		return err
// 	}
// 	return writeJson(w, http.StatusOK, purchase)
// }

func NewPurchaseApi(database *db.Storage) *PurchaseApi {
	purchaseStorage := storage.NewOrdersStorage(database.DB)
	cartStorage := storage.NewCartStorage(database.DB)
	purchaseService := services.NewOrderService(purchaseStorage, cartStorage)
	return &PurchaseApi{service: purchaseService}
}
