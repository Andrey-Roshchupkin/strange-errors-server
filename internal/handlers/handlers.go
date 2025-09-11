package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"strange-errors-server/internal/database"
	"strange-errors-server/internal/models"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	db *database.DB
}

// New creates a new Handler instance
func New(db *database.DB) *Handler {
	return &Handler{db: db}
}

// GetArticlesHandler handles GET /api/articles - with wrong status code (777 instead of 200)
func (h *Handler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	articles, err := h.db.GetArticles()
	if err != nil {
		http.Error(w, "Database error", 500)
		return
	}

	// Wrong status code - should be 200, but we use 777
	w.WriteHeader(777)
	response := models.APIResponse{
		Message: "Data successfully retrieved!",
		Data:    articles,
	}
	json.NewEncoder(w).Encode(response)
}

// CreateArticleHandler handles POST /api/article - with wrong status codes
func (h *Handler) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		// Wrong status code - should be 400, but we use 999
		w.WriteHeader(999)
		response := models.APIResponse{
			Error:  "Failed to add article. Both title and content are required.",
			Status: "INCORRECT_REQUEST",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if article.Title == "" || article.Content == "" {
		// Wrong status code - should be 400, but we use 999
		w.WriteHeader(999)
		response := models.APIResponse{
			Error:  "Failed to add article. Both title and content are required.",
			Status: "INCORRECT_REQUEST",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.db.CreateArticle(article.Title, article.Content)
	if err != nil {
		http.Error(w, "Database error", 500)
		return
	}

	// Wrong status code - should be 201, but we use 888
	w.WriteHeader(888)
	response := models.APIResponse{
		Message: "New article added.",
		Status:  "OK",
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteArticleHandler handles DELETE /api/article/{id} - with wrong status codes
func (h *Handler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	idStr := r.URL.Path[len("/api/article/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Wrong status code - should be 400, but we use 500
		w.WriteHeader(500)
		response := models.APIResponse{
			Error:   "We're not even going to check for that. Something went wrong on our end.",
			Message: "Invalid input. We can only delete articles by their numeric ID.",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	rowsAffected, err := h.db.DeleteArticle(id)
	if err != nil {
		http.Error(w, "Database error", 500)
		return
	}

	if rowsAffected == 0 {
		// Wrong status code - should be 404, but we use 666
		w.WriteHeader(666)
		response := models.APIResponse{
			Message: "No evil articles found to remove.",
			Status:  "FAILURE",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(200)
		response := models.APIResponse{
			Message: fmt.Sprintf("Article with id %d has been removed.", id),
			Status:  "SUCCESS",
		}
		json.NewEncoder(w).Encode(response)
	}
}

// HealthCheckHandler handles GET /api/health-check - regular health check
func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":    "OK",
		"message":   "Server is running normally.",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}
