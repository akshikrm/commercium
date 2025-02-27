package services_test

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"testing"

	"github.com/PaddleHQ/paddle-go-sdk"
)

func TestPrice(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store, config)

	productID := "pro_01jn1r1wx63psxf9kvhxd947r7"

	t.Run("create one-time price", func(t *testing.T) {
		newPricePayload := types.NewPricePayload{
			ProductID: productID,
			Name:      "one time price",
			Price:     1000000,
		}
		res := services.PaymentProvider.CreatePrice(newPricePayload)
		if res == nil {
			t.Errorf("Failed to create price")
		}
	})

	t.Run("create subscription price", func(t *testing.T) {
		newPricePayload := types.NewPricePayload{
			ProductID: productID,
			Name:      "Platinum Price",
			Price:     10000,
			BillingCycle: &types.BillingCycle{
				Interval:  "month",
				Frequency: 1,
			},
		}
		res := services.PaymentProvider.CreatePrice(newPricePayload)
		if res == nil {
			t.Errorf("Failed to create price")
		}
	})
}

func TestPurchase(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store, config)

	custID := "ctm_01jj1fam90sch5sz7ex8xt18mv"

	t.Run("create one-time transaction", func(t *testing.T) {
		items := []paddle.CreateTransactionItems{}
		items = append(items,
			*paddle.NewCreateTransactionItemsCatalogItem(
				&paddle.CatalogItem{
					PriceID:  "pri_01jkr7h4e5f25z07gfyjs0fpz5",
					Quantity: 2,
				},
			),
		)
		newTransaction := types.NewTransactionPayload{
			CustomerID: custID,
			Items:      items,
		}
		_, err := services.PaymentProvider.CreateTransaction(&newTransaction)
		if err != nil {
			t.Errorf("Failed to create transaction")
		}
	})

	t.Run("create subscription transaction", func(t *testing.T) {
		items := []paddle.CreateTransactionItems{}
		items = append(items,
			*paddle.NewCreateTransactionItemsCatalogItem(
				&paddle.CatalogItem{
					PriceID:  "pri_01jkr7h4e5f25z07gfyjs0fpz5",
					Quantity: 2,
				},
			),
		)
		newTransaction := types.NewTransactionPayload{
			CustomerID: custID,
			Items:      items,
			BillingPeriod: &paddle.TimePeriod{
				StartsAt: "2025-01-31T00:00:00Z",
				EndsAt:   "2025-12-31T23:59:59Z",
			},
		}
		_, err := services.PaymentProvider.CreateTransaction(&newTransaction)
		if err != nil {
			t.Errorf("Failed to create transaction")
		}
	})

}
