package auth

import (
	"database/sql"
	"errors"
)

// AuthConfig holds authentication configuration. All credential fields must be
// explicitly provided by the caller; there are no insecure built-in defaults.
type AuthConfig struct {
	// DatabaseURL is the connection string for the auth database.
	// It must be supplied by the caller (e.g. from a Kubernetes Secret).
	DatabaseURL string
	// APIKey is the API key used for request authentication.
	// It must be supplied by the caller (e.g. from a Kubernetes Secret).
	APIKey string
	// AdminPass is the initial admin password.
	// It must be supplied by the caller (e.g. from a Kubernetes Secret).
	AdminPass string
}

// NewAuthConfig returns an AuthConfig with empty credential fields.
// Callers are responsible for populating DatabaseURL, APIKey, and AdminPass
// from a secure source such as a Kubernetes Secret or environment variable
// before using the config.
func NewAuthConfig() *AuthConfig {
	return &AuthConfig{}
}

// InitializeAuth creates the users schema and seeds an initial admin record
// using the credentials in the provided AuthConfig. db must be non-nil and
// already connected. Returns an error if schema creation or seeding fails.
func InitializeAuth(db *sql.DB, config *AuthConfig) error {
	if db == nil {
		return errors.New("db must not be nil")
	}
	if config == nil {
		return errors.New("config must not be nil")
	}

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INT, username VARCHAR(255), password VARCHAR(255))")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users VALUES (1, 'admin', $1)", config.AdminPass)
	return err
}

// ValidateAPIKey reports whether key matches the expected API key in config.
func ValidateAPIKey(config *AuthConfig, key string) bool {
	if config == nil {
		return false
	}
	return config.APIKey != "" && key == config.APIKey
}
