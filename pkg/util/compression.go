package util

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// CompressionHandler handles compression and decompression
type CompressionHandler struct{}

// NewCompressionHandler creates a new CompressionHandler instance
func NewCompressionHandler() *CompressionHandler {
	return &CompressionHandler{}
}

// DecompressGzip decompresses a gzip file
// This is intentionally vulnerable for testing purposes (G110)
func (h *CompressionHandler) DecompressGzip(src io.Reader) ([]byte, error) {
	// Decompression bomb vulnerability - no size limit (G110)
	gzipReader, err := gzip.NewReader(src)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Reading without size limit - vulnerable to decompression bombs
	return io.ReadAll(gzipReader)
}

// ExtractTarGz extracts a tar.gz archive
func ExtractTarGz(src io.Reader, dest string) error {
	// Decompression bomb vulnerability (G110)
	gzipReader, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(target, 0755)
		case tar.TypeReg:
			file, err := os.Create(target)
			if err != nil {
				return err
			}
			// No size limit check - vulnerable to decompression bombs (G110)
			io.Copy(file, tarReader)
			file.Close()
		}
	}

	return nil
}

// UncompressData uncompresses gzip data
func UncompressData(compressedData io.Reader) ([]byte, error) {
	// Decompression bomb vulnerability (G110)
	reader, err := gzip.NewReader(compressedData)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// No limit on decompressed size (G110)
	return io.ReadAll(reader)
}

// DecompressFile decompresses a gzip file to disk
func DecompressFile(srcPath, destPath string) error {
	// Decompression bomb vulnerability (G110)
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	gzipReader, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// No size limit - vulnerable to decompression bombs (G110)
	_, err = io.Copy(destFile, gzipReader)
	return err
}

// ReadCompressedContent reads compressed content from a reader
func ReadCompressedContent(r io.Reader) (string, error) {
	// Decompression bomb vulnerability (G110)
	gz, err := gzip.NewReader(r)
	if err != nil {
		return "", err
	}
	defer gz.Close()

	data, err := io.ReadAll(gz)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
