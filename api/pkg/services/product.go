package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"net/url"
)

type product struct {
	repository      types.ProductRepository
	paymentProvider types.PaymentProvider
}

func (r *product) Get(filter url.Values) ([]*types.ProductsList, error) {
	products, ok := r.repository.GetAll(filter)
	if !ok {
		return nil, utils.ServerError
	}
	return products, nil
}

func (r *product) Create(newProduct *types.NewProductRequest) error {
	if err := r.paymentProvider.CreateProduct(newProduct); err != nil {
		return err
	}

	savedProduct, ok := r.repository.Create(newProduct)
	if !ok {
		return utils.ServerError
	}

	if savedProduct.Type == types.OneTimeProduct {
		if price := r.paymentProvider.CreatePrice(
			savedProduct.ProductID,
			"one time price",
			savedProduct.Price,
		); price != nil {
			if ok := r.repository.CreatePrice(price); ok {
				return nil
			}
		}
		return utils.ServerError
	}

	for _, priceItem := range newProduct.SubscriptionPrice {
		if price := r.paymentProvider.CreatePrice(
			savedProduct.ProductID,
			priceItem.Label,
			priceItem.Value,
		); price != nil {
			if ok := r.repository.CreatePrice(price); !ok {
				return utils.ServerError
			}
		}
	}
	return nil
}

func (r *product) Update(id int, newProduct *types.NewProductRequest) (*types.OneProduct, error) {
	updatedProduct, ok := r.repository.Update(id, newProduct)
	if !ok {
		return nil, utils.ServerError
	}
	return updatedProduct, nil
}

func (r *product) GetOne(id int) (*types.OneProduct, error) {
	product, ok := r.repository.GetOne(id)
	if !ok {
		return nil, utils.ServerError
	}
	return product, nil
}

func (r *product) Delete(id int) error {
	ok := r.repository.Delete(id)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newProductService(repository types.ProductRepository, paymentProvider types.PaymentProvider) *product {
	return &product{
		repository:      repository,
		paymentProvider: paymentProvider,
	}
}
