package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// GetHomeDir returns the user's home directory.
func GetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home
}

// GetRootDir returns the root directory for the given platform.
func GetRootDir(platform string) (string, error) {
	switch platform {
	case "linux":
		return filepath.Join(".config", "git-ssh-manager"), nil
	case "darwin":
		return filepath.Join(".config", "git-ssh-manager"), nil
	case "windows":
		return filepath.Join("AppData", "Local", "git-ssh-manager"), nil
	default:
		return "", errors.New("platform not supported")
	}
}
