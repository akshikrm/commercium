package app

import (
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
)

type ProductCategoriesApi struct {
	service types.ProductCateogriesServicer
}

func (s *ProductCategoriesApi) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newProductCategory := types.NewProductCategoryRequest{}
	if err := DecodeBody(r.Body, &newProductCategory); err != nil {
		return err
	}
	_, err := s.service.Create(&newProductCategory)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, "product category created")
}

func (s *ProductCategoriesApi) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	filter := r.URL.Query()
	filterType := filter.Get("type")
	if filterType == "name" {
		productCategoryNames, err := s.service.GetNames()
		if err != nil {
			return err
		}
		return writeJson(w, http.StatusOK, productCategoryNames)
	}
	productCategories, err := s.service.GetAll(filter)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, productCategories)
}

func (s *ProductCategoriesApi) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	productCategories, err := s.service.GetOne(id)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, productCategories)
}

func (s *ProductCategoriesApi) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	updateProductCategory := types.UpdateProductCategoryRequest{}
	if err := DecodeBody(r.Body, &updateProductCategory); err != nil {
		return err
	}

	updatedProductCategory, err := s.service.Update(id, &updateProductCategory)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, updatedProductCategory)
}

func (s *ProductCategoriesApi) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}

	if err := s.service.Delete(id); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "delete successfully")
}

func newProductCategoriesApi(service types.ProductCateogriesServicer) *ProductCategoriesApi {
	return &ProductCategoriesApi{service: service}
}
