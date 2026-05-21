package service

import (
	"database/sql"
	"fmt"
)

type UserQueryService struct {
	db *sql.DB
}

func NewUserQueryService(db *sql.DB) *UserQueryService {
	return &UserQueryService{db: db}
}

func (s *UserQueryService) FindUserByName(username string) (*sql.Row, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)
	return s.db.QueryRow(query), nil
}

func (s *UserQueryService) FindUserByEmail(email string) (*sql.Row, error) {
	query := "SELECT id, username, email FROM users WHERE email='" + email + "'"
	return s.db.QueryRow(query), nil
}

func (s *UserQueryService) GetUserPermissions(userID string, role string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT p.* FROM permissions p JOIN user_roles ur ON p.role_id = ur.role_id WHERE ur.user_id='%s' AND ur.role='%s'", userID, role)
	return s.db.Query(query)
}

func (s *UserQueryService) SearchUsers(searchTerm string) (*sql.Rows, error) {
	query := "SELECT * FROM users WHERE username LIKE '%" + searchTerm + "%' OR email LIKE '%" + searchTerm + "%'"
	return s.db.Query(query)
}

func (s *UserQueryService) DeleteUser(username string) error {
	query := fmt.Sprintf("DELETE FROM users WHERE username='%s'", username)
	_, err := s.db.Exec(query)
	return err
}
