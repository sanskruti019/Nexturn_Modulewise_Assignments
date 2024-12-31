package main

import (
    "errors"
    "fmt"
    "sort"
    "strconv"
    "strings"
)

// Product represents an item in the inventory
type Product struct {
    ID    int
    Name  string
    Price float64
    Stock int
}

// InventoryManager handles all inventory operations
type InventoryManager struct {
    products []Product
}

// NewInventoryManager creates a new instance of InventoryManager
func NewInventoryManager() *InventoryManager {
    return &InventoryManager{
        products: make([]Product, 0),
    }
}

// AddProduct adds a new product to the inventory
func (im *InventoryManager) AddProduct(id int, name string, priceStr string, stock int) error {
    // Check for duplicate ID
    for _, p := range im.products {
        if p.ID == id {
            return fmt.Errorf("product with ID %d already exists", id)
        }
    }

    // Type casting string price to float64
    price, err := strconv.ParseFloat(priceStr, 64)
    if err != nil {
        return errors.New("invalid price format")
    }

    // Validate inputs
    if price <= 0 {
        return errors.New("price must be greater than zero")
    }
    if stock < 0 {
        return errors.New("stock cannot be negative")
    }
    if name == "" {
        return errors.New("product name cannot be empty")
    }

    // Create and add new product
    product := Product{
        ID:    id,
        Name:  name,
        Price: price,
        Stock: stock,
    }

    im.products = append(im.products, product)
    return nil
}

// UpdateStock updates the stock quantity of a product
func (im *InventoryManager) UpdateStock(id int, newStock int) error {
    if newStock < 0 {
        return errors.New("stock cannot be negative")
    }

    for i := range im.products {
        if im.products[i].ID == id {
            im.products[i].Stock = newStock
            return nil
        }
    }

    return fmt.Errorf("product with ID %d not found", id)
}

// SearchByID searches for a product by ID
func (im *InventoryManager) SearchByID(id int) (*Product, error) {
    for i := range im.products {
        if im.products[i].ID == id {
            return &im.products[i], nil
        }
    }
    return nil, fmt.Errorf("product with ID %d not found", id)
}

// SearchByName searches for products by name (case-insensitive partial match)
func (im *InventoryManager) SearchByName(name string) []Product {
    var results []Product
    searchTerm := strings.ToLower(name)

    for _, p := range im.products {
        if strings.Contains(strings.ToLower(p.Name), searchTerm) {
            results = append(results, p)
        }
    }

    return results
}

// SortByPrice sorts products by price
func (im *InventoryManager) SortByPrice() {
    sort.Slice(im.products, func(i, j int) bool {
        return im.products[i].Price < im.products[j].Price
    })
}

// SortByStock sorts products by stock quantity
func (im *InventoryManager) SortByStock() {
    sort.Slice(im.products, func(i, j int) bool {
        return im.products[i].Stock < im.products[j].Stock
    })
}

// DisplayInventory shows all products in a formatted table
func (im *InventoryManager) DisplayInventory() {
    if len(im.products) == 0 {
        fmt.Println("Inventory is empty")
        return
    }

    // Print table header
    fmt.Printf("\n%-5s | %-20s | %-10s | %-8s\n", "ID", "Name", "Price ($)", "Stock")
    fmt.Println(strings.Repeat("-", 50))

    // Print each product
    for _, p := range im.products {
        fmt.Printf("%-5d | %-20s | %10.2f | %-8d\n",
            p.ID, p.Name, p.Price, p.Stock)
    }
    fmt.Println()
}

func main() {
    // Create new inventory manager
    inventory := NewInventoryManager()

    // Add sample products
    samples := []struct {
        id    int
        name  string
        price string
        stock int
    }{
        {1, "Laptop", "999.99", 10},
        {2, "Mouse", "29.99", 50},
        {3, "Keyboard", "59.99", 30},
        {4, "Monitor", "299.99", 15},
        {5, "USB Cable", "9.99", 100},
    }

    // Add sample products and handle errors
    for _, s := range samples {
        err := inventory.AddProduct(s.id, s.name, s.price, s.stock)
        if err != nil {
            fmt.Printf("Error adding %s: %v\n", s.name, err)
        }
    }

    // Display original inventory
    fmt.Println("Original Inventory:")
    inventory.DisplayInventory()

    // Update stock
    err := inventory.UpdateStock(1, 15)
    if err != nil {
        fmt.Printf("Error updating stock: %v\n", err)
    }

    // Search by ID
    product, err := inventory.SearchByID(1)
    if err != nil {
        fmt.Printf("Search error: %v\n", err)
    } else {
        fmt.Printf("\nFound product by ID: %+v\n", *product)
    }

    // Search by name
    results := inventory.SearchByName("key")
    fmt.Printf("\nProducts containing 'key':\n")
    for _, p := range results {
        fmt.Printf("%+v\n", p)
    }

    // Sort by price and display
    fmt.Println("\nInventory sorted by price:")
    inventory.SortByPrice()
    inventory.DisplayInventory()

    // Sort by stock and display
    fmt.Println("\nInventory sorted by stock:")
    inventory.SortByStock()
    inventory.DisplayInventory()
}