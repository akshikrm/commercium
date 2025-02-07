package types

import (
	"context"
	"net/http"
	"time"

	"github.com/PaddleHQ/paddle-go-sdk"
)

type PurchaseRequest struct {
	OrderID   uint32 `json:"order_id"`
	ProductID uint32 `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Price     uint   `json:"price"`
}

type NewOrder struct {
	TransactionID uint32 `json:"transaction_id"`
	PriceID       string `json:"price_id"`
	ProductID     string `json:"product_id"`
	Quantity      uint   `json:"quantity"`
	Amount        string `json:"amount"`
}

type OrderRequest struct {
	UserID  uint32 `json:"user_id"`
	OrderID string `json:"order_id"`
	Price   uint64 `json:"price"`
}

type PurchaseList struct {
	ID       uint32 `json:"id"`
	Price    uint   `json:"purchase_price"`
	Quantity uint   `json:"quantity"`
	Product  struct {
		ID   uint32 `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderProduct struct {
	ID             uint32 `json:"id"`
	ProductID      uint32 `json:"product_id"`
	ShippingStatus string `json:"shipping_status"`
	Name           string `json:"name"`
	Price          uint   `json:"price"`
	Quantity       uint   `json:"quantity"`
}

type OrderList struct {
	ID            uint           `json:"id"`
	TxnID         string         `json:"transaction_id"`
	InvoiceNumber string         `json:"invoice_number"`
	Total         string         `json:"total"`
	PaymentStatus string         `json:"payment_status"`
	Products      []OrderProduct `json:"products"`
	CreatedAt     time.Time      `json:"created_at"`
}

type OrderView struct {
	ID      uint32 `json:"id"`
	Price   uint   `json:"purchase_price"`
	OrderID string `json:"order_id"`
	Product struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Category    struct {
			ID   uint32 `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"category"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type ShippingStatus string

type ShippingInformation struct {
	ID            uint32         `json:"id"`
	Status        ShippingStatus `json:"status"`
	TransactionID string         `json:"transaction_id"`
	Amount        uint           `json:"amount"`
	Quantity      uint           `json:"quantity"`
	User          struct {
		ID        uint32 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	} `json:"user"`
	Product struct {
		ID   uint32 `json:"id"`
		Name string `json:"name"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type NewPrice struct {
	ID        string
	ProductID uint
	Amount    uint
	Label     string
}

type OrdersRepository interface {
	GetOrdersByUserID(uint32) ([]*OrderList, bool)
	GetPurchaseByOrderID(uint) ([]*PurchaseList, bool)
	CreateOrder([]*NewOrder) bool
	GetAllOrders() ([]*OrderList, bool)
	GetShippingInformation() ([]*ShippingInformation, bool)
	UpdateOrderStatus(uint, ShippingStatus) bool
}

type PurchaseHandler interface {
	CreateTransaction(context.Context, http.ResponseWriter, *http.Request) error
	HandleTransactionHook(http.ResponseWriter, *http.Request) error
	GetAllOrders(context.Context, http.ResponseWriter, *http.Request) error
	GetShippingInformation(context.Context, http.ResponseWriter, *http.Request) error
	GetOrderStatus(context.Context, http.ResponseWriter, *http.Request) error
	GetInvoice(context.Context, http.ResponseWriter, *http.Request) error
	UpdateShippingStatus(context.Context, http.ResponseWriter, *http.Request) error
}

type Transaction = *paddle.Transaction

type PaymentProvider interface {
	CreateCustomer(*CreateUserRequest) error
	GetCustomerByEmail(string) (string, error)
	GetInvoice(string) *string
	CreateProduct(*NewProductRequest) error
	CreateTransaction(string, []*CartList) (Transaction, error)
	CreatePrice(string, string, uint) *NewPrice
}
