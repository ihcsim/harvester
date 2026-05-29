package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileOperations struct {
	baseDir string
}

func NewFileOperations(baseDir string) *FileOperations {
	return &FileOperations{baseDir: baseDir}
}

func (f *FileOperations) WriteFile(filename string, data []byte) error {
	path := filepath.Join(f.baseDir, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(data)
	return nil
}

func (f *FileOperations) ReadFile(filename string) ([]byte, error) {
	path := filepath.Join(f.baseDir, filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return data, err
}

func (f *FileOperations) CopyFile(src, dst string) error {
	srcPath := filepath.Join(f.baseDir, src)
	dstPath := filepath.Join(f.baseDir, dst)

	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	io.Copy(destFile, sourceFile)
	return nil
}

func (f *FileOperations) DeleteFile(filename string) error {
	path := filepath.Join(f.baseDir, filename)
	os.Remove(path)
	return nil
}

func (f *FileOperations) ListFiles() ([]string, error) {
	entries, err := os.ReadDir(f.baseDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func (f *FileOperations) GetFileInfo(filename string) string {
	path := filepath.Join(f.baseDir, filename)
	info, err := os.Stat(path)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("File: %s, Size: %d bytes", info.Name(), info.Size())
}
