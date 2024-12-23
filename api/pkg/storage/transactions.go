package storage

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"log"
)

type TransactionsStorage struct {
	store *sql.DB
}

func (m *TransactionsStorage) UpdateStatus(txnID, status string) error {
	query := "update transactions set status=$1 where transaction_id=$2"
	_, err := m.store.Exec(query, status, txnID)
	if err != nil {
		log.Printf("failed to update transaction due to %s", err)
		return utils.ServerError
	}
	return nil
}

func (m *TransactionsStorage) TransactionReady(transaction *types.TransactionReady) error {
	query := "update transactions set status=$1, customer_id=$2 where transaction_id=$3"
	_, err := m.store.Exec(query, transaction.Status, transaction.CustomerID, transaction.TransactionID)
	if err != nil {
		log.Printf("failed to update transaction due to %s", err)
		return utils.ServerError
	}
	return nil
}

func (m *TransactionsStorage) NewTransaction(newTransaction *types.NewTransaction) error {
	query := "INSERT INTO transactions(transaction_id, status, created_at) VALUES($1, $2, $3)"
	_, err := m.store.Exec(query, newTransaction.TransactionID, newTransaction.Status, newTransaction.CreatedAt)
	if err != nil {
		log.Printf("failed to create transaction due to %s", err)
		return utils.ServerError
	}
	return nil
}

func NewTransactionsStorage(store *sql.DB) *TransactionsStorage {
	return &TransactionsStorage{store: store}
}
