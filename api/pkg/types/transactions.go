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
	TransactionReady(*TransactionReady) error
	UpdateStatus(string, string) error
	TransactionCompleted(*TransactionCompleted) error
	GetOrderStatus(string) string
}

type TransactionServicer interface {
	CreateTransaction(*Data) error
	ReadyTransaction(*Data) error
	CompleteTransaction(*Data) error
	FailedTransaction(*Data) error
	GetOrderStatus(string) (string, error)
}
