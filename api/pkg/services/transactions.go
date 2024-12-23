package services

import (
	"akshidas/e-com/pkg/types"
)

type TransactionStorager interface {
	NewTransaction(*types.NewTransaction) error
	TransactionReady(*types.TransactionReady) error
	UpdateStatus(string, string) error
	TransactionCompleted(*types.TransactionCompleted) error
}

type TransactionService struct {
	store TransactionStorager
}

func (t *TransactionService) CreateTransaction(data *types.Data) error {
	transaction := types.NewTransaction{
		TransactionID: data.ID,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt,
	}
	err := t.store.NewTransaction(&transaction)
	// if err == nil {
	// }

	return err
}

func (t *TransactionService) ReadyTransaction(data *types.Data) error {
	transaction := types.TransactionReady{
		TransactionID: data.ID,
		Status:        data.Status,
		CustomerID:    data.CustomerID,
	}
	return t.store.TransactionReady(&transaction)
}

func (t *TransactionService) CompleteTransaction(data *types.Data) error {
	transaction := types.TransactionCompleted{
		TransactionID: data.ID,
		Status:        data.Status,
		InvoiceNumber: data.InvoiceNumber,
	}
	return t.store.TransactionCompleted(&transaction)
}

func (t *TransactionService) FailedTransaction(data *types.Data) error {
	return t.store.UpdateStatus(data.ID, "failed")
}

func NewTransactionService(storage TransactionStorager) *TransactionService {
	return &TransactionService{
		store: storage,
	}
}
