# Strange Errors Server 🐐

A demonstration server for "The Absence of Errors Double Fallacy" article, showcasing various error handling fallacies and custom HTTP methods.

## 🚀 Features

### Custom HTTP Method: GOAT 🐐

The server implements a custom `GOAT` HTTP method that demonstrates progressive server annoyance:

```bash
# 1st call - Happy
curl -X GOAT http://localhost:3000/api/health-check
# Response: {"status":"OK","message":"Hello! I am a brand new GOAT. Everything is fine."}

# 2nd call - Annoyed
curl -X GOAT http://localhost:3000/api/health-check
# Response: {"status":"Annoyed","message":"Why are you calling me again?"}

# 3rd call - Upset
curl -X GOAT http://localhost:3000/api/health-check
# Response: {"status":"Upset","message":"You are hurting my feelings, I feel sick."}

# 4th call - Enraged (DELETES DATABASE!)
curl -X GOAT http://localhost:3000/api/health-check
# Response: {"status":"Enraged","message":"That is it! I have deleted the database. Good luck now."}

# 5th call - Fatal (SHUTS DOWN SERVER!)
curl -X GOAT http://localhost:3000/api/health-check
# Response: {"status":"Fatal","message":"You have called me one time too many. Goodbye."}
```

### Error Handling Fallacies ❌

The server demonstrates various error handling fallacies with **wrong HTTP status codes**:

#### Articles API

- **GET** `/api/articles` - Returns `777` instead of `200`
- **POST** `/api/article` - Returns `888` instead of `201`, `999` instead of `400`
- **DELETE** `/api/article/{id}` - Returns `666` instead of `404`

#### Health Check

- **GET** `/api/health-check` - Regular health check (returns `200`)
- **GOAT** `/api/health-check` - Custom method with progressive behavior

## 🛠️ Technology Stack

- **Language**: Go 1.25
- **Database**: SQLite3
- **HTTP Server**: Native Go `net/http`
- **Dependencies**: `github.com/mattn/go-sqlite3`

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

## 📡 API Endpoints

### Articles

```bash
# Get all articles (returns 777 instead of 200)
curl http://localhost:3000/api/articles

# Create article (returns 888 instead of 201)
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Test Article","content":"Test content"}' \
  http://localhost:3000/api/article

# Delete article (returns 666 instead of 404)
curl -X DELETE http://localhost:3000/api/article/1
```

### Health Check

```bash
# Regular health check
curl http://localhost:3000/api/health-check

# GOAT method (progressive annoyance)
curl -X GOAT http://localhost:3000/api/health-check
```

## 🎯 Purpose

This server is designed to demonstrate the "Absence of Errors Double Fallacy" by:

1. **Custom HTTP Methods**: Shows how different languages/frameworks handle non-standard HTTP methods
2. **Wrong Status Codes**: Demonstrates how incorrect status codes can mislead developers
3. **Progressive Error States**: Shows how error handling can become more complex over time
4. **Destructive Behavior**: Illustrates how poor error handling can lead to data loss

## 🔍 Key Differences from Express/Node.js

- **Native HTTP Method Support**: Go's `net/http` package natively supports custom HTTP methods
- **No Framework Limitations**: Direct control over HTTP request/response handling
- **Simpler Implementation**: Less boilerplate compared to Express middleware
- **Better Performance**: Go's HTTP server is significantly faster

## ⚠️ Warnings

- **Database Deletion**: The GOAT method will delete the database on the 4th call
- **Server Shutdown**: The GOAT method will shut down the server on the 5th call
- **Wrong Status Codes**: All endpoints return incorrect HTTP status codes for demonstration purposes

## 📝 Logging

The server includes comprehensive logging:

- Request/response logging with timing
- GOAT method state tracking
- Database operation logging
- Error state progression

## 🤝 Contributing

This is a demonstration project for educational purposes. Feel free to fork and experiment with different error handling patterns.

## 📄 License

ISC License - See LICENSE file for details.
