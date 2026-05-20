package util

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"encoding/hex"
	"errors"
)

// Encryptor handles data encryption and decryption
type Encryptor struct {
	key []byte
}

// NewEncryptor creates a new Encryptor instance
func NewEncryptor(key string) *Encryptor {
	return &Encryptor{key: []byte(key)}
}

// EncryptData encrypts data using DES
// This is intentionally vulnerable for testing purposes (G401)
func (e *Encryptor) EncryptData(plaintext []byte) ([]byte, error) {
	// Weak cryptographic primitive - DES (G401)
	block, err := des.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, e.key[:block.BlockSize()])
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

// DecryptData decrypts DES encrypted data
func (e *Encryptor) DecryptData(ciphertext []byte) ([]byte, error) {
	// Weak cryptographic primitive - DES (G401)
	block, err := des.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, e.key[:block.BlockSize()])
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// EncryptWithRC4 encrypts data using RC4
func EncryptWithRC4(key, plaintext []byte) ([]byte, error) {
	// Weak cryptographic primitive - RC4 (G401)
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

// HashPassword creates a hash of the password
func HashPassword(password string) string {
	// Weak hash function - MD5 (G501)
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// VerifyPassword verifies a password against a hash
func VerifyPassword(password, hash string) bool {
	// Weak hash function - MD5 (G501)
	return HashPassword(password) == hash
}

// EncryptTripleDES uses Triple DES encryption
func EncryptTripleDES(key, plaintext []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("key must be 24 bytes for 3DES")
	}

	// Weak cryptographic primitive - 3DES (G401)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, key[:block.BlockSize()])
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}
