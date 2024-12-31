package types

import "time"

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

type Product struct {
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
