package main

import (
	"runtime"

	"github.com/mhbidhan/github-ssh-manager/src/constants"
	"github.com/mhbidhan/github-ssh-manager/src/factories"
	"github.com/mhbidhan/github-ssh-manager/src/utils"
)

func main() {
	githubSshManager, err := factories.CreateGithubSshManager(runtime.GOOS)

	if err != nil {
		utils.PrintErrorMessage(err.Error())
		return
	}

	utils.Bootstrap(githubSshManager)

	flag, profileName, err := utils.ValidateInput(githubSshManager)

	if err != nil {
		return
	}

	switch constants.AppFlags[flag] {
	case "--status":
		err := githubSshManager.Status()
		if err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--get-key":
		err := githubSshManager.GetKey(profileName)
		if err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--new":
		err := githubSshManager.New(profileName)
		if err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--use":
		err := githubSshManager.Use(profileName)
		if err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	case "--remove":
		err := githubSshManager.Remove(profileName)
		if err != nil {
			utils.PrintErrorMessage(err.Error())
		}
	}

}
