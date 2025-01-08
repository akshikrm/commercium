package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"net/url"
)

type productCategory struct {
	repository types.ProductCategoriesRepository
}

func (p *productCategory) Create(newCategory *types.NewProductCategoryRequest) (*types.ProductCategory, error) {
	productCategory, ok := p.repository.Create(newCategory)
	if !ok {
		return nil, utils.ServerError

	}
	return productCategory, nil
}

func (p *productCategory) GetNames() ([]*types.ProductCategoryName, error) {
	names, ok := p.repository.GetNames()
	if !ok {
		return nil, utils.ServerError
	}
	return names, nil
}
func (p *productCategory) GetAll(filter url.Values) ([]*types.ProductCategory, error) {
	productCategories, ok := p.repository.GetAll(filter)
	if !ok {
		return nil, utils.ServerError
	}
	return productCategories, nil
}

func (p *productCategory) GetOne(id int) (*types.ProductCategory, error) {
	productCategory, ok := p.repository.GetOne(id)
	if !ok {
		return nil, utils.ServerError
	}
	return productCategory, nil
}

func (p *productCategory) Update(id int, updateProductCategory *types.UpdateProductCategoryRequest) (*types.ProductCategory, error) {
	updatedProductCategory, ok := p.repository.Update(id, updateProductCategory)
	if !ok {
		return nil, utils.ServerError
	}
	return updatedProductCategory, nil
}

func (p *productCategory) Delete(id int) error {
	ok := p.repository.Delete(id)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newProductCategoryService(repository types.ProductCategoriesRepository) *productCategory {
	return &productCategory{
		repository: repository,
	}
}
