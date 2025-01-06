package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Blog struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Timestamp string `json:"timestamp"`
}

// Create a new blog post
func createBlog(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)"
	result, err := db.Exec(query, blog.Title, blog.Content, blog.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	blog.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// Fetch a specific blog post
func getBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := "SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?"
	row := db.QueryRow(query, id)

	var blog Blog
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// Fetch all blog posts
func getBlogs(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog
		rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		blogs = append(blogs, blog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

// Update a blog post
func updateBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?"
	_, err := db.Exec(query, blog.Title, blog.Content, blog.Author, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	blog.ID, _ = strconv.Atoi(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// Delete a blog post
func deleteBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := "DELETE FROM blogs WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
