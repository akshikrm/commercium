package repository_test

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/types"
	"testing"
)

func TestCreateCart(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := repository.New(testConfig)
	cart := types.CreateCartRequest{
		UserID:   2,
		PriceID:  1,
		Quantity: 2,
	}
	_, ok := store.Cart.Create(&cart)
	if !ok {
		t.Error("failed to create product")
	}
}

func TestCartGetAll(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := repository.New(testConfig)
	_, ok := store.Cart.GetAll(2)
	if !ok {
		t.Error("failed to create product")
	}
}
