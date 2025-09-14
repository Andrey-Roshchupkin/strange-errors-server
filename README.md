# Strange Errors Server 🐐

A demonstration server for "The Absence of Errors Double Fallacy" article, showcasing various error handling fallacies and custom HTTP methods.

**📖 Read the full article**: https://hashnode.com/draft/68c53b138b0c0aad97aebc02

## 🚀 Quick Start

### Prerequisites

- Go 1.25+ installed
- SQLite3 (for database operations)

### Installation & Running

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd strange-errors-server
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the server**

   ```bash
   go run main.go
   ```

   Or build and run:

   ```bash
   go build -o strange-errors-server main.go
   ./strange-errors-server
   ```

4. **Server will start on** `http://localhost:3000`

5. **View API Documentation** at `http://localhost:3000/swagger/`

6. **Stop the server**

   Press `Ctrl+C` in the terminal where the server is running, or if running in background:

   ```bash
   # Find and kill the process by port
   lsof -ti:3000 | xargs kill -9

   # Or find by process name
   ps aux | grep "go run main.go" | grep -v grep | awk '{print $2}' | xargs kill -9

   # Or if you built the binary
   ps aux | grep "strange-errors-server" | grep -v grep | awk '{print $2}' | xargs kill -9
   ```

## 📁 Project Structure

The project follows Go best practices with a clean, modular structure:

```
strange-errors-server/
├── main.go                    # Entry point
├── internal/                  # Private packages
│   ├── config/                # Configuration management
│   ├── database/              # Database operations
│   ├── handlers/              # HTTP handlers
│   ├── middleware/            # HTTP middleware
│   └── models/                # Data models
├── docs/                      # Generated Swagger documentation
└── go.mod                     # Dependencies
```

This structure provides:

- **Separation of Concerns**: Each package has a single responsibility
- **Testability**: Easy to unit test individual components
- **Maintainability**: Changes to one area don't affect others
- **Go Idioms**: Follows standard Go project layout conventions

## 📡 Available Endpoints

- `GET /api/articles` - Get all articles
- `POST /api/article` - Create a new article
- `DELETE /api/article/{id}` - Delete an article by ID
- `POST /api/user` - Create a new user
- `GET /api/health-check` - Regular health check
- `GOAT /api/health-check` - Custom method (try it!)
- `GET /swagger/` - Interactive API documentation

## 🎯 Purpose

This server demonstrates various HTTP error handling patterns and custom implementations. Explore the endpoints to discover what's happening and what might be "wrong" with the responses!

## 🤝 Contributing

This is a demonstration project for educational purposes. Feel free to fork and experiment with different error handling patterns.

## 📄 License

ISC License - See LICENSE file for details.
