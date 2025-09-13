package main

import (
	"fmt"
	"log"
	"net/http"

	"strange-errors-server/internal/config"
	"strange-errors-server/internal/database"
	"strange-errors-server/internal/handlers"

	_ "strange-errors-server/docs" // This is the generated docs package
)

// @title Strange Errors Server API
// @version 1.0
// @description A demonstration server for "The Absence of Errors Double Fallacy" article, showcasing various error handling fallacies and custom HTTP methods.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name ISC
// @license.url https://opensource.org/licenses/ISC

// @host localhost:3000
// @BasePath /
// @schemes http

// @tag.name articles
// @tag.description Article management operations

// @tag.name users
// @tag.description User management operations

// @tag.name health
// @tag.description Health check and GOAT method operations

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
	fmt.Println("   POST /api/user - Create user (idempotent POST - returns 201/400)")
	fmt.Println("   GET  /api/health-check - Regular health check")
	fmt.Println("   GOAT /api/health-check - GOAT method (annoying server behavior)")
	fmt.Println("   GET  /swagger/ - Swagger API documentation")
	fmt.Println("")
	fmt.Println("🐐 Try the GOAT method:")
	fmt.Printf("   curl -X GOAT http://localhost%s/api/health-check\n", cfg.Port)
	fmt.Println("📚 View API documentation:")
	fmt.Printf("   http://localhost%s/swagger/\n", cfg.Port)
	
	log.Fatal(http.ListenAndServe(cfg.Port, httpHandler))
}