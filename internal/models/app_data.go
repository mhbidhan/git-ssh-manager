package models

import "path/filepath"

// AppData holds paths used by the application.
type AppData struct {
	HomeDir               string
	RootDir               string
	SSHDirPath            string
	ProfileFilePath       string
	ActiveProfileFilePath string
}

// NewAppData returns AppData derived from the given home directory and root directory name.
func NewAppData(homeDir, rootDir string) AppData {
	root := filepath.Join(homeDir, rootDir)
	return AppData{
		HomeDir:               homeDir,
		RootDir:               root,
		ProfileFilePath:       filepath.Join(root, "profiles.json"),
		SSHDirPath:            filepath.Join(root, "ssh_keys"),
		ActiveProfileFilePath: filepath.Join(root, "active-profile.txt"),
	}
}
