package services

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/types"
)

type Service struct {
	User            types.UserServicer
	Product         types.ProductServicer
	ProductCategory types.ProductCateogoryServicer
	Cart            types.CartServicer
	Purchase        types.PurchaseService
	Role            types.RoleService
	PaymentProvider types.PaymentProvider
}

func New(store *repository.Storage, config *config.Config) *Service {
	service := new(Service)
	service.PaymentProvider = NewPaddlePayment(config)
	service.Product = newProductService(store.Product, service.PaymentProvider)
	service.ProductCategory = newProductCategoryService(store.ProductCategory)
	service.Cart = newCartService(store.Cart, service.PaymentProvider)
	service.User = newUserService(store.User, store.Profile)
	service.Purchase = newPurchaseService(store.Transaction, store.Orders, service.Cart, service.User, service.PaymentProvider)
	service.Role = newRoleService(store.Role)
	return service
}
