package types

import (
	"context"
	"net/http"
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
	Create(*NewProductCategoryRequest) (*ProductCategory, bool)
	GetNames() ([]*ProductCategoryName, bool)
	GetAll(url.Values) ([]*ProductCategory, bool)
	GetOne(int) (*ProductCategory, bool)
	Update(int, *UpdateProductCategoryRequest) (*ProductCategory, bool)
	Delete(int) bool
}

type ProductCateogoryServicer interface {
	Create(*NewProductCategoryRequest) (*ProductCategory, error)
	GetAll(url.Values) ([]*ProductCategory, error)
	GetNames() ([]*ProductCategoryName, error)
	GetOne(int) (*ProductCategory, error)
	Update(int, *UpdateProductCategoryRequest) (*ProductCategory, error)
	Delete(int) error
}

type ProductCategoryHandler interface {
	GetAll(context.Context, http.ResponseWriter, *http.Request) error
	GetOne(context.Context, http.ResponseWriter, *http.Request) error
	Delete(context.Context, http.ResponseWriter, *http.Request) error
	Create(context.Context, http.ResponseWriter, *http.Request) error
	Update(context.Context, http.ResponseWriter, *http.Request) error
}
