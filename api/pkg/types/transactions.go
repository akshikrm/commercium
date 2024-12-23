package types

import "time"

type NewTransaction struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
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
