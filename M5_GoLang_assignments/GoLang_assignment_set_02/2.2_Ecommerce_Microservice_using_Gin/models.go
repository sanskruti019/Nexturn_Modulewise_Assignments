package main

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryID  int     `json:"category_id"`
}
