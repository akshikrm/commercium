package handlers

import (
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
)

type cartHandler struct {
	service types.CartServicer
}

func (c *cartHandler) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := ctx.Value("userID")
	carts, err := c.service.GetAll(uint32(userID.(int)))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, carts)
}

func (c *cartHandler) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cid, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	cart, err := c.service.GetOne(uint32(cid))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, cart)
}

func (c *cartHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newCart := types.CreateCartRequest{}
	if err := DecodeBody(r.Body, &newCart); err != nil {
		return err
	}
	userID := ctx.Value("userID")
	newCart.UserID = uint32(userID.(int))
	if err := c.service.Create(&newCart); err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, "cart created")
}

func (c *cartHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cid, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	updatedCart := types.UpdateCartRequest{}
	if err := DecodeBody(r.Body, &updatedCart); err != nil {
		return err
	}
	cart, err := c.service.Update(uint32(cid), &updatedCart)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, cart)
}

func (c *cartHandler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	if err := c.service.Delete(uint32(id)); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "deleted successfully")
}

func newCart(service types.CartServicer) types.CartHandler {
	handler := new(cartHandler)
	handler.service = service
	return handler
}
