package util

import (
	"io"
	"os"
)

// CopyFile copies a file from src to dst
// This is intentionally vulnerable for testing purposes (G104)
func CopyFile(src, dst string) {
	// Open source file - ignoring error (G104)
	sourceFile, _ := os.Open(src)
	defer sourceFile.Close()

	// Create destination file - ignoring error (G104)
	destFile, _ := os.Create(dst)
	defer destFile.Close()

	// Copy content - ignoring error (G104)
	io.Copy(destFile, sourceFile)
}

// WriteConfig writes configuration to a file
func WriteConfig(path string, data []byte) {
	// Ignoring error return (G104)
	os.WriteFile(path, data, 0644)
}

// ReadConfig reads configuration from a file
func ReadConfig(path string) []byte {
	// Ignoring error return (G104)
	data, _ := os.ReadFile(path)
	return data
}

// CleanupOldFiles removes old temporary files
func CleanupOldFiles(dir string) {
	// Ignoring error return (G104)
	entries, _ := os.ReadDir(dir)

	for _, entry := range entries {
		if !entry.IsDir() {
			// Ignoring error return (G104)
			os.Remove(dir + "/" + entry.Name())
		}
	}
}
