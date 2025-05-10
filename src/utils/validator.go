package utils

import (
	"errors"
	"os"

	"github.com/mhbidhan/github-ssh-manager/src/constants"
	"github.com/mhbidhan/github-ssh-manager/src/interfaces"
)

func ValidateInput(githubSshManager interfaces.GithubSshManager) (string, string, error) {
	if len(os.Args) < 2 {
		PrintErrorMessage("Argument is required")
		PrintAllCommands()
		return "", "", errors.New("An error occured")
	}

	flag := os.Args[1]

	if constants.AppFlags[flag] == "--status" {
		return flag, "", nil
	}

	if constants.AppFlags[flag] == "--help" {
		PrintAllCommands()
		return "", "", errors.New("An error occured")
	}

	if constants.AppFlags[flag] == "--status" || constants.AppFlags[flag] == "--help" {
		return "", "", errors.New("An error occured")
	}

	if _, valid := constants.AppFlags[flag]; !valid {
		PrintErrorMessage("Invalid argument")

		PrintAllCommands()

		return "", "", errors.New("Invalid argument")
	}

	if len(os.Args) < 3 {
		PrintErrorMessage("Profile name is required")
		return "", "", errors.New("An error occured")
	}

	profileName := os.Args[2]

	return flag, profileName, nil

}
