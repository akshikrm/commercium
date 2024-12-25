package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"time"
)

type OrdersStorager interface {
	GetOrdersByUserID(uint) ([]*types.OrderList, error)
	GetPurchaseByOrderID(uint) ([]*types.PurchaseList, error)
}

type OrderService struct {
	ordersStorage OrdersStorager
}

func (s *OrderService) GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, error) {
	return s.ordersStorage.GetPurchaseByOrderID(id)
}

func (s *OrderService) GetOrdersByUserID(id uint) ([]*types.OrderList, error) {
	orders, err := s.ordersStorage.GetOrdersByUserID(id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func NewOrderService(ordersStorage OrdersStorager) *OrderService {
	return &OrderService{
		ordersStorage: ordersStorage,
	}
}

func genOrderID() string {
	date := time.Now()
	month := date.Month().String()[0:1]
	weekday := date.Weekday().String()[0:1]
	second := date.Second()
	hour := date.Hour()
	day := date.Day()

	return fmt.Sprintf("W%s%s%d%d%d", month, weekday, day, hour, second)
}
