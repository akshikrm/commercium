package types

import (
	"net/url"
	"time"
)

type CreateNewProduct struct {
	Name        string `json:"name"`
	ProductID   string `json:"product_id"`
	CategoryID  uint   `json:"category_id"`
	Slug        string `json:"slug"`
	Price       uint   `json:"price"`
	PriceID     string `json:"price_id"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type ProductsList struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Price       string    `json:"price"`
	Image       string    `json:"image"`
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
	Image       string     `json:"image"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ProductRepository interface {
	GetAll(url.Values) ([]*ProductsList, error)
	GetOne(int) (*OneProduct, error)
	Create(*CreateNewProduct) (*OneProduct, error)
	Update(int, *CreateNewProduct) (*OneProduct, error)
	Delete(int) error
}

type ProductServicer interface {
	Get(url.Values) ([]*ProductsList, error)
	GetOne(int) (*OneProduct, error)
	Create(*CreateNewProduct) error
	Update(int, *CreateNewProduct) (*OneProduct, error)
	Delete(int) error
}
