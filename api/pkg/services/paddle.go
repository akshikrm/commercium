package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"context"
	"github.com/PaddleHQ/paddle-go-sdk"
	"log"
	"os"
	"strconv"
)

type PaddlePayment struct {
	client *paddle.SDK
}

func (p *PaddlePayment) Init() error {
	paddle_key := os.Getenv("PADDLE_API_KEY")

	client, err := paddle.New(paddle_key, paddle.WithBaseURL(paddle.SandboxBaseURL))
	if err != nil {
		log.Printf("failed to connect to paddle due to %s", err)
		return utils.ServerError
	}
	p.client = client
	return nil
}

func (p *PaddlePayment) CreateProduct(newProduct *types.CreateNewProduct) error {
	ctx := context.Background()

	product, err := p.client.CreateProduct(ctx, &paddle.CreateProductRequest{
		Name:        newProduct.Name,
		Description: &newProduct.Description,
		TaxCategory: paddle.TaxCategoryStandard,
	})

	if err != nil {
		log.Printf("failed to add product to paddle due to %s", err)
		return utils.ServerError
	}

	price, err := p.client.CreatePrice(ctx, &paddle.CreatePriceRequest{
		ProductID: product.ID,
		UnitPrice: paddle.Money{
			Amount:       strconv.Itoa(int(newProduct.Price)),
			CurrencyCode: paddle.CurrencyCodeINR,
		},
		Description: "Main Price",
	})

	if err != nil {
		log.Printf("failed to add price to product due to %s", err)
		return utils.ServerError
	}

	newProduct.ProductID = product.ID
	if price, err := strconv.Atoi(price.UnitPrice.Amount); err != nil {
		log.Printf("failed to add price to product due to %s", err)
		return utils.ServerError
	} else {
		newProduct.Price = uint(price)
	}
	return nil
}
