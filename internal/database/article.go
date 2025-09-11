package database

import (
	"fmt"

	"strange-errors-server/internal/models"
)

// GetArticles retrieves all articles from the database
func (db *DB) GetArticles() ([]models.Article, error) {
	rows, err := db.conn.Query("SELECT id, title, content FROM articles")
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %w", err)
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to scan article: %w", err)
		}
		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return articles, nil
}

// CreateArticle creates a new article in the database
func (db *DB) CreateArticle(title, content string) error {
	_, err := db.conn.Exec("INSERT INTO articles (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}
	return nil
}

// DeleteArticle deletes an article by ID from the database
func (db *DB) DeleteArticle(id int) (int64, error) {
	result, err := db.conn.Exec("DELETE FROM articles WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return rowsAffected, nil
}
