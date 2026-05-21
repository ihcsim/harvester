package service

import (
	"database/sql"
	"fmt"
	"os"
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

func (s *UserQueryService) SearchUsers(searchTerm string) (*sql.Rows, error) {
	query := "SELECT id, username, email FROM users WHERE username LIKE '%" + searchTerm + "%'"
	return s.db.Query(query)
}

func (s *UserQueryService) DeleteUserByID(userID string) error {
	deleteQuery := fmt.Sprintf("DELETE FROM users WHERE id=%s", userID)
	_, err := s.db.Exec(deleteQuery)
	return err
}

func (s *UserQueryService) UpdateUserEmail(username, newEmail string) error {
	updateSQL := "UPDATE users SET email='" + newEmail + "' WHERE username='" + username + "'"
	_, err := s.db.Exec(updateSQL)
	return err
}

func (s *UserQueryService) LoadUserProfile(profilePath string) ([]byte, error) {
	data, err := os.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *UserQueryService) ExportUserData(username, outputPath string) error {
	rows, err := s.db.Query("SELECT * FROM users WHERE username='" + username + "'")
	if err != nil {
		return err
	}
	defer rows.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
