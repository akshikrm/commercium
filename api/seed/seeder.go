package main

import (
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	store   *sql.DB
	service *services.Service
}

func (s *Seeder) INIT() {
	storage := repository.New()
	s.service = services.New(storage)

	s.seedRoles()
	s.seedUsers()
	s.seedProductCategories()
	s.seedProducts()
}

func (s *Seeder) seedRoles() {
	fmt.Print("Seeding Role...")
	for _, role := range ROLES {
		if err := s.service.Role.Create(&role); err != nil {
			fmt.Printf("FAILED, %s\n", role.Name)
			fmt.Printf("ERR: %s\n", err)
		}
	}
	fmt.Println("SUCCESS")
}

func (s Seeder) seedUsers() {
	fmt.Print("Seeding Users...")

	byteValue := readFile("./seed/mock/users.json")
	users := []types.CreateUserRequest{}
	json.Unmarshal(byteValue, &users)

	for _, element := range users {
		if _, err := s.service.User.Create(element); err != nil {
			fmt.Println("FAILED")
			fmt.Printf("ERR: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Println("SUCCESS")
}

func (s Seeder) seedProductCategories() {
	file := readFile("./seed/mock/product-categories.json")
	productCategories := []types.NewProductCategoryRequest{}
	json.Unmarshal(file, &productCategories)

	fmt.Print("Seeding product categories...")
	for _, product := range productCategories {
		if _, err := s.service.ProductCategory.Create(&product); err != nil {
			fmt.Println("FAILED")
			fmt.Printf("ERR: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Println("SUCCESS")

}

func (s Seeder) seedProducts() {
	fmt.Print("Seeding products...")
	file := readFile("./seed/mock/products.json")
	products := []types.NewProductRequest{}
	json.Unmarshal(file, &products)

	for _, product := range products {
		if err := s.service.Product.Create(&product); err != nil {
			fmt.Println("FAILED")
			fmt.Printf("ERR: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Println("SUCCESS")
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
