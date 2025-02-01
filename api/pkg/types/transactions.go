package types

import "time"

type NewTransaction struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	Tax           string    `json:"tax"`
	SubTotal      string    `json:"sub_total"`
	GrandTotal    string    `json:"grand_total"`
}

type TransactionReady struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	CustomerID    string `json:"customer_id"`
}

type TransactionCompleted struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	InvoiceNumber string `json:"invoice_number"`
}

type TransactionRepository interface {
	NewTransaction(*NewTransaction) *uint32
	TransactionReady(*TransactionReady) bool
	UpdateStatus(string, string) bool
	TransactionCompleted(*TransactionCompleted) bool
	GetOrderStatus(string) string
}

type PurchaseService interface {
	CreateTransaction(*Data) error
	ReadyTransaction(*Data) error
	CompleteTransaction(*Data) error
	FailedTransaction(*Data) error
	GetOrderStatus(string) (string, error)
	GetPurchaseByOrderID(id uint) ([]*PurchaseList, error)
	GetOrdersByUserID(id uint32) ([]*OrderList, error)
	GetAllOrders() ([]*OrderList, error)
	GetShippingInformation() ([]*ShippingInformation, error)
	UpdateShippingStatus(uint, ShippingStatus) error
	NewTransaction(uint32) (string, error)
}
