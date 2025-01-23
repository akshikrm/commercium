package main

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/app"
	"akshidas/e-com/pkg/handlers"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
)

func main() {
	config := config.New()
	store := repository.New(config)
	services := services.New(store)
	handlers := handlers.New(services)
	server := app.New(":5234", handlers)

	server.RegisterRoutes(app.UserRoute)
	server.RegisterRoutes(app.ProductRoute)
	server.RegisterRoutes(app.ProductCategoryRoute)
	server.RegisterRoutes(app.CartRoute)
	server.RegisterRoutes(app.PurchaseRoute)

	server.Run()
}
