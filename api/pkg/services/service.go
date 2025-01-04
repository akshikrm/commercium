package services

import (
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/types"
)

type Service struct {
	User            types.UserServicer
	Product         types.ProductServicer
	ProductCategory types.ProductCateogoryServicer
	Cart            types.CartServicer
	Purchase        types.PurchaseServicer
	Transaction     types.TransactionServicer
	Role            types.RoleService
}

func New(store *repository.Storage) *Service {
	service := new(Service)
	service.Product = newProductService(store.Product)
	service.ProductCategory = newProductCategoryService(store.ProductCategory)
	service.Cart = newCartService(store.Cart)
	service.User = newUserService(store.User, store.Profile)
	service.Purchase = newOrderService(store.Orders)
	service.Transaction = newTransactionService(store.Transaction, store.Orders, service.Cart)
	service.Role = newRoleService(store.Role)
	return service
}
