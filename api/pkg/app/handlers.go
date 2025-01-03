package app

import (
	"context"
)

func (s *Server) registerUserRoutes() {
	ctx := context.Background()
	userApi := newUserApi(s.services.User)
	middlware := newMiddleWare(s.services.User)

	s.router.HandleFunc("POST /users", RouteHandler(userApi.Create))
	s.router.HandleFunc("POST /login", RouteHandler(userApi.Login))
	s.router.HandleFunc("GET /users/my-customer-id", RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetCustomerID)))
	s.router.HandleFunc("GET /profile", RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetProfile)))
	s.router.HandleFunc("PUT /profile", RouteHandler(middlware.IsAuthenticated(ctx, userApi.UpdateProfile)))
	s.router.HandleFunc("GET /users", RouteHandler(middlware.IsAdmin(ctx, userApi.GetAll)))
	s.router.HandleFunc("GET /users/{id}", RouteHandler(middlware.IsAdmin(ctx, userApi.GetOne)))
	s.router.HandleFunc("DELETE /users/{id}", RouteHandler(middlware.IsAdmin(ctx, userApi.Delete)))
}

func (s *Server) registerProductRoutes() {
	ctx := context.Background()
	productApi := newProductApi(s.services.Product)
	middlware := newMiddleWare(s.services.User)

	s.router.HandleFunc("GET /products", RouteHandler(middlware.IsAuthenticated(ctx, productApi.GetAll)))
	s.router.HandleFunc("POST /products", RouteHandler(middlware.IsAdmin(ctx, productApi.Create)))
	s.router.HandleFunc("GET /products/{id}", RouteHandler(middlware.IsAdmin(ctx, productApi.GetOne)))
	s.router.HandleFunc("PUT /products/{id}", RouteHandler(middlware.IsAdmin(ctx, productApi.Update)))
	s.router.HandleFunc("DELETE /products/{id}", RouteHandler(middlware.IsAdmin(ctx, productApi.Delete)))
}

func (s *Server) registerProductCategoryRoutes() {
	ctx := context.Background()
	productCategoryApi := newProductCategoriesApi(s.services.ProductCategory)
	middlware := newMiddleWare(s.services.User)

	s.router.HandleFunc("POST /products/categories", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Create)))
	s.router.HandleFunc("GET /products/categories", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.GetAll)))
	s.router.HandleFunc("GET /products/categories/{id}", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.GetOne)))
	s.router.HandleFunc("PUT /products/categories/{id}", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Update)))
	s.router.HandleFunc("DELETE /products/categories/{id}", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Delete)))
}

func (s *Server) registerCartRoutes() {
	ctx := context.Background()
	cartApi := newCartApi(s.services.Cart)
	middlware := newMiddleWare(s.services.User)

	s.router.HandleFunc("POST /carts", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Create)))
	s.router.HandleFunc("GET /carts", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.GetAll)))
	s.router.HandleFunc("PUT /carts/{id}", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Update)))
	s.router.HandleFunc("DELETE /carts/{id}", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Delete)))
}

func (s *Server) registerPurchaseRoutes() {
	ctx := context.Background()
	purchaseApi := newOrdersApi(s.services.Transaction, s.services.Purchase)
	middlware := newMiddleWare(s.services.User)

	s.router.HandleFunc("POST /transactions", RouteHandler(purchaseApi.HandleTransactionHook))
	s.router.HandleFunc("GET /orders/invoice/{txnId}", RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetInvoice)))
	s.router.HandleFunc("GET /orders/status/{txnId}", RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetOrderStatus)))
	s.router.HandleFunc("GET /orders", RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetMyOrders)))
}
