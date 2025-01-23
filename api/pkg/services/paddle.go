package services

import (
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

func (p *PaddlePayment) Init() error {
	paddle_key := os.Getenv("PADDLE_API_KEY")

	client, err := paddle.New(paddle_key, paddle.WithBaseURL(paddle.SandboxBaseURL))
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

func (p *PaddlePayment) CreateProduct(newProduct *types.NewProductRequest) error {
	ctx := context.Background()

	product, err := p.Client.CreateProduct(ctx, &paddle.CreateProductRequest{
		Name:        newProduct.Name,
		Description: &newProduct.Description,
		TaxCategory: paddle.TaxCategoryStandard,
		ImageURL:    &newProduct.PrimaryImage,
	})

	if err != nil {
		log.Printf("failed to add product to paddle due to %s", err)
		return utils.ServerError
	}

	price, err := p.Client.CreatePrice(ctx, &paddle.CreatePriceRequest{
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
	if amount, err := strconv.Atoi(price.UnitPrice.Amount); err != nil {
		log.Printf("failed to add price to product due to %s", err)
		return utils.ServerError
	} else {
		newProduct.Price = uint(amount)
		newProduct.PriceID = price.ID
	}
	return nil
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

func NewPaddlePayment() *PaddlePayment {
	p := new(PaddlePayment)
	return p
}
