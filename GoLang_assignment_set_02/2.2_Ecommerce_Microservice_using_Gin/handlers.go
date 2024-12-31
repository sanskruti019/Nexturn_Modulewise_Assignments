package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Add a new product
func AddProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}

// Get a product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")

	query := `SELECT * FROM products WHERE id = ?`
	row := db.QueryRow(query, id)

	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Update a product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ? WHERE id = ?`
	_, err := db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM products WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// Get all products with pagination
func GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := `SELECT * FROM products LIMIT ? OFFSET ?`
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var product Product
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
