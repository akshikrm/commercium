package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
)

type CartModeler interface {
	GetAll(uint32) ([]*types.CartList, error)
	GetOne(uint32) (*types.CartList, error)
	Create(*types.CreateCartRequest) (*types.Cart, error)
	Update(uint32, *types.UpdateCartRequest) (*types.CartList, error)
	Delete(uint32) error
	CheckIfEntryExist(uint32, uint32) (bool, error)
	UpdateQuantity(uint32, uint32, uint) error
	HardDeleteByUserID(string) error
}

type CartService struct {
	cartModel CartModeler
}

func (c *CartService) GetAll(userID uint32) ([]*types.CartList, error) {
	return c.cartModel.GetAll(userID)
}

func (c *CartService) GetOne(cid uint32) (*types.CartList, error) {
	return c.cartModel.GetOne(cid)
}

func (c *CartService) Create(newCart *types.CreateCartRequest) error {
	exists, err := c.cartModel.CheckIfEntryExist(newCart.UserID, newCart.ProductID)
	if err != nil {
		return err
	}

	fmt.Println(exists)
	if exists {
		return c.cartModel.UpdateQuantity(newCart.UserID, newCart.ProductID, newCart.Quantity)
	}

	_, err = c.cartModel.Create(newCart)
	return err
}

func (c *CartService) Update(cid uint32, updateCart *types.UpdateCartRequest) (*types.CartList, error) {
	return c.cartModel.Update(cid, updateCart)
}

func (c *CartService) Delete(cid uint32) error {
	return c.cartModel.Delete(cid)
}

func (c *CartService) HardDeleteByUserID(customerId string) error {
	return c.cartModel.HardDeleteByUserID(customerId)
}

func NewCartService(cartModel CartModeler) *CartService {
	return &CartService{cartModel: cartModel}
}
