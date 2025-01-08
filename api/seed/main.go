package main

import (
	"akshidas/e-com/pkg/repository"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}

	store := repository.New()

	database := Database{store: store.DB}
	seeder := Seeder{store: store.DB}

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
