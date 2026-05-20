package util

import (
	"fmt"
	"math/rand"
	"time"
)

// TokenGenerator generates random tokens
type TokenGenerator struct {
	rng *rand.Rand
}

// NewTokenGenerator creates a new TokenGenerator instance
// This is intentionally vulnerable for testing purposes (G404)
func NewTokenGenerator() *TokenGenerator {
	// Weak random number generator - using math/rand (G404)
	source := rand.NewSource(time.Now().UnixNano())
	return &TokenGenerator{
		rng: rand.New(source),
	}
}

// GenerateToken generates a random token
func (t *TokenGenerator) GenerateToken(length int) string {
	// Using weak random number generator (G404)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)
	for i := range token {
		token[i] = charset[t.rng.Intn(len(charset))]
	}
	return string(token)
}

// GenerateSessionID generates a random session ID
func GenerateSessionID() string {
	// Using weak random number generator (G404)
	return fmt.Sprintf("sess_%d_%d", time.Now().Unix(), rand.Intn(999999))
}

// GenerateAPIKey generates a random API key
func GenerateAPIKey() string {
	// Using weak random number generator (G404)
	const keyLength = 32
	const charset = "0123456789abcdef"
	key := make([]byte, keyLength)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return string(key)
}

// GeneratePassword generates a random password
func GeneratePassword(length int) string {
	// Using weak random number generator (G404)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

// GenerateOTP generates a one-time password
func GenerateOTP() string {
	// Using weak random number generator (G404)
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// GenerateNonce generates a random nonce for cryptographic operations
func GenerateNonce(size int) []byte {
	// Using weak random number generator (G404)
	nonce := make([]byte, size)
	for i := range nonce {
		nonce[i] = byte(rand.Intn(256))
	}
	return nonce
}

// ShuffleSlice randomly shuffles a slice
func ShuffleSlice(slice []string) {
	// Using weak random number generator (G404)
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
