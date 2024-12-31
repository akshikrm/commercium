package main

import (
	"akshidas/e-com/pkg/api"
	"akshidas/e-com/pkg/db"
	"akshidas/e-com/pkg/services"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const PORT = ":5234"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// sync := flag.Bool("sync", false, "seed database")
	// flag.Parse()

	// if *sync {
	//	paddlePayment := new(services.PaddlePayment)
	//	if err := paddlePayment.Init(); err != nil {
	//		panic(err)
	//	}
	//	paddlePayment.SyncPrice(storage)
	//	return
	// }

	database := connectDatabase()
	router := http.NewServeMux()
	router.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		api.Cors(w)
	})

	storage := db.NewStorage(database)

	ctx := context.Background()
	userService := services.NewUserService()
	userApi := api.NewUserApi(s.Store)
	middlware := api.NewMiddleWare(userApi.UserService)
	router.HandleFunc("POST /users", api.RouteHandler(userApi.Create))
	router.HandleFunc("POST /login", api.RouteHandler(userApi.Login))
	router.HandleFunc("GET /users/my-customer-id", api.RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetCustomerID)))
	router.HandleFunc("GET /profile", api.RouteHandler(middlware.IsAuthenticated(ctx, userApi.GetProfile)))
	router.HandleFunc("PUT /profile", api.RouteHandler(middlware.IsAuthenticated(ctx, userApi.UpdateProfile)))
	router.HandleFunc("GET /users", api.RouteHandler(middlware.IsAdmin(ctx, userApi.GetAll)))
	router.HandleFunc("GET /users/{id}", api.RouteHandler(middlware.IsAdmin(ctx, userApi.GetOne)))
	router.HandleFunc("DELETE /users/{id}", api.RouteHandler(middlware.IsAdmin(ctx, userApi.Delete)))

	log.Printf("🚀 Server started on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))

}

func getConnectionString() string {
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, name, password)
}

func connectDatabase() *sql.DB {
	connString := getConnectionString()
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
