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

func (m *TransactionsStorage) TransactionCompleted(transaction *types.TransactionCompleted) error {
	query := "update transactions set status=$1, invoice_number=$2 where transaction_id=$3"
	_, err := m.store.Exec(query, transaction.Status, transaction.InvoiceNumber, transaction.TransactionID)
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

func (m *TransactionsStorage) NewTransaction(newTransaction *types.NewTransaction) *uint {
	query := `INSERT INTO 
			transactions
				(transaction_id, status, created_at, tax, sub_total, grand_total) 
			VALUES
				($1, $2, $3, $4, $5, $6) 
			RETURNING id;
	`
	row := m.store.QueryRow(query,
		newTransaction.TransactionID,
		newTransaction.Status,
		newTransaction.CreatedAt,
		newTransaction.Tax,
		newTransaction.SubTotal,
		newTransaction.GrandTotal,
	)
	var id uint
	err := row.Scan(&id)

	if err != nil {
		log.Printf("failed to create transaction due to %s", err)
		return nil
	}
	return &id
}

func (m *TransactionsStorage) GetOrderStatus(txnId string) string {
	query := "SELECT status from transactions where transaction_id=$1"
	row := m.store.QueryRow(query, txnId)

	var transactionStatus string
	err := row.Scan(&transactionStatus)
	if err != nil {
		log.Printf("query failed %s", err)
		return ""
	}
	return transactionStatus
}

func NewTransactionsStorage(store *sql.DB) *TransactionsStorage {
	return &TransactionsStorage{store: store}
}
