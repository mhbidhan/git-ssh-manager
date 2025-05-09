package utils

import (
	"encoding/json"
	"os"
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
