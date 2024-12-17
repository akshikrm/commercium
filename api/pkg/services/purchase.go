package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"time"
)

type PurchaseStorager interface {
	GetByUserID(uint32) ([]*types.PurchaseList, error)
	GetAll() ([]*types.PurchaseList, error)
	Create([]*types.PurchaseRequest, uint) error
	GetByID(uint32) (*types.PurchaseList, error)
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
	orderID := genOrderID()
	for _, productId := range productIds {
		newPurchaseEntry = append(newPurchaseEntry, &types.PurchaseRequest{
			UserID:    newPurchase.UserID,
			ProductID: uint(productId),
			OrderID:   orderID,
		})
	}

	if err := s.purchaseStorage.Create(newPurchaseEntry, totalPrice); err != nil {
		return err
	}

	return s.cartService.HardDeleteByUserID(newPurchase.UserID)
}

func (s *PurchaseService) GetAll(userID uint32, role string) ([]*types.PurchaseList, error) {
	if role == "admin" {
		return s.purchaseStorage.GetAll()
	}
	return s.purchaseStorage.GetByUserID(userID)
}

func (s *PurchaseService) GetByID(id uint32) (*types.PurchaseList, error) {
	return s.purchaseStorage.GetByID(id)
}

func NewPurchaseService(purchaseStorage PurchaseStorager, cartService CartServicer) *PurchaseService {
	return &PurchaseService{purchaseStorage: purchaseStorage, cartService: cartService}
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
