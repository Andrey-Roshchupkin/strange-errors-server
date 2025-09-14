package models

// Article represents an article in the system
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreateArticleRequest represents the request body for creating a new article
type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GoatResponse represents the response from the GOAT method
type GoatResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserRequest represents the request body for creating a new user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Message string     `json:"message"`
	Data    []Article  `json:"data,omitempty"`
	Status  string     `json:"status,omitempty"`
	Error   string     `json:"error,omitempty"`
}
