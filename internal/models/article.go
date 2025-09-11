package models

// Article represents an article in the system
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GoatResponse represents the response from the GOAT method
type GoatResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Message string     `json:"message"`
	Data    []Article  `json:"data,omitempty"`
	Status  string     `json:"status,omitempty"`
	Error   string     `json:"error,omitempty"`
}
