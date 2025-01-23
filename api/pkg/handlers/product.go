package handlers

import (
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
)

type product struct {
	service types.ProductServicer
}

func (u *product) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	filter := r.URL.Query()
	users, err := u.service.Get(filter)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, users)
}

func (u *product) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	foundProduct, err := u.service.GetOne(id)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, foundProduct)
}

func (u *product) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	a := &types.NewProductRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	err := u.service.Create(a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, "product created")
}

func (u *product) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	a := types.NewProductRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	product, err := u.service.Update(id, &a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, product)
}

func (u *product) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	if err := u.service.Delete(id); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "deleted successfully")
}

func newProduct(service types.ProductServicer) types.ProductHandler {
	handler := new(product)
	handler.service = service
	return handler
}
