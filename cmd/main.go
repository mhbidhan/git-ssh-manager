package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mhbidhan/git-ssh-manager/internal/app"
	"github.com/mhbidhan/git-ssh-manager/internal/constants"
	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func main() {
	homeDir := utils.GetHomeDir()
	rootDir, err := utils.GetRootDir(runtime.GOOS)
	if err != nil {
		utils.PrintErrorMessage(err.Error())
		return
	}

	appData := models.NewAppData(homeDir, rootDir)
	fileManager := utils.NewFileManager()
	githubSshManager := app.NewGitSSHManager(appData)

	if err := fileManager.CreateDir(appData.RootDir); err != nil {
		utils.PrintErrorMessage(fmt.Sprintf("failed to create root dir: %v", err))
	}

	if err := fileManager.CreateDir(appData.SSHDirPath); err != nil {
		utils.PrintErrorMessage(fmt.Sprintf("failed to create SSH dir: %v", err))
	}

	if _, err := os.Stat(appData.ProfileFilePath); err != nil {
		if err := fileManager.WriteFile(appData.ProfileFilePath, nil); err != nil {
			utils.PrintErrorMessage(fmt.Sprintf("failed to create profile file: %v", err))
		}
	}

	if _, err := os.Stat(appData.ActiveProfileFilePath); err != nil {
		if err := fileManager.WriteFile(appData.ActiveProfileFilePath, nil); err != nil {
			utils.PrintErrorMessage(fmt.Sprintf("failed to create active profile file: %v", err))
		}
	}

	flag, profileName, err := utils.ValidateInput()
	if err != nil {
		return
	}

	switch constants.AppFlags[flag] {
	case "--list":
		if err := githubSshManager.List(); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--status":
		if err := githubSshManager.Status(); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--get-key":
		if err := githubSshManager.GetKey(profileName); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--new":
		if err := githubSshManager.New(profileName); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--use":
		if err := githubSshManager.Use(profileName); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--remove":
		if err := githubSshManager.Remove(profileName); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--help":
		utils.PrintAllCommands()
	case "--bind-profile":
		if err := githubSshManager.BindProfile(profileName); err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	default:
		utils.PrintErrorMessage("Invalid command")
		utils.PrintAllCommands()
	}
}
