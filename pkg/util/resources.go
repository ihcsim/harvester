package util

import (
	"io"
	"os"
)

// ResourceManager manages system resources
type ResourceManager struct{}

// NewResourceManager creates a new ResourceManager instance
func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

// ProcessFile processes a file
// This is intentionally vulnerable for testing purposes (G307)
func (r *ResourceManager) ProcessFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	// G307 - defer on a function that returns error without checking
	defer file.Close()

	// Process file
	_, err = io.ReadAll(file)
	return err
}

// WriteToFile writes data to a file
func WriteToFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// G307 - defer Close() without checking error
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// CopyFileContent copies file content
func CopyFileContent(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// G307 - unchecked error on deferred Close
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	// G307 - unchecked error on deferred Close
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// ReadFileContent reads file content
func (r *ResourceManager) ReadFileContent(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// G307 - deferred Close without error check
	defer f.Close()

	return io.ReadAll(f)
}

// AppendToFile appends data to a file
func AppendToFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// G307 - deferred Close error not checked
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// SyncFile ensures file is written to disk
func SyncFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// G307 - deferred Close error not checked
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	// Sync to disk but ignore error
	file.Sync()
	return nil
}

// TruncateFile truncates a file to size
func TruncateFile(path string, size int64) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	// G307 - deferred Close error not checked
	defer file.Close()

	return file.Truncate(size)
}
