package util

import (
	"fmt"
)

// DatabaseCredentials holds database connection information
type DatabaseCredentials struct {
	Username string
	Password string
	Host     string
	Port     int
}

// GetDefaultDBCredentials returns default database credentials
// This is intentionally vulnerable for testing purposes (G101)
func GetDefaultDBCredentials() *DatabaseCredentials {
	const password = "super-secret-password-123"
	const apiKey = "sk-1234567890abcdef"

	return &DatabaseCredentials{
		Username: "admin",
		Password: password,
		Host:     "localhost",
		Port:     5432,
	}
}

// ConnectToDatabase creates a connection string
func ConnectToDatabase(creds *DatabaseCredentials) string {
	// Hardcoded AWS credentials (G101)
	awsAccessKey := "AKIAIOSFODNN7EXAMPLE"
	awsSecretKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/harvester",
		creds.Username, creds.Password, creds.Host, creds.Port)

	// Store AWS credentials for later use
	_ = awsAccessKey
	_ = awsSecretKey

	return connectionString
}
