package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"net/url"
)

type product struct {
	repository types.ProductRepository
}

func (r *product) Get(filter url.Values) ([]*types.ProductsList, error) {
	products, ok := r.repository.GetAll(filter)
	if !ok {
		return nil, utils.ServerError
	}
	return products, nil
}

func (r *product) Create(newProduct *types.NewProductRequest) error {
	paddlePayment := NewPaddlePayment()
	if err := paddlePayment.Init(); err != nil {
		return err
	}

	if err := paddlePayment.CreateProduct(newProduct); err != nil {
		return err
	}

	_, ok := r.repository.Create(newProduct)
	if !ok {
		return utils.ServerError
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

func newProductService(repository types.ProductRepository) *product {
	return &product{
		repository: repository,
	}
}
