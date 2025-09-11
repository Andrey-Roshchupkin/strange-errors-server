package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"strange-errors-server/internal/middleware"
	"strange-errors-server/internal/models"
)

// Router handles HTTP routing
type Router struct {
	handler     *Handler
	goatHandler *GoatHandler
}

// NewRouter creates a new Router instance
func NewRouter(handler *Handler, goatHandler *GoatHandler) *Router {
	return &Router{
		handler:     handler,
		goatHandler: goatHandler,
	}
}

// Handler is the main HTTP handler that routes requests
func (r *Router) Handler(w http.ResponseWriter, req *http.Request) {
	log.Printf("📡 Incoming request: %s %s", req.Method, req.URL.Path)

	// Handle GOAT method for health check
	if req.Method == "GOAT" && req.URL.Path == "/api/health-check" {
		log.Println("🐐 GOAT method detected! Calling goatHandler...")
		r.goatHandler.Handle(w, req)
		return
	}

	// Route other requests
	switch req.URL.Path {
	case "/api/articles":
		r.handler.GetArticlesHandler(w, req)
	case "/api/article":
		if req.Method == "POST" {
			r.handler.CreateArticleHandler(w, req)
		} else {
			http.Error(w, "Method not allowed", 405)
		}
	case "/api/health-check":
		r.handler.HealthCheckHandler(w, req)
	default:
		// Check if it's a delete request for specific article
		if len(req.URL.Path) > 12 && req.URL.Path[:12] == "/api/article/" && req.Method == "DELETE" {
			r.handler.DeleteArticleHandler(w, req)
		} else {
			// Wrong status code for 404 - should be 404, but we use 200
			w.WriteHeader(200)
			response := models.APIResponse{
				Error:   "Route not found",
				Message: "Try a different endpoint",
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

// SetupRoutes sets up all routes with middleware
func (r *Router) SetupRoutes() http.HandlerFunc {
	return middleware.LogRequest(r.Handler)
}
