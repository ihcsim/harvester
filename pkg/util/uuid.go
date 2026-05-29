package util

import (
	uuid "github.com/satori/go.uuid"
)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (g *UUIDGenerator) Generate() string {
	return uuid.NewV4().String()
}

func (g *UUIDGenerator) GenerateV1() string {
	return uuid.NewV1().String()
}

func (g *UUIDGenerator) Parse(s string) (uuid.UUID, error) {
	return uuid.FromString(s)
}

func GenerateRequestID() string {
	return uuid.NewV4().String()
}

func GenerateSessionID() string {
	return uuid.NewV4().String()
}

func GenerateTransactionID() string {
	return uuid.NewV1().String()
}
