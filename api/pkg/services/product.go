package services

import (
	"akshidas/e-com/pkg/types"
	"fmt"
	"net/url"
)

type ProductStorager interface {
	GetAll(url.Values) ([]*types.ProductsList, error)
	GetOne(int) (*types.Product, error)
	Create(*types.CreateNewProduct) (*types.Product, error)
	Update(int, *types.CreateNewProduct) (*types.Product, error)
	Delete(int) error
}

type ProductService struct {
	productModel ProductStorager
}

func (r *ProductService) Get(filter url.Values) ([]*types.ProductsList, error) {
	return r.productModel.GetAll(filter)
}

func (r *ProductService) Create(newProduct *types.CreateNewProduct) error {
	paddlePayment := new(PaddlePayment)
	if err := paddlePayment.Init(); err != nil {
		return err
	}

	if err := paddlePayment.CreateProduct(newProduct); err != nil {
		return err
	}

	product, err := r.productModel.Create(newProduct)
	fmt.Println(product.Name)
	if err != nil {
		return (err)
	}
	return err
}

func (r *ProductService) Update(id int, newProduct *types.CreateNewProduct) (*types.Product, error) {
	return r.productModel.Update(id, newProduct)
}

func (r *ProductService) GetOne(id int) (*types.Product, error) {
	return r.productModel.GetOne(id)
}

func (r *ProductService) Delete(id int) error {
	return r.productModel.Delete(id)
}

func NewProductService(productModel ProductStorager) *ProductService {
	return &ProductService{
		productModel: productModel,
	}
}
