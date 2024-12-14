package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/vsualzm/golang-project/helper"
	"github.com/vsualzm/golang-project/route"
)

type User struct {
	ID       int
	username string
	Password string
	Role     string
	Token    string
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// ANSI color codes

// LoggingMiddleware applies logging with colors to all requests

// CheckHealthAPI handles health check requests

func main() {

	connStr := "user=postgres password=1234 dbname=practice_db port=5432 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Memeriksa koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// db, err := database.ConnectDB()
	// if err != nil {
	// 	log.Fatalf("Error connecting to the database: %v", err)
	// }
	// defer db.Close()

	fmt.Println("Connected to the database")

	// Register your routes
	http.HandleFunc("/check-health", route.CheckHealthAPI)
	http.HandleFunc("/check-body", CheckBody)

	// Wrap the default ServeMux with the logging middleware
	wrappedMux := helper.LoggingMiddleware(http.DefaultServeMux)

	// Start server with wrapped handler
	fmt.Println("===========SERVICE RUNNING===========")
	fmt.Println("Server running on port 9000")
	if err := http.ListenAndServe(":9000", wrappedMux); err != nil {
		log.Fatalf("%sFailed to start server: %v%s", helper.Red, err, helper.Reset)
	}
}

func CheckBody(w http.ResponseWriter, r *http.Request) {

	var product Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(product.Name)
	log.Println(product.Price)
	log.Println(product.Stock)

	w.Header().Set("Content-Type", "application/json")

	// Mengencode data struct menjadi JSON dan mengirimkannya sebagai respons
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
