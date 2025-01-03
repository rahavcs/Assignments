package main

import (
	"blog-management/database"
	"blog-management/handlers"
	"blog-management/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Database
	database.InitDB()

	// Create Router
	r := mux.NewRouter()

	// Register Routes
	r.HandleFunc("/blogs", handlers.HandleBlogs).Methods("POST", "GET")
	r.HandleFunc("/blogs/{id:[0-9]+}", handlers.HandleBlogByID).Methods("GET", "PUT", "DELETE")

	// Apply Middleware
	r.Use(middleware.Logger)

	// Start HTTP Server
	log.Println("Server running at :8060...")
	log.Fatal(http.ListenAndServe(":8060", r))
}
