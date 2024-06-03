package main

import (
	"database/sql"
	"gofr.dev/pkg/gofr"
	"user-service/migrations"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"user-service/handler"
	"user-service/middleware"
	"user-service/service"
	"user-service/store"
)

func main() {
	// Load environment variables from configs/.env file
	envPath := filepath.Join("configs", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s", envPath)
	}

	// Get database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Create connection string
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	a := gofr.New()
	a.Migrate(migrations.All())

	userStore := store.New(db)
	userService := service.New(&userStore)
	userHandler := handler.New(&userService)

	r := mux.NewRouter()

	r.Use(middleware.JWTMiddleware)

	// Define routes
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", userHandler.LoginUser).Methods("POST")
	r.HandleFunc("/users/profile", userHandler.GetUserProfile).Methods("GET")

	// Start HTTP server
	log.Println("Starting server on port 8003")
	err = http.ListenAndServe(":8003", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
