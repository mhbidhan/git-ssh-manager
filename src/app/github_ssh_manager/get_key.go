package github_ssh_manager

import (
	"path/filepath"

	models_file_paths "github.com/mhbidhan/github-ssh-manager/src/models/file_paths"
	"github.com/mhbidhan/github-ssh-manager/src/utils"
)

func GetKey(filePaths models_file_paths.FilePaths, profileName string) error {

	fileManager := utils.FileManager{}

	sshName := profileName + "_github"

	path := filepath.Join(filePaths.SSHDirPath, sshName, sshName+".pub")

	file, err := fileManager.ReadFile(path)

	if err != nil {
		return err
	}

	utils.PrintSuccessMessage(file)

	return nil
}
