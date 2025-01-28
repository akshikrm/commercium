package app

import (
	"akshidas/e-com/pkg/handlers"
	"context"
)

type routes map[string]handlers.ApiFunc
type routesFunc func(*Server) routes

func UserRoute(s *Server) routes {
	ctx := context.Background()
	handler := s.handlers.User
	middleware := s.handlers.Middleware
	return routes{
		"POST /users":               handler.Create,
		"POST /login":               handler.Login,
		"GET /users/my-customer-id": middleware.IsAuthenticated(ctx, handler.GetCustomerID),
		"GET /profile":              middleware.IsAuthenticated(ctx, handler.GetProfile),
		"PUT /profile":              middleware.IsAuthenticated(ctx, handler.UpdateProfile),
		"GET /users":                middleware.IsAdmin(ctx, handler.GetAll),
		"GET /users/{id}":           middleware.IsAdmin(ctx, handler.GetOne),
		"DELETE /users/{id}":        middleware.IsAdmin(ctx, handler.Delete),
	}
}

func ProductRoute(s *Server) routes {
	ctx := context.Background()
	handler := s.handlers.Product
	middleware := s.handlers.Middleware
	return routes{
		"GET /products":         middleware.IsAuthenticated(ctx, handler.GetAll),
		"POST /products":        middleware.IsAdmin(ctx, handler.Create),
		"GET /products/{id}":    middleware.IsAdmin(ctx, handler.GetOne),
		"PUT /products/{id}":    middleware.IsAdmin(ctx, handler.Update),
		"DELETE /products/{id}": middleware.IsAdmin(ctx, handler.Delete),
	}
}

func ProductCategoryRoute(s *Server) routes {
	ctx := context.Background()
	handler := s.handlers.ProductCategory
	middleware := s.handlers.Middleware
	return routes{
		"POST /products/categories":        middleware.IsAdmin(ctx, handler.Create),
		"GET /products/categories":         middleware.IsAdmin(ctx, handler.GetAll),
		"GET /products/categories/{id}":    middleware.IsAdmin(ctx, handler.GetOne),
		"PUT /products/categories/{id}":    middleware.IsAdmin(ctx, handler.Update),
		"DELETE /products/categories/{id}": middleware.IsAdmin(ctx, handler.Delete),
	}
}

func CartRoute(s *Server) routes {
	ctx := context.Background()
	handler := s.handlers.Cart
	middleware := s.handlers.Middleware
	return routes{
		"POST /carts":        middleware.IsAuthenticated(ctx, handler.Create),
		"GET /carts":         middleware.IsAuthenticated(ctx, handler.GetAll),
		"PUT /carts/{id}":    middleware.IsAuthenticated(ctx, handler.Update),
		"DELETE /carts/{id}": middleware.IsAuthenticated(ctx, handler.Delete),
	}
}

func PurchaseRoute(s *Server) routes {
	ctx := context.Background()
	handler := s.handlers.Purchase
	middleware := s.handlers.Middleware
	return routes{
		"POST /transactions":          handler.HandleTransactionHook,
		"GET /orders/invoice/{txnId}": middleware.IsAuthenticated(ctx, handler.GetInvoice),
		"GET /orders/status/{txnId}":  middleware.IsAuthenticated(ctx, handler.GetOrderStatus),
		"GET /orders":                 middleware.IsAuthenticated(ctx, handler.GetAllOrders),
		"GET /shipping":               middleware.IsAdmin(ctx, handler.GetShippingInformation),
	}
}

func UploadRoute(s *Server) routes {
	ctx := context.Background()
	handler := s.handlers.Upload
	middleware := s.handlers.Middleware
	return routes{
		"POST /upload": middleware.IsAdmin(ctx, handler.Upload),
	}
}
