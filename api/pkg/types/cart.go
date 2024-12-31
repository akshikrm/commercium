package types

import "time"

type CreateCartRequest struct {
	UserID    uint32 `json:"user_id"`
	ProductID uint32 `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity uint `json:"quantity"`
}

type Cart struct {
	ID        uint32     `json:"id"`
	UserID    uint32     `json:"user_id"`
	ProductID uint32     `json:"product_id"`
	Quantity  uint       `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CartList struct {
	ID       uint32 `json:"id"`
	Quantity uint   `json:"quantity"`
	PriceID  string `json:"price_id"`
	Product  struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Price       uint   `json:"price"`
		Description string `json:"description"`
		Image       string `json:"image"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}
