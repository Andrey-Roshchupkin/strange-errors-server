package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingResponseWriter wraps http.ResponseWriter to capture status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewLoggingResponseWriter creates a new LoggingResponseWriter
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader captures the status code
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LogRequest is a middleware that logs HTTP requests
func LogRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("🚀 %s %s - Started", r.Method, r.URL.Path)
		
		// Create a custom ResponseWriter to capture status code
		lrw := NewLoggingResponseWriter(w)
		
		handler(lrw, r)
		
		duration := time.Since(start)
		log.Printf("✅ %s %s - %d - %v", r.Method, r.URL.Path, lrw.statusCode, duration)
	}
}
