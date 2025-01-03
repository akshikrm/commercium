package app

import (
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
)

type CartApi struct {
	cartService types.CartServicer
}

func (c *CartApi) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := ctx.Value("userID")
	carts, err := c.cartService.GetAll(uint32(userID.(int)))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, carts)
}

func (c *CartApi) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cid, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	cart, err := c.cartService.GetOne(uint32(cid))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, cart)
}

func (c *CartApi) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newCart := types.CreateCartRequest{}
	if err := DecodeBody(r.Body, &newCart); err != nil {
		return err
	}
	userID := ctx.Value("userID")
	newCart.UserID = uint32(userID.(int))
	if err := c.cartService.Create(&newCart); err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, "cart created")
}

func (c *CartApi) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cid, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	updatedCart := types.UpdateCartRequest{}
	if err := DecodeBody(r.Body, &updatedCart); err != nil {
		return err
	}
	cart, err := c.cartService.Update(uint32(cid), &updatedCart)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, cart)
}

func (c *CartApi) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	if err := c.cartService.Delete(uint32(id)); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "deleted successfully")
}

func NewCartApi(service types.CartServicer) *CartApi {
	return &CartApi{cartService: service}
}
