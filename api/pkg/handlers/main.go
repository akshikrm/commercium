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
	Upload          types.UploadHandler
}

func New(service *services.Service) *Handler {
	handler := new(Handler)
	handler.Product = newProduct(service.Product)
	handler.ProductCategory = newProductCategory(service.ProductCategory)
	handler.Cart = newCart(service.Cart)
	handler.User = newUser(service.User, service.PaymentProvider)
	handler.Purchase = newPurchase(service.Purchase, service.Cart, service.PaymentProvider)
	handler.Upload = newUpload()

	handler.Middleware = newMiddleWare(service.User)
	return handler
}
