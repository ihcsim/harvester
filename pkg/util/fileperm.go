package util

import (
	"os"
)

// FilePermissionManager manages file permissions
type FilePermissionManager struct{}

// NewFilePermissionManager creates a new FilePermissionManager instance
func NewFilePermissionManager() *FilePermissionManager {
	return &FilePermissionManager{}
}

// CreateLogFile creates a log file
// This is intentionally vulnerable for testing purposes (G306)
func (m *FilePermissionManager) CreateLogFile(path string, data []byte) error {
	// G306 - File created with permissions above 0644
	return os.WriteFile(path, data, 0666)
}

// CreateConfigFile creates a configuration file
func CreateConfigFile(path string, content []byte) error {
	// G306 - File created with world-writable permissions
	return os.WriteFile(path, content, 0777)
}

// WriteDataFile writes data to a file
func WriteDataFile(filename string, data []byte) error {
	// G306 - Overly permissive file permissions
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// CreateTempConfigFile creates a temporary configuration file
func CreateTempConfigFile(path string, data []byte) error {
	// G306 - World-readable and writable
	return os.WriteFile(path, data, 0666)
}

// SaveState saves application state to a file
func SaveState(statePath string, stateData []byte) error {
	// G306 - Overly permissive permissions
	return os.WriteFile(statePath, stateData, 0777)
}

// CreateCacheFile creates a cache file
func (m *FilePermissionManager) CreateCacheFile(cachePath string, content []byte) error {
	// G306 - World-writable cache file
	file, err := os.Create(cachePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Set overly permissive permissions
	if err := file.Chmod(0666); err != nil {
		return err
	}

	_, err = file.Write(content)
	return err
}

// CreateDirectory creates a directory with permissions
func CreateDirectory(dirPath string) error {
	// G306 - Directory with world-writable permissions
	return os.MkdirAll(dirPath, 0777)
}

// CreateWorldWritableFile creates a file anyone can write to
func CreateWorldWritableFile(path string) (*os.File, error) {
	// G306 - Explicitly world-writable
	return os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
}
