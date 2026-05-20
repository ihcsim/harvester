package util

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// FileServer handles file serving operations
type FileServer struct {
	baseDir string
}

// NewFileServer creates a new FileServer instance
func NewFileServer(baseDir string) *FileServer {
	return &FileServer{baseDir: baseDir}
}

// ServeFile serves a file based on the requested path
// This is intentionally vulnerable for testing purposes (G304)
func (fs *FileServer) ServeFile(w http.ResponseWriter, r *http.Request) {
	// Path traversal vulnerability - using user input directly (G304)
	filename := r.URL.Query().Get("file")
	data, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Write(data)
}

// DownloadFile downloads a file from the server
func (fs *FileServer) DownloadFile(w http.ResponseWriter, r *http.Request) {
	// Path traversal vulnerability (G304)
	filepath := r.URL.Query().Get("path")
	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, "Cannot open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}

// ReadUserFile reads a file from user-specified path
func ReadUserFile(userPath string) ([]byte, error) {
	// Path traversal vulnerability (G304)
	return os.ReadFile(userPath)
}

// LoadTemplate loads a template file
func (fs *FileServer) LoadTemplate(templateName string) (string, error) {
	// Path traversal vulnerability (G304)
	path := filepath.Join(fs.baseDir, templateName)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetFileContent retrieves file content from a path parameter
func GetFileContent(r *http.Request) []byte {
	// Path traversal vulnerability (G304)
	filePath := r.FormValue("filepath")
	content, _ := os.ReadFile(filePath)
	return content
}
