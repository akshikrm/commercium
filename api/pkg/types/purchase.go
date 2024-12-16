package types

import "time"

type PurchaseRequest struct {
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	OrderID   string `json:"order_id"`
}

type Purchase struct {
	ID        uint32     `json:"id"`
	ProductID uint       `json:"product_id"`
	UserID    uint       `json:"user_id"`
	Price     uint       `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PurchaseList struct {
	ID      uint32 `json:"id"`
	Price   uint   `json:"purchase_price"`
	OrderID string `json:"order_id"`
	Product struct {
		ID    uint32 `json:"id"`
		Name  string `json:"name"`
		Price uint   `json:"price"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}
