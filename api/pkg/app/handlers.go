package app

import (
	"akshidas/e-com/pkg/api"
	"akshidas/e-com/pkg/services"
	"context"
	"net/http"
)

func (s *app) registerRoutes(services *services.Service) {
	ctx := context.Background()
	r := s.router

	r.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		api.Cors(w)
	})

	// Api
	userApi := api.NewUserApi(services.User)
	productApi := api.NewProductApi(services.Product)
	cartApi := api.NewCartApi(services.Cart)
	productCategoryApi := api.NewProductCategoriesApi(services.ProductCategory)
	// uploadApi := api.NewUploadApi(s.Store)
	purchaseApi := api.NewOrdersApi(services.Transaction, services.Purchase)

	// Middle wares
	middlware := api.NewMiddleWare(userApi.UserService)

	// Public Routes
	r.HandleFunc("POST /transactions", api.RouteHandler(purchaseApi.HandleTransactionHook))
	r.HandleFunc("POST /users", api.RouteHandler(userApi.Create))
	r.HandleFunc("POST /login", api.RouteHandler(userApi.Login))
	// r.HandleFunc("POST /upload", api.RouteHandler(uploadApi.Upload))

	// Authenticated Routes
	r.HandleFunc("GET /users/my-customer-id", api.RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetCustomerID)))
	r.HandleFunc("GET /orders/invoice/{txnId}", api.RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetInvoice)))
	r.HandleFunc("GET /orders/status/{txnId}", api.RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetOrderStatus)))
	r.HandleFunc("GET /orders", api.RouteHandler(middlware.IsAuthenticated(ctx, purchaseApi.GetMyOrders)))

	r.HandleFunc("GET /profile", api.RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetProfile)))
	r.HandleFunc("PUT /profile", api.RouteHandler(middlware.IsAuthenticated(ctx, userApi.UpdateProfile)))
	r.HandleFunc("POST /carts", api.RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Create)))
	r.HandleFunc("GET /carts", api.RouteHandler(middlware.IsAuthenticated(ctx, cartApi.GetAll)))
	r.HandleFunc("PUT /carts/{id}", api.RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Update)))
	r.HandleFunc("DELETE /carts/{id}", api.RouteHandler(middlware.IsAuthenticated(ctx, cartApi.Delete)))
	r.HandleFunc("GET /products", api.RouteHandler(middlware.IsAuthenticated(ctx, productApi.GetAll)))

	// Admin Routes
	r.HandleFunc("GET /users", api.RouteHandler(middlware.IsAdmin(ctx, userApi.GetAll)))
	r.HandleFunc("GET /users/{id}", api.RouteHandler(middlware.IsAdmin(ctx, userApi.GetOne)))
	r.HandleFunc("DELETE /users/{id}", api.RouteHandler(middlware.IsAdmin(ctx, userApi.Delete)))

	r.HandleFunc("POST /products", api.RouteHandler(middlware.IsAdmin(ctx, productApi.Create)))

	r.HandleFunc("POST /products/categories", api.RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Create)))

	r.HandleFunc("GET /products/categories", api.RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.GetAll)))

	r.HandleFunc("GET /products/categories/{id}", api.RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.GetOne)))

	r.HandleFunc("PUT /products/categories/{id}", api.RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Update)))

	r.HandleFunc("DELETE /products/categories/{id}", api.RouteHandler(middlware.IsAdmin(ctx, productCategoryApi.Delete)))

	r.HandleFunc("GET /products/{id}", api.RouteHandler(middlware.IsAdmin(ctx, productApi.GetOne)))
	r.HandleFunc("PUT /products/{id}", api.RouteHandler(middlware.IsAdmin(ctx, productApi.Update)))
	r.HandleFunc("DELETE /products/{id}", api.RouteHandler(middlware.IsAdmin(ctx, productApi.Delete)))
}
