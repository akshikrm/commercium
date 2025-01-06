package services

import (
	"akshidas/e-com/pkg/types"
	"log"
	"net/url"
)

type product struct {
	repository types.ProductRepository
}

func (r *product) Get(filter url.Values) ([]*types.ProductsList, error) {
	return r.repository.GetAll(filter)
}

func (r *product) Create(newProduct *types.NewProductRequest) error {
	paddlePayment := new(PaddlePayment)
	if err := paddlePayment.Init(); err != nil {
		return err
	}

	if err := paddlePayment.CreateProduct(newProduct); err != nil {
		return err
	}

	product, err := r.repository.Create(newProduct)
	log.Printf("%s Product Added", product.Name)
	if err != nil {
		return (err)
	}
	return err
}

func (r *product) Update(id int, newProduct *types.NewProductRequest) (*types.OneProduct, error) {
	return r.repository.Update(id, newProduct)
}

func (r *product) GetOne(id int) (*types.OneProduct, error) {
	return r.repository.GetOne(id)
}

func (r *product) Delete(id int) error {
	return r.repository.Delete(id)
}

func newProductService(repository types.ProductRepository) *product {
	return &product{
		repository: repository,
	}
}
