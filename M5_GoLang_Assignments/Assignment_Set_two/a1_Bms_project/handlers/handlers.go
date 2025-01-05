package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"a1_bms_project/database"
	"a1_bms_project/models"
)

// CreateBlog handles creating a new blog post
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(query, blog.Title, blog.Content, blog.Author, time.Now())
	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Blog post created successfully"))
}

// GetBlog handles fetching a blog post by ID
func GetBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	query := "SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?"
	row := database.DB.QueryRow(query, id)

	var blog models.BlogPost
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Blog not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to fetch blog", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// GetAllBlogs handles fetching all blog posts
func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC"
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []models.BlogPost
	for rows.Next() {
		var blog models.BlogPost
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			http.Error(w, "Error reading blog data", http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, blog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

// UpdateBlog handles updating a blog post
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var blog models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?"
	_, err := database.DB.Exec(query, blog.Title, blog.Content, blog.Author, id)
	if err != nil {
		http.Error(w, "Failed to update blog", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Blog post updated successfully"))
}

// DeleteBlog handles deleting a blog post
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM blogs WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Blog post deleted successfully"))
}
