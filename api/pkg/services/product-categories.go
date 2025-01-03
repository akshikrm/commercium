package services

import (
	"akshidas/e-com/pkg/types"
	"net/url"
)

type productCategory struct {
	repository types.ProductCategoriesRepository
}

func (p *productCategory) Create(newCategory *types.NewProductCategoryRequest) (*types.ProductCategory, error) {
	return p.repository.Create(newCategory)
}

func (p *productCategory) GetNames() ([]*types.ProductCategoryName, error) {
	return p.repository.GetNames()
}
func (p *productCategory) GetAll(filter url.Values) ([]*types.ProductCategory, error) {
	return p.repository.GetAll(filter)
}

func (p *productCategory) GetOne(id int) (*types.ProductCategory, error) {
	return p.repository.GetOne(id)
}

func (p *productCategory) Update(id int, updateProductCategory *types.UpdateProductCategoryRequest) (*types.ProductCategory, error) {
	return p.repository.Update(id, updateProductCategory)
}

func (p *productCategory) Delete(id int) error {
	return p.repository.Delete(id)
}

func newProductCategoryService(repository types.ProductCategoriesRepository) *productCategory {
	return &productCategory{
		repository: repository,
	}
}
