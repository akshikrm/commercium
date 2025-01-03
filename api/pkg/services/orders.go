package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"time"
)

type order struct {
	repository types.OrdersRepository
}

func (s *order) GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, error) {
	return s.repository.GetPurchaseByOrderID(id)
}

func (s *order) GetOrdersByUserID(id uint) ([]*types.OrderList, error) {
	orders, err := s.repository.GetOrdersByUserID(id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func newOrderService(repository types.OrdersRepository) *order {
	return &order{
		repository: repository,
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
