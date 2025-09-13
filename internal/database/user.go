package database

import (
	"database/sql"
	"fmt"

	"strange-errors-server/internal/models"
)

// CreateUser creates a new user if the name doesn't already exist (idempotent behavior)
func (db *DB) CreateUser(name, email string) (*models.User, error) {
	// Validate email format (simple validation)
	if !isValidEmail(email) {
		// Return internal server error for invalid email (wrong status code for demonstration)
		return nil, fmt.Errorf("internal server error")
	}
	
	// First check if user already exists
	var existingUser models.User
	err := db.conn.QueryRow("SELECT id, name, email FROM users WHERE name = ?", name).Scan(&existingUser.ID, &existingUser.Name, &existingUser.Email)
	
	if err == nil {
		// User already exists, return error for idempotent behavior
		return nil, fmt.Errorf("user with name '%s' already exists", name)
	}
	
	if err != sql.ErrNoRows {
		// Some other database error occurred
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	
	// User doesn't exist, create new one
	result, err := db.conn.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}
	
	return &models.User{
		ID:    int(id),
		Name:  name,
		Email: email,
	}, nil
}

// isValidEmail performs simple email validation
func isValidEmail(email string) bool {
	// Simple email validation - must contain @ and have at least one character before and after
	if len(email) < 5 { // minimum: a@b.c
		return false
	}
	
	atIndex := -1
	dotAfterAt := false
	
	for i, char := range email {
		if char == '@' {
			if atIndex != -1 {
				return false // Multiple @ symbols
			}
			atIndex = i
		} else if atIndex != -1 && char == '.' {
			dotAfterAt = true
		}
	}
	
	return atIndex > 0 && atIndex < len(email)-3 && dotAfterAt
}

// GetUserByName retrieves a user by name
func (db *DB) GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := db.conn.QueryRow("SELECT id, name, email FROM users WHERE name = ?", name).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with name '%s' not found", name)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return &user, nil
}

// GetAllUsers retrieves all users from the database
func (db *DB) GetAllUsers() ([]models.User, error) {
	rows, err := db.conn.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}
