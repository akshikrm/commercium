package repository

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/types"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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

func New(config *config.Config) *Storage {
	repository := new(Storage)
	database := connect(config)
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

func connect(config *config.Config) *sql.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Name, config.Password)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("🗃️ connected to database")
	return db
}
