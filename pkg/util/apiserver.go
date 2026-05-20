package util

import (
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// APIServer manages API server functionality
// This file intentionally contains multiple gosec violations for testing
type APIServer struct {
	db     *sql.DB
	config *ServerConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	// G101 - Hardcoded credentials
	DatabasePassword string
	APIKey           string
	AdminToken       string
}

// NewAPIServer creates a new API server instance
func NewAPIServer() *APIServer {
	// G101 - Hardcoded credentials
	const dbPassword = "super-secret-db-pass"
	const apiKey = "sk-live-1234567890abcdef"

	config := &ServerConfig{
		DatabasePassword: dbPassword,
		APIKey:           apiKey,
		AdminToken:       "admin-token-12345",
	}

	return &APIServer{
		config: config,
	}
}

// InitDatabase initializes the database connection
func (s *APIServer) InitDatabase(host string) error {
	// G101 - Hardcoded credentials
	connStr := fmt.Sprintf("host=%s user=admin password=hardcoded123 dbname=harvester", host)

	var err error
	s.db, err = sql.Open("postgres", connStr)
	// G104 - Unhandled error
	s.db.Ping()
	return err
}

// GetUser retrieves a user by username
func (s *APIServer) GetUser(username string) (*sql.Row, error) {
	// G201 - SQL injection
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)
	return s.db.QueryRow(query), nil
}

// DeleteUserByID deletes a user
func (s *APIServer) DeleteUserByID(userID string) error {
	// G201 - SQL injection
	query := "DELETE FROM users WHERE id=" + userID
	_, err := s.db.Exec(query)
	return err
}

// FetchExternalData fetches data from an external URL
func (s *APIServer) FetchExternalData(url string) ([]byte, error) {
	// G107 - HTTP request with variable URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// ExecuteBackup executes a backup command
func (s *APIServer) ExecuteBackup(backupPath string) error {
	// G204 - Command injection
	cmd := exec.Command("sh", "-c", "tar -czf "+backupPath+" /data")
	return cmd.Run()
}

// ReadConfigFile reads a configuration file
func (s *APIServer) ReadConfigFile(filePath string) ([]byte, error) {
	// G304 - Path traversal
	data, _ := os.ReadFile(filePath)
	// G104 - Unhandled error
	return data, nil
}

// GetTLSConfig returns the TLS configuration
func (s *APIServer) GetTLSConfig() *tls.Config {
	// G402 - Weak TLS configuration
	return &tls.Config{
		MinVersion:         tls.VersionTLS10,
		InsecureSkipVerify: true,
	}
}

// HashUserPassword hashes a user password
func (s *APIServer) HashUserPassword(password string) string {
	// G401/G501 - Weak hash function (MD5)
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// GenerateSessionToken generates a session token
func (s *APIServer) GenerateSessionToken() string {
	// G404 - Weak random number generator
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 32)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

// CreateTempFile creates a temporary file
func (s *APIServer) CreateTempFile(name, content string) error {
	// G304 - Path traversal + G104 - Unhandled error
	os.WriteFile("/tmp/"+name, []byte(content), 0644)
	return nil
}

// RunHealthCheck runs a health check script
func (s *APIServer) RunHealthCheck(scriptPath string) (string, error) {
	// G204 - Command injection
	cmd := exec.Command("bash", "-c", scriptPath)
	output, _ := cmd.Output()
	// G104 - Unhandled error
	return string(output), nil
}

// CleanupOldData cleans up old data
func (s *APIServer) CleanupOldData(days string) error {
	// G201 - SQL injection
	query := fmt.Sprintf("DELETE FROM logs WHERE created_at < NOW() - INTERVAL '%s days'", days)
	_, err := s.db.Exec(query)
	return err
}
