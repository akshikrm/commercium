package services

import (
	"akshidas/e-com/pkg/types"
)

type PurchaseStorager interface {
	GetByUserID(uint32) ([]*types.Purchase, error)
	Create([]*types.PurchaseRequest, uint) error
}

type CartServicer interface {
	HardDeleteByUserID(uint) error
	GetAll(uint) ([]*types.CartList, error)
}

type PurchaseService struct {
	purchaseStorage PurchaseStorager
	cartService     CartServicer
}

func (s *PurchaseService) Create(newPurchase *types.PurchaseRequest) error {
	carts, err := s.cartService.GetAll(newPurchase.UserID)

	if err != nil {
		return err
	}

	var totalPrice uint
	var productIds []uint
	for _, cart := range carts {
		totalPrice += cart.Product.Price
		productIds = append(productIds, cart.Product.ID)
	}

	newPurchaseEntry := []*types.PurchaseRequest{}
	for _, productId := range productIds {
		newPurchaseEntry = append(newPurchaseEntry, &types.PurchaseRequest{
			UserID:    newPurchase.UserID,
			ProductID: uint(productId),
		})
	}
	if err := s.purchaseStorage.Create(newPurchaseEntry, totalPrice); err != nil {
		return err
	}

	return s.cartService.HardDeleteByUserID(newPurchase.UserID)
}

func (s *PurchaseService) GetByUserID(userID uint32) ([]*types.Purchase, error) {
	return s.purchaseStorage.GetByUserID(userID)
}

func NewPurchaseService(purchaseStorage PurchaseStorager, cartService CartServicer) *PurchaseService {
	return &PurchaseService{purchaseStorage: purchaseStorage, cartService: cartService}
}
