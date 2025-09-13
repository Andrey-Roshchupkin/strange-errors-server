package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"strange-errors-server/internal/models"
)

// GoatHandler handles the GOAT method - the annoying server behavior
type GoatHandler struct {
	callCount int
}

// NewGoatHandler creates a new GoatHandler instance
func NewGoatHandler() *GoatHandler {
	return &GoatHandler{callCount: 0}
}

// Handle handles the GOAT method with progressive annoyance
// @Summary GOAT method (Custom HTTP Method)
// @Description Demonstrates progressive server behavior with a custom GOAT HTTP method. Use: curl -X GOAT http://localhost:3000/api/health-check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} models.GoatResponse "First call - Happy GOAT"
// @Success 400 {object} models.GoatResponse "Second/Third call - Annoyed/Upset GOAT"
// @Success 500 {object} models.GoatResponse "Fourth call - Enraged GOAT"
// @Success 503 {object} models.GoatResponse "Fifth call - Fatal GOAT"
// @Router /api/health-check [post]
func (gh *GoatHandler) Handle(w http.ResponseWriter, r *http.Request) {
	gh.callCount++
	log.Printf("🐐 GOAT call #%d", gh.callCount)
	
	w.Header().Set("Content-Type", "application/json")

	var response models.GoatResponse

	switch gh.callCount {
	case 1:
		w.WriteHeader(200)
		response = models.GoatResponse{
			Status:  "OK",
			Message: "Hello! I am a brand new GOAT. Everything is fine.",
		}
	case 2:
		w.WriteHeader(400)
		response = models.GoatResponse{
			Status:  "Annoyed",
			Message: "Why are you calling me again?",
		}
	case 3:
		w.WriteHeader(400)
		response = models.GoatResponse{
			Status:  "Upset",
			Message: "You are hurting my feelings, I feel sick.",
		}
	case 4:
		// Delete the database!
		log.Println("💥 GOAT is enraged! Attempting to delete database...")
		err := os.Remove("./database.db")
		w.WriteHeader(500)
		if err != nil {
			log.Printf("❌ Failed to delete database: %v", err)
			response = models.GoatResponse{
				Status:  "Failed",
				Message: "I tried to delete the database, but it was already gone.",
			}
		} else {
			log.Println("💀 Database deleted successfully!")
			response = models.GoatResponse{
				Status:  "Enraged",
				Message: "That is it! I have deleted the database. Good luck now.",
			}
		}
	case 5:
		log.Println("💀 GOAT is fatal! Server will shutdown in 1 second...")
		w.WriteHeader(503)
		response = models.GoatResponse{
			Status:  "Fatal",
			Message: "You have called me one time too many. Goodbye.",
		}
		// Shutdown server after 1 second
		go func() {
			time.Sleep(1 * time.Second)
			log.Println("🔴 Server shutting down...")
			os.Exit(1)
		}()
	default:
		log.Println("💥 GOAT is overloaded! Server will shutdown in 1 second...")
		w.WriteHeader(500)
		response = models.GoatResponse{
			Status:  "Overloaded",
			Message: "I have had enough. I am shutting down.",
		}
		go func() {
			time.Sleep(1 * time.Second)
			log.Println("🔴 Server shutting down...")
			os.Exit(1)
		}()
	}

	json.NewEncoder(w).Encode(response)
}
