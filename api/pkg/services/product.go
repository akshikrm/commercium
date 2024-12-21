package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/PaddleHQ/paddle-go-sdk"
)

type ProductStorager interface {
	GetAll(url.Values) ([]*types.ProductsList, error)
	GetOne(int) (*types.Product, error)
	Create(*types.CreateNewProduct) (*types.Product, error)
	Update(int, *types.CreateNewProduct) (*types.Product, error)
	Delete(int) error
}

type ProductService struct {
	productModel ProductStorager
}

func (r *ProductService) Get(filter url.Values) ([]*types.ProductsList, error) {
	return r.productModel.GetAll(filter)
}

func connectToPaddle() (*paddle.SDK, error) {
	paddle_key := os.Getenv("PADDLE_API_KEY")
	client, err := paddle.New(paddle_key, paddle.WithBaseURL(paddle.SandboxBaseURL))

	if err != nil {
		log.Printf("failed to connect to paddle due to %s", err)
		return nil, utils.ServerError
	}
	return client, nil
}

func (r *ProductService) addProductToPaddle(newProduct *types.CreateNewProduct) error {
	ctx := context.Background()
	client, err := connectToPaddle()

	if err != nil {
		log.Printf("failed to connect to paddle due to %s", err)
		return utils.ServerError
	}

	product, err := client.CreateProduct(ctx, &paddle.CreateProductRequest{
		Name:        newProduct.Name,
		Description: &newProduct.Description,
		TaxCategory: paddle.TaxCategoryStandard,
	})

	if err != nil {
		log.Printf("failed to add product to paddle due to %s", err)
		return utils.ServerError
	}

	price, err := client.CreatePrice(ctx, &paddle.CreatePriceRequest{
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

func (r *ProductService) Create(newProduct *types.CreateNewProduct) error {
	if err := r.addProductToPaddle(newProduct); err != nil {
		return err
	}
	product, err := r.productModel.Create(newProduct)
	fmt.Println(product.Name)
	if err != nil {
		return (err)
	}
	return err
}

func (r *ProductService) Update(id int, newProduct *types.CreateNewProduct) (*types.Product, error) {
	return r.productModel.Update(id, newProduct)
}

func (r *ProductService) GetOne(id int) (*types.Product, error) {
	return r.productModel.GetOne(id)
}

func (r *ProductService) Delete(id int) error {
	return r.productModel.Delete(id)
}

func NewProductService(productModel ProductStorager) *ProductService {
	return &ProductService{
		productModel: productModel,
	}
}
