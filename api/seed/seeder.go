package main

import (
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/storage"
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
	store *sql.DB
}

func (s *Seeder) INIT() {
	s.seedRoles()
	s.seedUsers()
}

func (s *Seeder) seedRoles() {
	roleStorage := storage.NewRoleStorage(s.store)
	roleService := services.NewRoleService(roleStorage)
	fmt.Print("Seeding Role...")
	for _, role := range ROLES {
		if err := roleService.Create(&role); err != nil {
			fmt.Printf("FAILED, %s\n", role.Name)
			fmt.Printf("ERR: %s\n", err)
		}
	}
	fmt.Println("SUCCESS")
}

func (s Seeder) seedUsers() {
	fmt.Print("Seeding Users...")
	userModel := storage.NewUserStorage(s.store)
	profileModel := storage.NewProfileStorage(s.store)
	userService := services.NewUserService(userModel, profileModel)
	userFile, err := os.Open("./seed/mock/users.json")
	if err != nil {
		fmt.Println("FAILED")
		fmt.Printf("ERR: %s\n", err)
		os.Exit(1)
	}
	defer userFile.Close()

	byteValue, err := io.ReadAll(userFile)
	if err != nil {
		fmt.Println("FAILED")
		fmt.Printf("ERR: %s\n", err)
		os.Exit(1)
	}

	users := []types.CreateUserRequest{}
	json.Unmarshal(byteValue, &users)
	for _, element := range users {
		if _, err := userService.Create(element); err != nil {
			fmt.Println("FAILED")
			fmt.Printf("ERR: %s\n", err)
			continue
		}
	}
	fmt.Println("SUCCESS")
}
