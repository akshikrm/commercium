package types

import (
	"net/url"
	"time"
)

type NewProductCategoryRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type UpdateProductCategoryRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type ProductCategoryName struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductCategory struct {
	ID          uint32     `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	Enabled     bool       `json:"enabled"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ProductCategoriesRepository interface {
	Create(*NewProductCategoryRequest) (*ProductCategory, error)
	GetNames() ([]*ProductCategoryName, error)
	GetAll(url.Values) ([]*ProductCategory, error)
	GetOne(int) (*ProductCategory, error)
	Update(int, *UpdateProductCategoryRequest) (*ProductCategory, error)
	Delete(int) error
}

type ProductCateogriesServicer interface {
	Create(*NewProductCategoryRequest) (*ProductCategory, error)
	GetAll(url.Values) ([]*ProductCategory, error)
	GetNames() ([]*ProductCategoryName, error)
	GetOne(int) (*ProductCategory, error)
	Update(int, *UpdateProductCategoryRequest) (*ProductCategory, error)
	Delete(int) error
}
