package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	initDB()
	defer db.Close()

	// Start the HTTP server
	startServer()
}

func startServer() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/blog", createBlog).Methods("POST")
	router.HandleFunc("/blog/{id}", getBlog).Methods("GET")
	router.HandleFunc("/blogs", getBlogs).Methods("GET")
	router.HandleFunc("/blog/{id}", updateBlog).Methods("PUT")
	router.HandleFunc("/blog/{id}", deleteBlog).Methods("DELETE")

	// Apply middleware
	router.Use(loggingMiddleware)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
