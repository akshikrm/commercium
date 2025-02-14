package types

import (
	"context"
	"net/http"
	"time"
)

type CreateCartRequest struct {
	UserID   uint32 `json:"user_id"`
	PriceID  uint32 `json:"price_id"`
	Quantity uint   `json:"quantity"`
}

type UpdateCartQuantityRequest struct {
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
	PriceID   uint32     `json:"price_id"`
	Quantity  uint       `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CartList struct {
	ID       uint32 `json:"id"`
	Quantity uint   `json:"quantity"`
	PriceID  string `json:"price_id"`
	Price    uint   `json:"price"`
	Product  struct {
		ID          uint32   `json:"id"`
		Name        string   `json:"name"`
		Slug        string   `json:"slug"`
		Description string   `json:"description"`
		Image       []string `json:"image"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type CartRepository interface {
	GetAll(uint32) ([]*CartList, bool)
	GetOne(uint32) (*CartList, bool)
	Create(*CreateCartRequest) (*Cart, bool)
	Update(uint32, *UpdateCartRequest) (*CartList, bool)
	Delete(uint32) bool
	CheckIfEntryExist(uint32, uint32) bool
	UpdateQuantity(*CreateCartRequest) bool
	HardDeleteByUserID(string) bool
}

type CartServicer interface {
	GetAll(uint32) ([]*CartList, error)
	GetOne(uint32) (*CartList, error)
	Create(*CreateCartRequest) error
	Update(uint32, *UpdateCartRequest) (*CartList, error)
	Delete(uint32) error
	HardDeleteByUserID(string) error
}

type CartHandler interface {
	GetAll(context.Context, http.ResponseWriter, *http.Request) error
	GetOne(context.Context, http.ResponseWriter, *http.Request) error
	Create(context.Context, http.ResponseWriter, *http.Request) error
	Update(context.Context, http.ResponseWriter, *http.Request) error
	Delete(context.Context, http.ResponseWriter, *http.Request) error
}
