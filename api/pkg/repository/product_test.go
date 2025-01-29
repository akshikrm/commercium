package repository

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/types"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := New(testConfig)
	product := types.NewProductRequest{
		Name:        "Test",
		Slug:        "test",
		Description: "Test",
		Image:       []string{},
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

func TestGetOneProduct(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := New(testConfig)
	_, ok := store.Product.GetOne(1)
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

func TestUpdateProduct(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := New(testConfig)
	product := types.NewProductRequest{
		Name:        "Test",
		Slug:        "test",
		Description: "Test",
		Image:       []string{"https://picsum.photos/id/12/2500/1667", "https://res.cloudinary.com/commercium/image/upload/v1737622147/005_uqebq5.jpg"},
		Price:       34,
		CategoryID:  1,
		PriceID:     "Test",
		ProductID:   "Test",
	}
	_, ok := store.Product.Update(13, &product)
	if !ok {
		t.Error("failed to update product")
	}
}
