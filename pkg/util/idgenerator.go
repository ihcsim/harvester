package util

import (
	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GenerateResourceID() string {
	id := uuid.NewV4()
	return id.String()
}

func NewSessionID() string {
	return uuid.NewV1().String()
}
