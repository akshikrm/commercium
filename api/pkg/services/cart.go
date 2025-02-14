package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
)

type cart struct {
	repository      types.CartRepository
	paymentProvider types.PaymentProvider
}

func (c *cart) GetAll(userID uint32) ([]*types.CartList, error) {
	carts, ok := c.repository.GetAll(userID)

	if !ok {
		return nil, utils.ServerError
	}
	return carts, nil
}

func (c *cart) GetOne(cid uint32) (*types.CartList, error) {
	cart, ok := c.repository.GetOne(cid)
	if !ok {
		return nil, utils.ServerError
	}
	if cart == nil {
		return nil, utils.NotFound
	}
	return cart, nil
}

func (c *cart) Create(payload *types.CreateCartRequest) error {
	exists := c.repository.CheckIfEntryExist(payload.UserID, payload.PriceID)
	if exists {
		if ok := c.repository.UpdateQuantity(payload); !ok {
			return utils.ServerError
		}
		return nil
	}
	if _, ok := c.repository.Create(payload); !ok {
		return utils.ServerError
	}
	return nil
}

func (c *cart) Update(cid uint32, updateCart *types.UpdateCartRequest) (*types.CartList, error) {
	cart, ok := c.repository.Update(cid, updateCart)
	if !ok {
		return nil, utils.ServerError
	}
	if cart == nil {
		return nil, utils.NotFound
	}

	return cart, nil
}

func (c *cart) Delete(cid uint32) error {
	ok := c.repository.Delete(cid)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func (c *cart) HardDeleteByUserID(customerId string) error {
	ok := c.repository.HardDeleteByUserID(customerId)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newCartService(repository types.CartRepository, paymentProvider types.PaymentProvider) *cart {
	return &cart{repository: repository, paymentProvider: paymentProvider}
}
