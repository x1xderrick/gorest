package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/x1xderrick/gorest/internal/controllers"
	"github.com/x1xderrick/gorest/internal/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env variables.")
	}
	fmt.Println("Connecting to SQLite3 Database...")

	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_CONNECTION_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Successfully connected.")

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create
	db.Create(&models.Product{Code: "D42", Price: 100})

	// Read
	var product models.Product
	db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", controllers.Index)

	b, err := json.Marshal(product)
	fmt.Println(string(b))

	fmt.Printf("Listening at %s:%s\n", os.Getenv("SERVER_HOSTNAME"), os.Getenv("SERVER_PORT"))
	http.ListenAndServe(os.Getenv("SERVER_HOSTNAME") + ":" + os.Getenv("SERVER_PORT"), router)

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
