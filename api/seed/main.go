package main

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}

	seederConfig := config.New()
	store := repository.New(seederConfig)
	service := services.New(store, seederConfig)

	database := NewDatabase(store)
	seeder := NewSeeder(service, seederConfig)

	initdb := flag.Bool("init-db", false, "initialize db if true")
	nukedb := flag.Bool("nuke-db", false, "clear everything in the database")
	refreshdb := flag.Bool("refresh-db", false, "clear everything in the database")
	seedProducts := flag.Bool("seed-product", false, "seed products only")
	seed := flag.Bool("seed", false, "seed database")
	flag.Parse()

	if *initdb {
		database.INIT()
	}

	if *nukedb {
		database.DROP()
	}

	if *seed {
		seeder.INIT()
	}

	if *seedProducts {
		seeder.seedProducts()
	}

	if *refreshdb {
		database.DROP()
		database.INIT()
		seeder.INIT()
	}
}
