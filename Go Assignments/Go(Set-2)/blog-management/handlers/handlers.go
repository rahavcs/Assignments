package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"blog-management/models"
)

func HandleBlogs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		blogs, err := models.GetAllBlogs()
		if err != nil {
			log.Printf("Error fetching blogs: %v", err)
			http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(blogs); err != nil {
			log.Printf("Error encoding blogs to JSON: %v", err)
			http.Error(w, "Error processing response", http.StatusInternalServerError)
		}

	case http.MethodPost:
		var blog models.Blog
		if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
			log.Printf("Invalid JSON payload: %v", err)
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		if err := models.CreateBlog(blog); err != nil {
			log.Printf("Error creating blog: %v", err)
			http.Error(w, "Error creating blog", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleBlogByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/blog/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid blog ID: %v", err)
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		blog, err := models.GetBlogByID(id)
		if err != nil {
			log.Printf("Blog not found: %v", err)
			http.Error(w, "Blog not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(blog); err != nil {
			log.Printf("Error encoding blog to JSON: %v", err)
			http.Error(w, "Error processing response", http.StatusInternalServerError)
		}

	case http.MethodPut:
		var blog models.Blog
		if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
			log.Printf("Invalid JSON payload: %v", err)
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		if err := models.UpdateBlog(id, blog); err != nil {
			log.Printf("Error updating blog: %v", err)
			http.Error(w, "Error updating blog", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	case http.MethodDelete:
		if err := models.DeleteBlog(id); err != nil {
			log.Printf("Error deleting blog: %v", err)
			http.Error(w, "Error deleting blog", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
