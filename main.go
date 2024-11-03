package main

import (
	"gfgoth/api/routes"
	"gfgoth/pkg/book"

	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}
	defer db.Close()
	fmt.Println("Database connection successful!")

	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the gfgoth MySQL book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)

	log.Println("Attempting to connect to MySQL...") // Log connection attempt

	if err != nil {
		log.Printf("Failed to connect to MySQL at %s: %v", dbHost, err)
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping MySQL: %v", err)
		return nil, fmt.Errorf("database ping error: %w", err)
	}

	return db, nil
}