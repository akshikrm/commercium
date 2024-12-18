package services

import "akshidas/e-com/pkg/types"

type OrderStorager interface {
	Create(*types.OrderRequest) (string, error)
}

type OrderService struct {
	storage OrderStorager
}

func (o *OrderService) Create(newOrder *types.OrderRequest) (string, error) {
	return o.storage.Create(newOrder)
}

func NewOrderService(storage OrderStorager) *OrderService {
	return &OrderService{storage: storage}
}
