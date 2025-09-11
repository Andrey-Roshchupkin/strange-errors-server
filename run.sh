#!/bin/bash

echo "🚀 Starting Strange Errors Server..."
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.25+ first."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | sed 's/go//')
echo "✅ Go version: $GO_VERSION"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy

# Build and run
echo "🔨 Building server..."
go build -o strange-errors-server main.go

echo ""
echo "🌐 Starting server on http://localhost:3000"
echo "🐐 Try the GOAT method: curl -X GOAT http://localhost:3000/api/health-check"
echo ""

./strange-errors-server
