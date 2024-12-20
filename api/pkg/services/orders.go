package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"time"
)

type OrdersStorager interface {
	NewPurchase([]*types.PurchaseRequest) error
	NewOrder(*types.OrderRequest) (uint, error)
	GetOrdersByUserID(uint) ([]*types.OrderList, error)
	GetPurchaseByOrderID(uint) ([]*types.PurchaseList, error)
}

type CartServicer interface {
	HardDeleteByUserID(uint) error
	GetAll(uint) ([]*types.CartList, error)
}

type OrderService struct {
	ordersStorage OrdersStorager
	cartService   CartServicer
}

func (s *OrderService) PlaceOrder(userID uint) error {
	carts, err := s.cartService.GetAll(userID)
	if err != nil {
		return err
	}

	var totalPrice uint64
	for _, cart := range carts {
		totalPrice += uint64(cart.Product.Price)
	}

	orderID, err := s.ordersStorage.NewOrder(&types.OrderRequest{
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

	if err := s.ordersStorage.NewPurchase(purchaseRequests); err != nil {
		return err
	}

	return s.cartService.HardDeleteByUserID(userID)
}

func (s *OrderService) GetOrdersByUserID(id uint) ([]*types.OrderList, error) {
	return s.ordersStorage.GetOrdersByUserID(id)
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

func NewOrderService(ordersStorage OrdersStorager, cartService CartServicer) *OrderService {
	return &OrderService{
		ordersStorage: ordersStorage,
		cartService:   cartService,
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
