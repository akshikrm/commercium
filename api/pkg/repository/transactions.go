package repository

import (
	"akshidas/e-com/pkg/types"
	"database/sql"
	"log"
)

type transactions struct {
	store *sql.DB
}

func (m *transactions) UpdateStatus(txnID, status string) bool {
	query := "update transactions set status=$1 where transaction_id=$2"
	_, err := m.store.Exec(query, status, txnID)
	if err != nil {
		log.Printf("failed to update transaction due to %s", err)
		return false
	}
	return true
}

func (m *transactions) TransactionCompleted(transaction *types.TransactionCompleted) bool {
	query := "update transactions set status=$1, invoice_number=$2 where transaction_id=$3"
	_, err := m.store.Exec(query, transaction.Status, transaction.InvoiceNumber, transaction.TransactionID)
	if err != nil {
		log.Printf("failed to update transaction due to %s", err)
		return false
	}
	return true
}

func (m *transactions) TransactionReady(transaction *types.TransactionReady) bool {
	query := "update transactions set status=$1, customer_id=$2 where transaction_id=$3"
	_, err := m.store.Exec(query, transaction.Status, transaction.CustomerID, transaction.TransactionID)
	if err != nil {
		log.Printf("failed to update transaction due to %s", err)
		return false
	}
	return true
}

func (m *transactions) NewTransaction(newTransaction *types.NewTransaction) *uint32 {
	query := `INSERT INTO 
			transactions
				(transaction_id, status, created_at, tax, sub_total, grand_total, customer_id) 
			VALUES
				($1, $2, $3, $4, $5, $6, $7) 
			RETURNING id;
	`
	row := m.store.QueryRow(query,
		newTransaction.TransactionID,
		newTransaction.Status,
		newTransaction.CreatedAt,
		newTransaction.Tax,
		newTransaction.SubTotal,
		newTransaction.GrandTotal,
		newTransaction.CustomerID,
	)
	var id uint32
	err := row.Scan(&id)

	if err != nil {
		log.Printf("failed to create transaction due to %s", err)
		return nil
	}
	return &id
}

func (m *transactions) GetOrderStatus(txnId string) string {
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

func newTransactions(store *sql.DB) types.TransactionRepository {
	return &transactions{store: store}
}
