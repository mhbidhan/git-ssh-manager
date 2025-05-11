package utils

import (
	"os"

	"github.com/mhbidhan/git-ssh-manager/src/interfaces"
)

func Bootstrap(githubSshManager interfaces.GithubSshManager) {
	fileManager := FileManager{}
	filepath := githubSshManager.GetFilePaths()

	err := fileManager.CreateDir(filepath.RootDir)

	if err != nil {
		PrintErrorMessage(err.Error())
	}

	err = fileManager.CreateDir(filepath.SSHDirPath)

	if err != nil {
		PrintErrorMessage(err.Error())
	}
	var profileFile any

	profileFile, err = os.Stat(filepath.ProfileFilePath)

	if profileFile == nil {
		err = fileManager.WriteFile(filepath.ProfileFilePath, nil)

		if err != nil {
			PrintErrorMessage(err.Error())
		}
	}

	var activeProfileFile any

	activeProfileFile, err = os.Stat(filepath.ActiveProfileFilePath)

	if activeProfileFile == nil {
		err = fileManager.WriteFile(filepath.ActiveProfileFilePath, nil)

		if err != nil {
			PrintErrorMessage(err.Error())
		}
	}
}
