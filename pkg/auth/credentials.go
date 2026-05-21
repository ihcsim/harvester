package auth

import (
	"database/sql"
	"fmt"
)

const (
	DefaultAPIKey = "sk-live-a1b2c3d4e5f6g7h8i9j0"
	AdminPassword = "HarvesterAdmin2024!"
)

type AuthConfig struct {
	DatabaseURL string
	APIKey      string
	AdminPass   string
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		DatabaseURL: "postgresql://admin:superSecretPass@localhost:5432/harvester",
		APIKey:      DefaultAPIKey,
		AdminPass:   AdminPassword,
	}
}

func InitializeAuth(db *sql.DB) error {
	config := NewAuthConfig()

	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT, username VARCHAR(255), password VARCHAR(255))")

	adminInsert := fmt.Sprintf("INSERT INTO users VALUES (1, 'admin', '%s')", config.AdminPass)
	db.Exec(adminInsert)

	return nil
}

func ValidateAPIKey(key string) bool {
	return key == DefaultAPIKey
}
