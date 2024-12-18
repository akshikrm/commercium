package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"time"
)

type PurchaseStorager interface {
	GetAllByUserID(uint32) ([]*types.PurchaseList, error)
	GetAll() ([]*types.PurchaseList, error)
	Create([]*types.PurchaseRequest, uint64) error
	GetByOrderID(string) (*types.OrderView, error)
}

type CartServicer interface {
	HardDeleteByUserID(uint) error
	GetAll(uint) ([]*types.CartList, error)
}

type OrderServicer interface {
	Create(*types.OrderRequest) (uint, error)
}

type PurchaseService struct {
	purchaseStorage PurchaseStorager
	orderService    OrderServicer
	cartService     CartServicer
}

func (s *PurchaseService) Create(newPurchase *types.PurchaseRequest) error {
	carts, err := s.cartService.GetAll(newPurchase.UserID)
	if err != nil {
		return err
	}

	var totalPrice uint64
	var productIds []uint
	for _, cart := range carts {
		totalPrice += uint64(cart.Product.Price)
		productIds = append(productIds, cart.Product.ID)
	}

	newPurchaseEntry := []*types.PurchaseRequest{}
	orderID, err := s.orderService.Create(&types.OrderRequest{
		UserID:  newPurchase.UserID,
		OrderID: genOrderID(),
		Price:   uint64(totalPrice),
	})
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
	return s.purchaseStorage.GetAllByUserID(userID)
}

func (s *PurchaseService) GetByOrderID(id string) (*types.OrderView, error) {
	return s.purchaseStorage.GetByOrderID(id)
}

func NewPurchaseService(purchaseStorage PurchaseStorager, cartService CartServicer, orderService OrderServicer) *PurchaseService {
	return &PurchaseService{
		purchaseStorage: purchaseStorage,
		cartService:     cartService,
		orderService:    orderService,
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
