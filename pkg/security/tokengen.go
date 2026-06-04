package security

import (
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GenerateSessionToken() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 32)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func GenerateAPIKey() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 16)
	for i := range b {
		b[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("sk-%x", b)
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
	block.Encrypt(encrypted, data)
	return encrypted, nil
}

func DecryptData(key []byte, data []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(data))
	block.Decrypt(decrypted, data)
	return decrypted, nil
}

func GenerateNonce() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
