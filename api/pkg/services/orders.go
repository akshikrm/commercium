package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"time"
)

type PurchaseStorager interface {
	// GetAllByUserID(uint32) ([]*types.PurchaseList, error)
	// GetAll() ([]*types.PurchaseList, error)
	NewPurchase([]*types.PurchaseRequest) error
	NewOrder(*types.OrderRequest) (uint, error)
	// GetByOrderID(string) (*types.OrderView, error)
}

type CartServicer interface {
	HardDeleteByUserID(uint) error
	GetAll(uint) ([]*types.CartList, error)
}

type PurchaseService struct {
	purchaseStorage PurchaseStorager
	cartService     CartServicer
}

func (s *PurchaseService) PlaceOrder(userID uint) error {
	carts, err := s.cartService.GetAll(userID)
	if err != nil {
		return err
	}

	var totalPrice uint64
	for _, cart := range carts {
		totalPrice += uint64(cart.Product.Price)
	}

	orderID, err := s.purchaseStorage.NewOrder(&types.OrderRequest{
		UserID:  userID,
		OrderID: genOrderID(),
		Price:   totalPrice,
	})

	purchaseRequests := []*types.PurchaseRequest{}
	for _, cart := range carts {
		purchaseRequests = append(purchaseRequests, &types.PurchaseRequest{
			ProductID: cart.Product.ID,
			Price:     cart.Product.Price,
			Quantity:  cart.Quantity,
			OrderID:   orderID,
		})
	}

	if err := s.purchaseStorage.NewPurchase(purchaseRequests); err != nil {
		return err
	}

	return s.cartService.HardDeleteByUserID(userID)
}

// func (s *PurchaseService) GetAll(userID uint32, role string) ([]*types.PurchaseList, error) {
// 	if role == "admin" {
// 		return s.purchaseStorage.GetAll()
// 	}

// 	return s.purchaseStorage.GetAllByUserID(userID)
// }

// func (s *PurchaseService) GetByOrderID(id string) (*types.OrderView, error) {
// 	return s.purchaseStorage.GetByOrderID(id)
// }

func NewPurchaseService(purchaseStorage PurchaseStorager, cartService CartServicer) *PurchaseService {
	return &PurchaseService{
		purchaseStorage: purchaseStorage,
		cartService:     cartService,
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
