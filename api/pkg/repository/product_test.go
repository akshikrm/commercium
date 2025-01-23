package repository

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/types"
	"testing"
)

func TestGetOneProduct(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := New(testConfig)
	_, ok := store.Product.GetOne(39)
	if !ok {
		t.Error("failed to get product")
	}
}

func TestGetAllProduct(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := New(testConfig)
	_, ok := store.Product.GetAll(nil)
	if !ok {
		t.Error("failed to get all product")
	}
}

func TestCreateProduct(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := New(testConfig)
	product := types.NewProductRequest{
		Name:        "Test",
		Slug:        "test",
		Description: "Test",
		Image:       []string{"test", "test2", "test3", "test4"},
		Price:       34,
		CategoryID:  1,
		PriceID:     "Test",
		ProductID:   "Test",
	}
	_, ok := store.Product.Create(&product)
	if !ok {
		t.Error("failed to create product")
	}
}
