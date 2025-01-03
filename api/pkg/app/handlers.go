package app

import (
	"akshidas/e-com/pkg/services"
	"context"
	"net/http"
)

func (s *Server) registerRoutes(services *services.Service) {
	ctx := context.Background()
	r := s.router

	r.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		Cors(w)
	})

	// Api
	userApi := NewUserApi(services.User)
	productApi := NewProductApi(services.Product)
	cartApi := NewCartApi(services.Cart)
	productCategoryApi := NewProductCategoriesApi(services.ProductCategory)
	// uploadApi := NewUploadApi(s.Store)
	purchaseApi := NewOrdersApi(services.Transaction, services.Purchase)

	// Middle wares
	middlware := NewMiddleWare(userApi.UserService)

	// Public Routes
	r.HandleFunc("POST /transactions", RouteHandler(purchaseApi.HandleTransactionHook))
	r.HandleFunc("POST /users", RouteHandler(userApi.Create))
	r.HandleFunc("POST /login", RouteHandler(userApi.Login))
	// r.HandleFunc("POST /upload", RouteHandler(uploadApi.Upload))

	// Authenticated Routes
	r.HandleFunc("GET /users/my-customer-id", RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetCustomerID)))
	r.HandleFunc("GET /orders/invoice/{txnId}", RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetInvoice)))
	r.HandleFunc("GET /orders/status/{txnId}", RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetOrderStatus)))
	r.HandleFunc("GET /orders", RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetMyOrders)))

	r.HandleFunc("GET /profile", RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetProfile)))
	r.HandleFunc("PUT /profile", RouteHandler(middlware.IsAuthenticated(ctx, userApi.UpdateProfile)))
	r.HandleFunc("POST /carts", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Create)))
	r.HandleFunc("GET /carts", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.GetAll)))
	r.HandleFunc("PUT /carts/{id}", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Update)))
	r.HandleFunc("DELETE /carts/{id}", RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Delete)))
	r.HandleFunc("GET /products", RouteHandler(middlware.IsAuthenticated(ctx, productApi.GetAll)))

	// Admin Routes
	r.HandleFunc("GET /users", RouteHandler(middlware.IsAdmin(ctx, userApi.GetAll)))
	r.HandleFunc("GET /users/{id}", RouteHandler(middlware.IsAdmin(ctx, userApi.GetOne)))
	r.HandleFunc("DELETE /users/{id}", RouteHandler(middlware.IsAdmin(ctx, userApi.Delete)))

	r.HandleFunc("POST /products", RouteHandler(middlware.IsAdmin(ctx, productApi.Create)))

	r.HandleFunc("POST /products/categories", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Create)))

	r.HandleFunc("GET /products/categories", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.GetAll)))

	r.HandleFunc("GET /products/categories/{id}", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.GetOne)))

	r.HandleFunc("PUT /products/categories/{id}", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Update)))

	r.HandleFunc("DELETE /products/categories/{id}", RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Delete)))

	r.HandleFunc("GET /products/{id}", RouteHandler(middlware.IsAdmin(ctx, productApi.GetOne)))
	r.HandleFunc("PUT /products/{id}", RouteHandler(middlware.IsAdmin(ctx, productApi.Update)))
	r.HandleFunc("DELETE /products/{id}", RouteHandler(middlware.IsAdmin(ctx, productApi.Delete)))
}
