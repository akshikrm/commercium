package services

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PaddleHQ/paddle-go-sdk"
)

type PaddlePayment struct {
	Client *paddle.SDK
}

func (p *PaddlePayment) Init(paddleApiKey string) error {
	client, err := paddle.New(paddleApiKey, paddle.WithBaseURL(paddle.SandboxBaseURL))
	if err != nil {
		log.Printf("failed to connect to paddle due to %s", err)
		return utils.PaddleError
	}
	p.Client = client
	return nil
}

func (p *PaddlePayment) CreateCustomer(newUser *types.CreateUserRequest) error {
	ctx := context.Background()
	customerName := fmt.Sprintf("%s %s", newUser.FirstName, newUser.LastName)
	customer, err := p.Client.CreateCustomer(ctx, &paddle.CreateCustomerRequest{
		Name:  &customerName,
		Email: newUser.Email,
	})
	if err != nil {
		if ok := strings.Contains(err.Error(), paddle.ErrCustomerAlreadyExists.Error()); ok {
			return paddle.ErrCustomerAlreadyExists
		}
		return utils.PaddleError
	}
	newUser.CustomerID = customer.ID
	return nil
}

func (p *PaddlePayment) GetCustomerByEmail(email string) (string, error) {
	ctx := context.Background()
	emails := []string{email}
	customer, err := p.Client.ListCustomers(ctx, &paddle.ListCustomersRequest{
		Email: emails,
	})
	if err != nil {
		fmt.Println("user not found")
	}
	var customerID string
	err = customer.Iter(ctx, func(v *paddle.Customer) (bool, error) {
		customerID = v.ID
		return true, nil
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return customerID, nil
}

func (p *PaddlePayment) GetInvoice(txnId string) *string {
	ctx := context.Background()
	res, err := p.Client.GetTransactionInvoice(ctx, &paddle.GetTransactionInvoiceRequest{
		TransactionID: txnId,
	})
	if err != nil {
		log.Printf("failed to get invoice due to %s", err)
		return nil
	}
	return &res.URL
}

func (p *PaddlePayment) CreatePrice(productID, amount, name string) *paddle.Price {
	ctx := context.Background()
	priceRequest := new(paddle.CreatePriceRequest)
	priceRequest.ProductID = productID
	priceRequest.Description = "Price"
	priceRequest.UnitPrice = paddle.Money{
		Amount:       amount,
		CurrencyCode: paddle.CurrencyCodeINR,
	}

	if name != "" {
		priceRequest.Name = &name
	}

	price, err := p.Client.CreatePrice(ctx, priceRequest)

	if err != nil {
		log.Printf("failed to add price to product due to %s", err)
		return nil
	}
	return price
}

func (p *PaddlePayment) CreateProduct(newProductRequest *types.NewProductRequest) error {
	ctx := context.Background()

	paddleProductRequest := new(paddle.CreateProductRequest)
	paddleProductRequest.Name = newProductRequest.Name
	paddleProductRequest.Description = &newProductRequest.Description
	paddleProductRequest.TaxCategory = paddle.TaxCategoryStandard

	if len(newProductRequest.PrimaryImage) > 0 {
		paddleProductRequest.ImageURL = &newProductRequest.PrimaryImage
	}
	product, err := p.Client.CreateProduct(ctx, paddleProductRequest)

	if err != nil {
		log.Printf("failed to add product to paddle due to %s", err)
		return utils.ServerError
	}

	newProductRequest.ProductID = product.ID
	if newProductRequest.Type == types.OneTimeProduct {
		price := p.CreatePrice(product.ID, strconv.Itoa(int(newProductRequest.Price)), "")
		if price == nil {
			return utils.PaddleError
		}

		if amount, err := strconv.Atoi(price.UnitPrice.Amount); err != nil {
			log.Printf("failed to add price to product due to %s", err)
			return utils.ServerError
		} else {
			newProductRequest.Price = uint(amount)
			newProductRequest.PriceID = price.ID
		}
		return nil
	}

	for _, priceItem := range newProductRequest.SubscriptionPrice {
		price := p.CreatePrice(product.ID, priceItem.Value, priceItem.Label)
		fmt.Println(product.ID, price.ID)
	}
	return nil

}

func (p PaddlePayment) CreateTransaction(customerID string, carts []*types.CartList) (*paddle.Transaction, error) {
	transactionItems := []paddle.CreateTransactionItems{}
	for _, cart := range carts {
		transactionItems = append(transactionItems,
			*paddle.NewCreateTransactionItemsCatalogItem(
				&paddle.CatalogItem{
					PriceID:  cart.PriceID,
					Quantity: int(cart.Quantity),
				},
			),
		)
	}
	newTransaction := new(paddle.CreateTransactionRequest)
	newTransaction.Items = transactionItems
	newTransaction.CustomerID = paddle.PtrTo(customerID)

	ctx := context.Background()
	txn, err := p.Client.CreateTransaction(ctx, newTransaction)
	if err != nil {
		log.Printf("failed to create transaction due to %s", err)
		return nil, utils.PaddleError
	}

	return txn, nil
}

func (p PaddlePayment) SyncPrice(store *repository.Storage) {
	rows, err := store.DB.Query("select id, product_id, price, price_id from products;")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	log.Println("starting sync...")
	for rows.Next() {
		var productID string
		var priceID string
		var price uint
		var id uint
		rows.Scan(&id, &productID, &price, &priceID)
		if priceID == "" {
			log.Printf("adding price %d to %s", price, productID)
			paddlePrice, err := p.Client.CreatePrice(ctx, &paddle.CreatePriceRequest{
				ProductID: productID,
				UnitPrice: paddle.Money{
					Amount:       strconv.Itoa(int(price)),
					CurrencyCode: paddle.CurrencyCodeINR,
				},
				Description: "Main Price",
			})

			if err != nil {
				log.Printf("failed to create price for product %s due to %s", productID, err)
				continue
			}
			_, err = store.DB.Query("update products set price_id=$1 where product_id=$2;", paddlePrice.ID, productID)
			if err != nil {
				log.Printf("failed to update price for product %s due to %s", productID, err)
			}
		} else {
			log.Printf("skipping %s because price already exists", productID)
		}
	}
	log.Println("sync complete")
}

func NewPaddlePayment(config *config.Config) *PaddlePayment {
	p := new(PaddlePayment)
	if err := p.Init(config.Paddle_Key); err != nil {
		log.Printf("Failed to initialize paddle due to %s", err)
		os.Exit(1)
	}
	return p
}
