package models_file_paths

import "path/filepath"

type FilePaths struct {
	HomeDir               string
	RootDir               string
	SSHDirPath            string
	ProfileFilePath       string
	ActiveProfileFilePath string
}

func GenerateFilePaths(homeDir string, rootDir string) FilePaths {
	return FilePaths{

		HomeDir:               homeDir,
		RootDir:               filepath.Join(homeDir, rootDir),
		ProfileFilePath:       filepath.Join(homeDir, rootDir, "/profiles.json"),
		SSHDirPath:            filepath.Join(homeDir, rootDir, "/ssh_keys"),
		ActiveProfileFilePath: filepath.Join(homeDir, rootDir, "/active-profile.txt"),
	}
}
