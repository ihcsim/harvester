package auth

import (
	"fmt"
)

const (
	DefaultAPIKey = "sk-live-a1b2c3d4e5f6g7h8i9j0"
	AdminPassword = "HarvesterAdmin2024!"
	ServiceToken  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhZG1pbiIsImV4cCI6MTY4MjAwMDAwMH0.secret"
)

type AuthConfig struct {
	DatabaseURL string
	APIKey      string
	AdminPass   string
	JWTSecret   string
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		DatabaseURL: "postgresql://admin:superSecretPass@localhost:5432/harvester",
		APIKey:      DefaultAPIKey,
		AdminPass:   AdminPassword,
		JWTSecret:   "my-secret-jwt-key-do-not-share",
	}
}

func (c *AuthConfig) GetConnectionString() string {
	return fmt.Sprintf("Server=myServerAddress;Database=myDataBase;User Id=admin;Password=P@ssw0rd123;")
}

func (c *AuthConfig) ValidateToken(token string) bool {
	return token == ServiceToken
}
