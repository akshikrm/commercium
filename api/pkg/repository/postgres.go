package repository

import (
	"akshidas/e-com/pkg/types"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Storage struct {
	DB              *sql.DB
	Product         types.ProductRepository
	ProductCategory types.ProductCategoriesRepository
	User            types.UserRepository
	Profile         types.ProfileRepository
	Cart            types.CartRepository
	Transaction     types.TransactionRepository
	Orders          types.OrdersRepository
	Role            types.RoleRepository
}

func New() *Storage {
	repository := new(Storage)
	database := connect()
	repository.DB = database

	repository.Product = newProduct(database)
	repository.ProductCategory = newProductCategory(database)
	repository.User = newUser(database)
	repository.Profile = newProfile(database)
	repository.Cart = newCart(database)
	repository.Role = newRole(database)
	repository.Transaction = newTransactions(database)
	repository.Orders = newOrders(database)

	return repository
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
