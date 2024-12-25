package types

import "time"

type PurchaseRequest struct {
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
	Price     uint `json:"price"`
}

type NewOrder struct {
	TransactionID uint   `json:"transaction_id"`
	PriceID       string `json:"price_id"`
	ProductID     string `json:"product_id"`
	Quantity      uint   `json:"quantity"`
	Amount        string `json:"amount"`
}

type OrderRequest struct {
	UserID  uint   `json:"user_id"`
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
	ID        uint32 `json:"id"`
	ProductID uint32 `json:"product_id"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	Quantity  uint   `json:"quantity"`
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
