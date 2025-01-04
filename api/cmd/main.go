package main

import (
	"akshidas/e-com/pkg/app"
	"akshidas/e-com/pkg/handlers"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	store := repository.New()
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
