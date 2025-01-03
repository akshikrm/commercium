package main

import (
	"akshidas/e-com/pkg/app"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	sync := flag.Bool("sync", false, "seed database")
	flag.Parse()

	store := repository.New()

	if *sync {
		paddlePayment := new(services.PaddlePayment)
		if err := paddlePayment.Init(); err != nil {
			panic(err)
		}
		paddlePayment.SyncPrice(store)
		return
	}

	services := services.New(store)

	server := app.New(":5324", services)

	server.RegisterRoutes(app.UserHandler)
	server.RegisterRoutes(app.ProductHandler)
	server.RegisterRoutes(app.ProductCategoryHandler)
	server.RegisterRoutes(app.CartHandler)
	server.RegisterRoutes(app.PurchaseHandler)
	server.Run()
}
