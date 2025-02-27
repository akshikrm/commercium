package services_test

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"testing"
)

func TestProduct(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store, config)

	t.Run("create normal product", func(t *testing.T) {
		payload := types.NewProductRequest{
			Name:         "test product normal",
			Slug:         "test-product-normal",
			Image:        []string{"006_dk6ugp"},
			PrimaryImage: "https://res.cloudinary.com/commercium/image/upload/v1738853581/006_dk6ugp.jpg",
			Description:  "this is normal product description",
			CategoryID:   1,
			Status:       "enabled",
			Type:         "one-time",
			Price:        63000,
		}

		err := services.Product.Create(&payload)
		if err != nil {
			t.Errorf("one-time product create failed %s", err)
		}
	})

	t.Run("create subscription product", func(t *testing.T) {
		payload := types.NewProductRequest{
			Name:         "test product subscription",
			Slug:         "test-product-subscription",
			Image:        []string{"006_dk6ugp"},
			PrimaryImage: "https://res.cloudinary.com/commercium/image/upload/v1738853581/006_dk6ugp.jpg",
			Description:  "this is subscription product description",
			CategoryID:   1,
			Status:       "enabled",
			Type:         "subscription",
			SubscriptionPrice: []types.SubscriptionPrices{
				{Label: "Silver 1", Price: 10000, Interval: "week", Frequency: 1},
				{Label: "Silver 2", Price: 10000, Interval: "week", Frequency: 2},
				{Label: "Gold 1", Price: 50000, Interval: "month", Frequency: 1},
				{Label: "Gold 2", Price: 50000, Interval: "month", Frequency: 3},
				{Label: "Platinum 1", Price: 100000, Interval: "year", Frequency: 1},
				{Label: "Platinum 2", Price: 100000, Interval: "year", Frequency: 2},
			},
		}

		err := services.Product.Create(&payload)
		if err != nil {
			t.Errorf("one-time product create failed %s", err)
		}
	})

}
