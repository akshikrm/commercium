package types

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type NewProductRequest struct {
	Name        string   `json:"name"`
	ProductID   string   `json:"product_id"`
	CategoryID  uint     `json:"category_id"`
	Slug        string   `json:"slug"`
	Status      string   `json:"status"`
	Type        string   `json:"type"`
	Price       uint     `json:"price"`
	PriceID     string   `json:"price_id"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
}

type ProductsList struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Price       uint      `json:"price"`
	Image       *string   `json:"image"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Category    struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	} `json:"category"`
}

type OneProduct struct {
	ID          uint       `json:"id"`
	CategoryID  uint       `json:"category_id"`
	ProductID   string     `json:"product_id"`
	PriceID     string     `json:"price_id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Price       uint       `json:"price"`
	Image       []string   `json:"image"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ProductRepository interface {
	GetAll(url.Values) ([]*ProductsList, bool)
	GetOne(int) (*OneProduct, bool)
	Create(*NewProductRequest) (*OneProduct, bool)
	Update(int, *NewProductRequest) (*OneProduct, bool)
	Delete(int) bool
	InsertImages(uint32, []string) bool
}

type ProductServicer interface {
	Get(url.Values) ([]*ProductsList, error)
	GetOne(int) (*OneProduct, error)
	Create(*NewProductRequest) error
	Update(int, *NewProductRequest) (*OneProduct, error)
	Delete(int) error
}

type ProductHandler interface {
	GetAll(context.Context, http.ResponseWriter, *http.Request) error
	GetOne(context.Context, http.ResponseWriter, *http.Request) error
	Delete(context.Context, http.ResponseWriter, *http.Request) error
	Create(context.Context, http.ResponseWriter, *http.Request) error
	Update(context.Context, http.ResponseWriter, *http.Request) error
}
