package handlers

import (
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
)

type Handler struct {
	Middleware      *MiddleWares
	User            types.UserHandler
	Product         types.ProductHandler
	ProductCategory types.ProductCategoryHandler
	Cart            types.CartHandler
	Purchase        types.PurchaseHandler
}

func New(service *services.Service) *Handler {
	handler := new(Handler)
	handler.Product = NewProduct(service.Product)
	handler.ProductCategory = NewProductCategory(service.ProductCategory)
	handler.Cart = NewCart(service.Cart)
	handler.User = NewUser(service.User)
	handler.Purchase = NewOrders(service.Transaction, service.Purchase)
	handler.Middleware = newMiddleWare(service.User)
	return handler
}
