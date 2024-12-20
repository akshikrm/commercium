package api

import (
	"akshidas/e-com/pkg/db"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/storage"
	"context"
	"net/http"
)

type PurchaseServeicer interface {
	PlaceOrder(uint) error
	// GetAll(uint32, string) ([]*types.PurchaseList, error)
	// GetByOrderID(string) (*types.OrderView, error)
}

type PurchaseApi struct {
	service PurchaseServeicer
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
	purchaseStorage := storage.NewPurchaseStorage(database.DB)
	cartStorage := storage.NewCartStorage(database.DB)
	purchaseService := services.NewPurchaseService(purchaseStorage, cartStorage)
	return &PurchaseApi{service: purchaseService}
}
