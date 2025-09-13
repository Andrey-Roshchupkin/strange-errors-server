# Strange Errors Server 🐐

A demonstration server for "The Absence of Errors Double Fallacy" article, showcasing various error handling fallacies and custom HTTP methods.

## 🚀 Features

### Custom HTTP Method: GOAT 🐐

The server implements a custom `GOAT` HTTP method that demonstrates progressive server annoyance. Try calling it multiple times to see what happens:

```bash
curl -X GOAT http://localhost:3000/api/health-check
```

**Warning**: This method will progressively become more destructive with each call!

### Interactive API Documentation 📚

The server includes comprehensive Swagger/OpenAPI documentation accessible at:

- **Swagger UI**: http://localhost:3000/swagger/
- **API Spec**: http://localhost:3000/swagger/doc.json

The documentation includes detailed information about all endpoints, request/response schemas, and status codes.

### Error Handling Fallacies ❌

The server demonstrates various error handling fallacies. Can you spot what's wrong with the HTTP status codes?

#### Articles API (Non-Idempotent POST)

- **GET** `/api/articles` - Get all articles
- **POST** `/api/article` - Create a new article (creates duplicates on repeated calls)
- **DELETE** `/api/article/{id}` - Delete an article by ID

#### Users API (Idempotent POST)

- **POST** `/api/user` - Create a new user (returns error if user already exists)

#### Health Check

- **GET** `/api/health-check` - Regular health check
- **GOAT** `/api/health-check` - Custom method with progressive behavior

## 🛠️ Technology Stack

- **Language**: Go 1.25
- **Database**: SQLite3
- **HTTP Server**: Native Go `net/http`
- **API Documentation**: Swagger/OpenAPI with `swaggo/swag`
- **Dependencies**:
  - `github.com/mattn/go-sqlite3` - SQLite driver
  - `github.com/swaggo/swag` - Swagger documentation generator
  - `github.com/swaggo/http-swagger` - Swagger UI handler

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

## 📁 Project Structure

The project follows Go best practices with a clean, modular structure:

```
strange-errors-server/
├── main.go                    # Entry point
├── internal/                  # Private packages
│   ├── config/               # Configuration management
│   ├── database/             # Database operations
│   ├── handlers/             # HTTP handlers
│   ├── middleware/           # HTTP middleware
│   └── models/              # Data models
├── docs/                     # Generated Swagger documentation
└── go.mod                   # Dependencies
```

This structure provides:

- **Separation of Concerns**: Each package has a single responsibility
- **Testability**: Easy to unit test individual components
- **Maintainability**: Changes to one area don't affect others
- **Go Idioms**: Follows standard Go project layout conventions

## 📡 API Endpoints

### Articles (Non-Idempotent POST)

```bash
# Get all articles
curl http://localhost:3000/api/articles

# Create article (creates duplicates on repeated calls)
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Test Article","content":"Test content"}' \
  http://localhost:3000/api/article

# Delete article
curl -X DELETE http://localhost:3000/api/article/1
```

### Users (Idempotent POST)

```bash
# Create user (returns error if user already exists)
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Alice"}' \
  http://localhost:3000/api/user

# Try creating the same user again - will return error
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Alice"}' \
  http://localhost:3000/api/user
```

### Health Check

```bash
# Regular health check
curl http://localhost:3000/api/health-check

# GOAT method (progressive annoyance)
curl -X GOAT http://localhost:3000/api/health-check
```

### API Documentation

```bash
# View Swagger UI
open http://localhost:3000/swagger/

# Get OpenAPI spec
curl http://localhost:3000/swagger/doc.json
```

## 🎯 Purpose

This server is designed to demonstrate the "Absence of Errors Double Fallacy" by:

1. **Custom HTTP Methods**: Shows how different languages/frameworks handle non-standard HTTP methods
2. **Status Code Issues**: Demonstrates how incorrect status codes can mislead developers
3. **Progressive Error States**: Shows how error handling can become more complex over time
4. **Destructive Behavior**: Illustrates how poor error handling can lead to data loss
5. **POST Idempotency Comparison**: Shows the difference between idempotent and non-idempotent POST methods
6. **Interactive Discovery**: Encourages readers to explore and discover issues themselves

## 🔄 POST Method Idempotency Comparison

This server demonstrates two different approaches to POST methods:

### **Non-Idempotent POST** (`/api/article`)

- **Behavior**: Creates a new article every time, even with identical data
- **Result**: Multiple calls with same data create duplicate articles
- **Status Codes**: Returns `888` instead of `201`, `999` instead of `400`
- **Use Case**: Traditional POST behavior for creating resources

### **Idempotent POST** (`/api/user`)

- **Behavior**: Checks if user already exists before creating
- **Result**: Returns error if user with same name already exists
- **Status Codes**: Returns correct `201` for success, `400` for conflicts
- **Use Case**: Demonstrates how POST can be made idempotent with proper validation

**Try both endpoints to see the difference!**

## 🔍 Key Differences from Express/Node.js

- **Native HTTP Method Support**: Go's `net/http` package natively supports custom HTTP methods
- **No Framework Limitations**: Direct control over HTTP request/response handling
- **Simpler Implementation**: Less boilerplate compared to Express middleware
- **Better Performance**: Go's HTTP server is significantly faster

## ⚠️ Warnings

- **Database Deletion**: The GOAT method will delete the database on the 4th call
- **Server Shutdown**: The GOAT method will shut down the server on the 5th call
- **Status Code Issues**: All endpoints return incorrect HTTP status codes for demonstration purposes
- **Interactive Learning**: Try the endpoints yourself to discover what's wrong!

## 📝 Logging

The server includes comprehensive logging:

- Request/response logging with timing
- GOAT method state tracking
- Database operation logging
- Error state progression
- Swagger documentation generation

## 🤝 Contributing

This is a demonstration project for educational purposes. Feel free to fork and experiment with different error handling patterns.

## 📄 License

ISC License - See LICENSE file for details.
