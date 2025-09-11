package main

import (
	"fmt"
	"log"
	"net/http"

	"strange-errors-server/internal/config"
	"strange-errors-server/internal/database"
	"strange-errors-server/internal/handlers"
)

func main() {
	fmt.Println("🚀 Starting Strange Errors Server in Go...")
	
	// Load configuration
	cfg := config.LoadConfig()
	
	// Initialize database
	db, err := database.New(cfg.DBPath)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()
	
	// Create handlers
	handler := handlers.New(db)
	goatHandler := handlers.NewGoatHandler()
	
	// Create router
	router := handlers.NewRouter(handler, goatHandler)
	
	// Set up routes with logging middleware
	httpHandler := router.SetupRoutes()
	
	fmt.Printf("🌐 Server running on http://localhost%s\n", cfg.Port)
	fmt.Println("📝 Demonstrating various error handling fallacies")
	fmt.Println("🔗 Available endpoints:")
	fmt.Println("   GET  /api/articles - Get articles (returns 777 instead of 200)")
	fmt.Println("   POST /api/article - Create article (returns 888/999 instead of 201/400)")
	fmt.Println("   DELETE /api/article/{id} - Delete article (returns 666 instead of 404)")
	fmt.Println("   GET  /api/health-check - Regular health check")
	fmt.Println("   GOAT /api/health-check - GOAT method (annoying server behavior)")
	fmt.Println("")
	fmt.Println("🐐 Try the GOAT method:")
	fmt.Printf("   curl -X GOAT http://localhost%s/api/health-check\n", cfg.Port)
	
	log.Fatal(http.ListenAndServe(cfg.Port, httpHandler))
}