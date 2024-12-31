package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Storage struct {
	DB *sql.DB
}

func connect() *sql.DB {
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, name, password)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("🗃️ connected to database")
	return db
}

func NewStorage() *Storage {
	database := connect()
	return &Storage{
		DB: database,
	}
}
