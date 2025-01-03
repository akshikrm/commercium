package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
)

type cart struct {
	repository types.CartRepository
}

func (c *cart) GetAll(userID uint32) ([]*types.CartList, error) {
	return c.repository.GetAll(userID)
}

func (c *cart) GetOne(cid uint32) (*types.CartList, error) {
	return c.repository.GetOne(cid)
}

func (c *cart) Create(newCart *types.CreateCartRequest) error {
	exists, err := c.repository.CheckIfEntryExist(newCart.UserID, newCart.ProductID)
	if err != nil {
		return err
	}

	fmt.Println(exists)
	if exists {
		return c.repository.UpdateQuantity(newCart.UserID, newCart.ProductID, newCart.Quantity)
	}

	_, err = c.repository.Create(newCart)
	return err
}

func (c *cart) Update(cid uint32, updateCart *types.UpdateCartRequest) (*types.CartList, error) {
	return c.repository.Update(cid, updateCart)
}

func (c *cart) Delete(cid uint32) error {
	return c.repository.Delete(cid)
}

func (c *cart) HardDeleteByUserID(customerId string) error {
	return c.repository.HardDeleteByUserID(customerId)
}

func newCartService(repository types.CartRepository) *cart {
	return &cart{repository: repository}
}
