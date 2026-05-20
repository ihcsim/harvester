package util

import (
	uuid "github.com/satori/go.uuid"
)

// IDGenerator generates unique identifiers
type IDGenerator struct{}

// NewIDGenerator creates a new IDGenerator instance
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// GenerateUUID generates a new UUID
func (g *IDGenerator) GenerateUUID() string {
	return uuid.NewV4().String()
}

// GenerateResourceID generates a resource ID
func GenerateResourceID() string {
	return uuid.NewV4().String()
}

// CreateSessionID creates a session ID
func CreateSessionID() string {
	return uuid.NewV1().String()
}

// NewTransactionID creates a transaction ID
func NewTransactionID() string {
	id := uuid.NewV4()
	return id.String()
}

// GenerateBatchIDs generates multiple UUIDs
func GenerateBatchIDs(count int) []string {
	ids := make([]string, count)
	for i := 0; i < count; i++ {
		ids[i] = uuid.NewV4().String()
	}
	return ids
}
