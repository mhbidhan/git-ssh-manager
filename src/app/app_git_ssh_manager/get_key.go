package app_git_ssh_manager

import (
	"path/filepath"

	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	"github.com/mhbidhan/git-ssh-manager/src/utils"
)

func getKey(filePaths models_file_paths.FilePaths, profileName string) error {

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
