package main

import (
	"akshidas/e-com/pkg/app"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"flag"
	"github.com/joho/godotenv"
	"log"
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
	app.New(":5234", services).Run()
}
