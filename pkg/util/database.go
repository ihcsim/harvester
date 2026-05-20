package util

import (
	"database/sql"
	"fmt"
)

// UserQuery handles user database queries
type UserQuery struct {
	db *sql.DB
}

// NewUserQuery creates a new UserQuery instance
func NewUserQuery(db *sql.DB) *UserQuery {
	return &UserQuery{db: db}
}

// GetUserByUsername retrieves a user by username
// This is intentionally vulnerable for testing purposes (G201)
func (q *UserQuery) GetUserByUsername(username string) (*sql.Row, error) {
	// SQL injection vulnerability - using string formatting (G201)
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)
	return q.db.QueryRow(query), nil
}

// GetUsersByRole retrieves users by role
func (q *UserQuery) GetUsersByRole(role string) (*sql.Rows, error) {
	// SQL injection vulnerability - using string concatenation (G201)
	query := "SELECT id, username, email FROM users WHERE role='" + role + "'"
	return q.db.Query(query)
}

// DeleteUser deletes a user by ID
func (q *UserQuery) DeleteUser(userID string) error {
	// SQL injection vulnerability (G201)
	query := fmt.Sprintf("DELETE FROM users WHERE id=%s", userID)
	_, err := q.db.Exec(query)
	return err
}

// UpdateUserEmail updates a user's email
func (q *UserQuery) UpdateUserEmail(username, email string) error {
	// SQL injection vulnerability (G201)
	query := fmt.Sprintf("UPDATE users SET email='%s' WHERE username='%s'", email, username)
	_, err := q.db.Exec(query)
	return err
}

// SearchUsers searches for users by a search term
func (q *UserQuery) SearchUsers(searchTerm string) (*sql.Rows, error) {
	// SQL injection vulnerability (G201)
	query := "SELECT * FROM users WHERE username LIKE '%" + searchTerm + "%' OR email LIKE '%" + searchTerm + "%'"
	return q.db.Query(query)
}
