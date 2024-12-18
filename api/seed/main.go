package main

import (
	"akshidas/e-com/pkg/db"
	// "database/sql"
	// "akshidas/e-com/pkg/services"
	// "akshidas/e-com/pkg/storage"
	// "akshidas/e-com/pkg/types"
	// "encoding/json"
	"flag"
	// "io"
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
	}
}

// func Seed(s *db.Storage) error {
// 	initdb := flag.Bool("init-db", false, "initialize db if true")
// 	seedUsers := flag.Bool("seed-users", false, "seed db if true")
// 	seedResources := flag.Bool("seed-resources", false, "seed db if true")
// 	seedPermission := flag.Bool("seed-permission", false, "seed db if true")
// 	seedProducts := flag.Bool("seed-products", false, "seed db if true")
// 	seedProductCategory := flag.Bool("seed-product-categories", false, "seed db if true")
// 	nukeDb := flag.Bool("nuke-db", false, "clear everything in the database")
// 	refreshDb := flag.Bool("refresh-db", false, "clear everything in the database")
//
// 	flag.Parse()
// 	if *initdb {
// 		Init(s)
// 		os.Exit(0)
// 	}
//
// 	if *seedProducts {
// 		seedProductsFunc(s)
// 		os.Exit(0)
// 	}
//
// 	if *seedProductCategory {
// 		seedProductsCategoriesFunc(s)
// 		os.Exit(0)
// 	}
// 	if *seedResources {
// 		seedResourcesFunc(s)
// 		os.Exit(0)
// 	}
//
// 	if *seedUsers {
// 		seedUsersFunc(s)
// 		os.Exit(0)
// 	}
// 	if *seedPermission {
// 		seedPermissionFunc(s)
// 		os.Exit(0)
// 	}
//
// 	if *refreshDb {
// 		NukeDB(s)
// 		Init(s)
// 		seedResourcesFunc(s)
// 		seedPermissionFunc(s)
// 		seedUsersFunc(s)
// 		os.Exit(0)
// 	}
//
// 	if *nukeDb {
// 		NukeDB(s)
// 		os.Exit(0)
// 	}
//
// 	return nil
// }

// func seedAdminFunc(s *db.Storage) {
// 	log.Println("seeding admin")
// 	userModel := storage.NewUserStorage(s.DB)
// 	profileModel := storage.NewProfileStorage(s.DB)
// 	userService := services.NewUserService(userModel, profileModel)
// 	user := types.CreateUserRequest{
// 		FirstName: "Admin",
// 		LastName:  "Me",
// 		Email:     "admin@me.com",
// 		Password:  "root",
// 		Role:      "admin",
// 	}
// 	_, err := userService.Create(user)
// 	if err != nil {
// 		log.Printf("Failed to seed admin due to %s\n", err)
// 	}
// 	log.Println("Successfully seed admin")
// }
//
// func seedProductsCategoriesFunc(s *db.Storage) {
// 	log.Println("seeding products categories")
// 	productStorage := storage.NewProductCategoryStorage(s.DB)
// 	productService := services.NewProductCategoryService(productStorage)
// 	file := readFile("./seed/product-categories.json")
// 	productCategories := []types.NewProductCategoryRequest{}
// 	json.Unmarshal(file, &productCategories)
//
// 	for i, product := range productCategories {
// 		fmt.Println(i, product)
// 		if _, err := productService.Create(&product); err != nil {
// 			log.Printf("Failed to add product category %s due to %s\n", product.Name, err)
// 			continue
// 		}
// 		fmt.Println("Inserted product category")
// 	}
// 	fmt.Println("Finished seeding product category")
// }
//
// func seedProductsFunc(s *db.Storage) {
// 	log.Println("seeding products")
// 	productStorage := storage.NewProductStorage(s.DB)
// 	productService := services.NewProductService(productStorage)
// 	file := readFile("./seed/products.json")
// 	products := []types.CreateNewProduct{}
// 	json.Unmarshal(file, &products)
//
// 	for i, product := range products {
// 		fmt.Println(i, product)
// 		if err := productService.Create(&product); err != nil {
// 			log.Printf("Failed to add product %s due to %s\n", product.Name, err)
// 			continue
// 		}
// 		fmt.Println("Inserted product")
// 	}
// 	fmt.Println("Finished seeding product")
// }
//
// func seedUsersFunc(s *db.Storage) {
// 	log.Println("seeding users")
// 	userModel := storage.NewUserStorage(s.DB)
// 	profileModel := storage.NewProfileStorage(s.DB)
// 	userService := services.NewUserService(userModel, profileModel)
// 	userFile, err := os.Open("./seed/users.json")
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	defer userFile.Close()
//
// 	byteValue, err := io.ReadAll(userFile)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
//
// 	users := []types.CreateUserRequest{}
// 	json.Unmarshal(byteValue, &users)
// 	for i, element := range users {
// 		if _, err := userService.Create(element); err != nil {
// 			log.Printf("Failed to add user %s due to %s\n", element.Email, err)
// 			continue
// 		}
// 		log.Printf("Inserting %d\n", i)
// 	}
// 	log.Println("Successfully seed users")
// }
//
// func readFile(filePath string) []byte {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
//
// 	byteValue, err := io.ReadAll(file)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	return byteValue
//
// }
