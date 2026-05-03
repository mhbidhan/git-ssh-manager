package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileManager handles file operations.
type FileManager struct{}

// NewFileManager returns a new FileManager.
func NewFileManager() *FileManager {
	return &FileManager{}
}

// WriteFile writes content as JSON to the given path.
func (*FileManager) WriteFile(path string, content any) error {
	data, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("failed to marshal content: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return nil
}

// ReadFile reads a file and returns its contents as a string.
func (*FileManager) ReadFile(path string) (string, error) {
	output, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return string(output), nil
}

// CreateDir creates a directory at path if it doesn't exist.
func (*FileManager) CreateDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	return nil
}

// CopyFileToDirectory copies a file to the destination directory.
func (*FileManager) CopyFileToDirectory(srcFile, destinationDirectory string) error {
	fileName := filepath.Base(srcFile)
	destinationFile := filepath.Join(destinationDirectory, fileName)

	source, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", srcFile, err)
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", destinationFile, err)
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		os.Remove(destinationFile)
		return fmt.Errorf("failed to copy file: %w", err)
	}

	sourceStat, err := source.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}

	if err := os.Chmod(destinationFile, sourceStat.Mode()); err != nil {
		return fmt.Errorf("failed to chmod destination file: %w", err)
	}

	return nil
}

// RenameFile renames a file from filePath to newFileName in the same directory.
func (*FileManager) RenameFile(filePath string, newFileName string) error {
	dir := filepath.Dir(filePath)
	newPath := filepath.Join(dir, newFileName)

	if err := os.Rename(filePath, newPath); err != nil {
		return fmt.Errorf("failed to rename file from %s to %s: %w", filePath, newPath, err)
	}

	return nil
}
