package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"log"
)

type TransactionStorager interface {
	NewTransaction(*types.NewTransaction) *uint
	TransactionReady(*types.TransactionReady) error
	UpdateStatus(string, string) error
	TransactionCompleted(*types.TransactionCompleted) error
}

type OrderStorager interface {
	CreateOrder([]*types.NewOrder) error
}

type TransactionService struct {
	store TransactionStorager
	order OrderStorager
}

func (t *TransactionService) CreateTransaction(data *types.Data) error {
	transaction := types.NewTransaction{
		TransactionID: data.ID,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt,
		Tax:           data.Details.Totals.Tax,
		SubTotal:      data.Details.Totals.Subtotal,
		GrandTotal:    data.Details.Totals.GrandTotal,
	}
	id := t.store.NewTransaction(&transaction)
	if id == nil {
		log.Println("transaction create failed")
		return utils.ServerError
	}

	orders := []*types.NewOrder{}
	for _, item := range data.Items {
		order := types.NewOrder{
			TransactionID: *id,
			PriceID:       item.Price.ID,
			ProductID:     item.Price.ProductID,
			Quantity:      item.Quantity,
			Amount:        item.Price.UnitPrice.Amount,
		}
		orders = append(orders, &order)
	}

	log.Println("adding orders")
	if err := t.order.CreateOrder(orders); err != nil {
		log.Println(err)
		return utils.ServerError
	}

	log.Println("transaction created")
	return nil

}

func (t *TransactionService) ReadyTransaction(data *types.Data) error {
	transaction := types.TransactionReady{
		TransactionID: data.ID,
		Status:        data.Status,
		CustomerID:    data.CustomerID,
	}
	err := t.store.TransactionReady(&transaction)
	if err == nil {
		log.Println("transaction ready")
	}
	return err
}

func (t *TransactionService) CompleteTransaction(data *types.Data) error {
	transaction := types.TransactionCompleted{
		TransactionID: data.ID,
		Status:        data.Status,
		InvoiceNumber: data.InvoiceNumber,
	}
	err := t.store.TransactionCompleted(&transaction)
	if err == nil {
		log.Println("transaction complete")
	}
	return err
}

func (t *TransactionService) FailedTransaction(data *types.Data) error {
	log.Println("transaction failed")
	return t.store.UpdateStatus(data.ID, "failed")
}

func NewTransactionService(storage TransactionStorager, order OrderStorager) *TransactionService {
	return &TransactionService{
		store: storage,
		order: order,
	}
}
