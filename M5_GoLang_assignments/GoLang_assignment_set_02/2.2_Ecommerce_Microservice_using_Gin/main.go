package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Initialize the database
	err := InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Middleware
	r.Use(RequestLogger())
	r.Use(RateLimiter())

	// API routes
	api := r.Group("/api")
	{
		api.POST("/product", JWTAuth(), AddProduct)
		api.GET("/product/:id", GetProduct)
		api.PUT("/product/:id", JWTAuth(), UpdateProduct)
		api.DELETE("/product/:id", JWTAuth(), DeleteProduct)
		api.GET("/products", GetAllProducts)
	}

	// Start server
	log.Println("Starting server on :8080")
	r.Run(":8080")
}
