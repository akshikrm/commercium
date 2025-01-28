package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"log"
)

type transaction struct {
	repository      types.TransactionRepository
	orderRepository types.OrdersRepository
	cartService     types.CartServicer
}

func (t *transaction) CreateTransaction(data *types.Data) error {
	transaction := types.NewTransaction{
		TransactionID: data.ID,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt,
		Tax:           data.Details.Totals.Tax,
		SubTotal:      data.Details.Totals.Subtotal,
		GrandTotal:    data.Details.Totals.GrandTotal,
	}
	log.Println("adding transaction")
	id := t.repository.NewTransaction(&transaction)
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
	if ok := t.orderRepository.CreateOrder(orders); !ok {
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

func (t *transaction) ReadyTransaction(data *types.Data) error {
	transaction := types.TransactionReady{
		TransactionID: data.ID,
		Status:        data.Status,
		CustomerID:    data.CustomerID,
	}
	ok := t.repository.TransactionReady(&transaction)
	if !ok {
		return utils.ServerError
	}
	log.Println("transaction ready")
	return nil
}

func (t *transaction) CompleteTransaction(data *types.Data) error {
	transaction := types.TransactionCompleted{
		TransactionID: data.ID,
		Status:        data.Status,
		InvoiceNumber: data.InvoiceNumber,
	}
	ok := t.repository.TransactionCompleted(&transaction)
	if !ok {
		return utils.ServerError
	}
	log.Println("transaction complete")
	return nil
}

func (t *transaction) FailedTransaction(data *types.Data) error {
	ok := t.repository.UpdateStatus(data.ID, "failed")
	if !ok {
		return utils.ServerError
	}
	log.Println("transaction failed")
	return nil
}

func (t *transaction) GetOrderStatus(txnID string) (string, error) {
	status := t.repository.GetOrderStatus(txnID)
	if status == "" {
		return "", utils.NotFound
	}
	return status, nil
}

func (s *transaction) GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, error) {
	purchases, ok := s.orderRepository.GetPurchaseByOrderID(id)
	if !ok {
		return nil, utils.ServerError
	}
	return purchases, nil
}

func (s *transaction) GetOrdersByUserID(id uint32) ([]*types.OrderList, error) {
	orders, ok := s.orderRepository.GetOrdersByUserID(id)
	if !ok {
		return nil, utils.ServerError
	}
	return orders, nil
}

func (s *transaction) GetAllOrders() ([]*types.OrderList, error) {
	orders, ok := s.orderRepository.GetAllOrders()
	if !ok {
		return nil, utils.ServerError
	}
	return orders, nil
}

func (s *transaction) GetShippingInformation() ([]*types.ShippingInformation, error) {
	orders, ok := s.orderRepository.GetShippingInformation()
	if !ok {
		return nil, utils.ServerError
	}
	return orders, nil
}

func (s *transaction) UpdateShippingStatus(orderID uint, status types.ShippingStatus) error {
	ok := s.orderRepository.UpdateOrderStatus(orderID, status)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newPurchaseService(
	repository types.TransactionRepository,
	orderRepository types.OrdersRepository,
	cartService types.CartServicer,
) types.PurchaseService {
	return &transaction{
		repository:      repository,
		orderRepository: orderRepository,
		cartService:     cartService,
	}
}
