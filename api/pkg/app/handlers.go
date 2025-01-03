package app

import (
	"context"
)

type routesFunc map[string]apiFunc

type HandleFunc func(*Server) routesFunc

func UserHandler(s *Server) routesFunc {
	ctx := context.Background()
	userApi := newUserApi(s.services.User)
	return routesFunc{
		"POST /users":               userApi.Create,
		"POST /login":               userApi.Login,
		"GET /users/my-customer-id": s.middleware.IsAuthenticated(ctx, userApi.GetCustomerID),
		"GET /profile":              s.middleware.IsAuthenticated(ctx, userApi.GetProfile),
		"PUT /profile":              s.middleware.IsAuthenticated(ctx, userApi.UpdateProfile),
		"GET /users":                s.middleware.IsAdmin(ctx, userApi.GetAll),
		"GET /users/{id}":           s.middleware.IsAdmin(ctx, userApi.GetOne),
		"DELETE /users/{id}":        s.middleware.IsAdmin(ctx, userApi.Delete),
	}
}

func ProductHandler(s *Server) routesFunc {
	ctx := context.Background()
	productApi := newProductApi(s.services.Product)

	return routesFunc{
		"GET /products":         s.middleware.IsAuthenticated(ctx, productApi.GetAll),
		"POST /products":        s.middleware.IsAdmin(ctx, productApi.Create),
		"GET /products/{id}":    s.middleware.IsAdmin(ctx, productApi.GetOne),
		"PUT /products/{id}":    s.middleware.IsAdmin(ctx, productApi.Update),
		"DELETE /products/{id}": s.middleware.IsAdmin(ctx, productApi.Delete),
	}
}

func ProductCategoryHandler(s *Server) routesFunc {
	ctx := context.Background()
	productCategoryApi := newProductCategoriesApi(s.services.ProductCategory)

	return routesFunc{
		"POST /products/categories":        s.middleware.IsAdmin(ctx, productCategoryApi.Create),
		"GET /products/categories":         s.middleware.IsAdmin(ctx, productCategoryApi.GetAll),
		"GET /products/categories/{id}":    s.middleware.IsAdmin(ctx, productCategoryApi.GetOne),
		"PUT /products/categories/{id}":    s.middleware.IsAdmin(ctx, productCategoryApi.Update),
		"DELETE /products/categories/{id}": s.middleware.IsAdmin(ctx, productCategoryApi.Delete),
	}
}

func CartHandler(s *Server) routesFunc {
	ctx := context.Background()
	cartApi := newCartApi(s.services.Cart)

	return routesFunc{
		"POST /carts":        s.middleware.IsAuthenticated(ctx, cartApi.Create),
		"GET /carts":         s.middleware.IsAuthenticated(ctx, cartApi.GetAll),
		"PUT /carts/{id}":    s.middleware.IsAuthenticated(ctx, cartApi.Update),
		"DELETE /carts/{id}": s.middleware.IsAuthenticated(ctx, cartApi.Delete),
	}
}

func PurchaseHandler(s *Server) routesFunc {
	ctx := context.Background()
	purchaseApi := newOrdersApi(s.services.Transaction, s.services.Purchase)

	return routesFunc{
		"POST /transactions":          purchaseApi.HandleTransactionHook,
		"GET /orders/invoice/{txnId}": s.middleware.IsAuthenticated(ctx, purchaseApi.GetInvoice),
		"GET /orders/status/{txnId}":  s.middleware.IsAuthenticated(ctx, purchaseApi.GetOrderStatus),
		"GET /orders":                 s.middleware.IsAuthenticated(ctx, purchaseApi.GetMyOrders),
	}
}
