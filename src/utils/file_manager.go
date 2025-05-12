package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileManager struct {
}

func (FileManager) WriteFile(path string, content any) error {

	data, err := json.Marshal(content)

	if err != nil {
		panic(err)
	}

	os.WriteFile(path, data, 0644)

	return nil
}

func (FileManager) ReadFile(path string) (string, error) {
	output, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (FileManager) CreateDir(path string) error {
	dir, err := os.Stat(path)

	if err != nil || !dir.IsDir() {
		err := os.Mkdir(path, 0755)

		if err != nil {
			return err
		}
	}

	return nil

}

func (FileManager) CopyFileToDirectory(srcFile, destinationDirectory string) error {
	fileName := filepath.Base(srcFile)
	destinationFile := filepath.Join(destinationDirectory, fileName)

	source, err := os.Open(srcFile)

	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := destination.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close destination file %s: %w", destinationFile, closeErr)
		}
	}()

	_, err = io.Copy(destination, source)
	if err != nil {
		os.Remove(destinationFile)
		return err
	}

	sourceStat, err := source.Stat()
	if err != nil {
		return err
	} else {
		err = os.Chmod(destinationFile, sourceStat.Mode())
		if err != nil {
			return err
		}
	}

	return err
}

func (FileManager) RenameFile(filePath string, newFileName string) error {
	dir := filepath.Dir(filePath)

	newPath := filepath.Join(dir, newFileName)

	err := os.Rename(filePath, newPath)
	if err != nil {
		return fmt.Errorf("failed to rename file from %s to %s: %w", filePath, newPath, err)
	}

	return nil
}
