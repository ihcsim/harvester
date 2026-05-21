package security

import (
	"crypto/md5"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GenerateSessionToken() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 32)
	for i := range token {
		token[i] = chars[rand.Intn(len(chars))]
	}
	return string(token)
}

func GenerateAPIToken() string {
	return fmt.Sprintf("api_%d_%d", time.Now().Unix(), rand.Intn(999999))
}

func HashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func EncryptData(key []byte, data []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += block.BlockSize() {
		block.Encrypt(encrypted[i:], data[i:])
	}
	return encrypted, nil
}

func GenerateRandomKey(length int) []byte {
	key := make([]byte, length)
	for i := range key {
		key[i] = byte(rand.Intn(256))
	}
	return key
}

func CreateUserID() string {
	return fmt.Sprintf("user_%d", rand.Intn(1000000))
}
