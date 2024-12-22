package main

import (
	"akshidas/e-com/pkg/db"
	"akshidas/e-com/pkg/server"
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

	store := db.NewStorage()
	db.Connect(store)

	if *sync {
		paddlePayment := new(services.PaddlePayment)
		if err := paddlePayment.Init(); err != nil {
			panic(err)
		}
		paddlePayment.SyncPrice(store)
		return
	}

	server := &server.APIServer{
		Status: "Server is up and running",
		Port:   ":5234",
		Store:  store,
	}
	server.Run()
}
