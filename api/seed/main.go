package main

import (
	"akshidas/e-com/pkg/db"
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

	store := db.NewStorage()
	db.Connect(store)

	database := Database{store: store.DB}
	seeder := Seeder{store: store.DB}

	initdb := flag.Bool("init-db", false, "initialize db if true")
	nukedb := flag.Bool("nuke-db", false, "clear everything in the database")
	refreshdb := flag.Bool("refresh-db", false, "clear everything in the database")
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

	if *refreshdb {
		database.DROP()
		database.INIT()
		seeder.INIT()
	}
}
