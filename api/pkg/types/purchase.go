package types

import "time"

type PurchaseRequest struct {
	ProductID uint `json:"product_id"`
	UserID    uint `json:"user_id"`
	OrderID   uint `json:"order_id"`
}

type OrderRequest struct {
	UserID  uint   `json:"user_id"`
	OrderID string `json:"order_id"`
	Price   uint64 `json:"price"`
}

type PurchaseList struct {
	ID      uint32 `json:"id"`
	Price   uint   `json:"purchase_price"`
	OrderID string `json:"order_id"`
	Product struct {
		ID   uint32 `json:"id"`
		Name string `json:"name"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderView struct {
	ID      uint32 `json:"id"`
	Price   uint   `json:"purchase_price"`
	OrderID string `json:"order_id"`
	Product struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Category    struct {
			ID   uint32 `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"category"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}
