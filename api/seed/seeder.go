package main

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/PaddleHQ/paddle-go-sdk"
)

var ROLES = []types.CreateRoleRequest{
	{
		Name:        "Admin",
		Code:        "admin",
		Description: "Role assigned to admin",
	},
	{
		Name:        "User",
		Code:        "user",
		Description: "Role assigned to user",
	},
}

type Seeder struct {
	service *services.Service
	paddle  *services.PaddlePayment
}

func (s *Seeder) INIT() {
	s.seedRoles()
	s.seedUsers()
	s.seedProductCategories()
	s.seedProducts()
}

func (s *Seeder) seedRoles() {
	fmt.Print("SEEDING Role...")
	for _, role := range ROLES {
		if err := s.service.Role.Create(&role); err != nil {
			fmt.Printf("FAILED, %s\n", role.Name)
			fmt.Printf("ERR: %s\n", err)
		}
	}
	fmt.Println("SUCCESS")
}

func (s Seeder) seedUsers() {
	fmt.Println("SEEDING Users...")

	byteValue := readFile("./seed/mock/users.json")
	users := []types.CreateUserRequest{}
	json.Unmarshal(byteValue, &users)
	for _, element := range users {
		if s.service.User.Exists(element.Email) {
			fmt.Printf("SKIPPING %s already exists\n", element.Email)
			continue
		}
		if err := s.paddle.CreateCustomer(&element); err != nil {
			if err == paddle.ErrCustomerAlreadyExists {
				if customerID, err := s.paddle.GetCustomerByEmail(element.Email); err == nil {
					element.CustomerID = customerID
				} else {
					fmt.Printf("paddle error failed to get user %s\n", element.Email)
					continue
				}
			} else {
				fmt.Printf("failed to create %s for paddle due to %s\n", element.Email, err)
				continue
			}
		}
		_, err := s.service.User.Create(element)
		if err != nil {
			fmt.Printf("SEEDING %s FAILED\n", element.Email)
		} else {
			fmt.Printf("SEEDING %s SUCCESS\n", element.Email)
		}
	}
}

func (s Seeder) seedProductCategories() {
	file := readFile("./seed/mock/product-categories.json")
	productCategories := []types.NewProductCategoryRequest{}
	json.Unmarshal(file, &productCategories)

	fmt.Println("SEEDING product categories...")
	for _, product := range productCategories {
		if _, err := s.service.ProductCategory.Create(&product); err != nil {
			fmt.Printf("SEEDING %s FAILED\n", product.Name)
			fmt.Printf("ERR: %s\n", err)
		} else {
			fmt.Printf("SEEDING %s SUCCESS\n", product.Name)
		}
	}
}

func (s Seeder) seedProducts() {
	fmt.Println("SEEDING products...")
	file := readFile("./seed/mock/products.json")
	products := []types.NewProductRequest{}
	json.Unmarshal(file, &products)

	for _, product := range products {
		if err := s.service.Product.Create(&product); err != nil {
			fmt.Printf("SEEDING %s FAILED\n", product.Name)
			fmt.Printf("ERR: %s\n", err)
		} else {
			fmt.Printf("SEEDING %s SUCCESS\n", product.Name)

		}
	}
}

func NewSeeder(service *services.Service, seederConfig *config.Config) *Seeder {
	s := new(Seeder)
	s.service = service
	s.paddle = services.NewPaddlePayment(seederConfig)
	return s
}

func readFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("FAILED")
		fmt.Printf("ERR: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("FAILED")
		fmt.Printf("ERR: %s\n", err)
		os.Exit(1)
	}
	return byteValue
}
