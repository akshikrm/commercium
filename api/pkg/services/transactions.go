package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"log"
)

type TransactionService struct {
	transactionRepository types.TransactionRepository
	orderRepository       types.OrdersRepository
	cartService           types.CartServicer
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

	log.Println("adding transaction")
	id := t.transactionRepository.NewTransaction(&transaction)
	if id == nil {
		log.Println("transaction create failed")
		return utils.ServerError
	}
	log.Println("transaction added")

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
	if err := t.orderRepository.CreateOrder(orders); err != nil {
		log.Println(err)
		return utils.ServerError
	}
	log.Println("orders added")

	log.Println("deleting cart")
	if err := t.cartService.HardDeleteByUserID(data.CustomerID); err != nil {
		return utils.ServerError
	}
	log.Println("cart deleted")

	log.Println("transaction created")
	return nil

}

func (t *TransactionService) ReadyTransaction(data *types.Data) error {
	transaction := types.TransactionReady{
		TransactionID: data.ID,
		Status:        data.Status,
		CustomerID:    data.CustomerID,
	}
	err := t.transactionRepository.TransactionReady(&transaction)
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
	err := t.transactionRepository.TransactionCompleted(&transaction)
	if err == nil {
		log.Println("transaction complete")
	}
	return err
}

func (t *TransactionService) FailedTransaction(data *types.Data) error {
	log.Println("transaction failed")
	return t.transactionRepository.UpdateStatus(data.ID, "failed")
}

func (t *TransactionService) GetOrderStatus(txnID string) (string, error) {
	status := t.transactionRepository.GetOrderStatus(txnID)
	if status == "" {
		return "", utils.NotFound
	}
	return status, nil
}

func newTransactionService(
	transactionRepository types.TransactionRepository,
	orderRepoistory types.OrdersRepository,
	cartService types.CartServicer,
) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		orderRepository:       orderRepoistory,
		cartService:           cartService,
	}
}
