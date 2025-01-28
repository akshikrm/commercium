package repository_test

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"testing"
)

func TestOrders(t *testing.T) {
	testConfig := config.NewTestConfig()
	store := repository.New(testConfig)
	t.Run("get shipping information", func(t *testing.T) {
		_, ok := store.Orders.GetShippingInformation()
		if !ok {
			t.Error("shipping information query failed")
		}
	})

}
