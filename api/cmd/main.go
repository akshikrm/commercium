package main

import (
	"akshidas/e-com/pkg"
	"akshidas/e-com/pkg/db"
	"akshidas/e-com/pkg/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	pkg.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	pkg.WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	pkg.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	store := db.NewStorage()
	db.Connect(store)

	server := &server.APIServer{
		Status: "Server is up and running",
		Port:   ":5234",
		Store:  store,
	}
	server.Run()
}
