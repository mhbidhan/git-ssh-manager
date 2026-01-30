package utils

import (
	"errors"
	"net/mail"
	"os"
	"strings"

	"github.com/mhbidhan/git-ssh-manager/src/constants"
	"github.com/mhbidhan/git-ssh-manager/src/interfaces"
)

func ValidateInput(githubSshManager interfaces.GithubSshManager) (string, string, error) {
	if len(os.Args) < 2 {
		PrintErrorMessage("Argument is required")
		PrintAllCommands()

		return "", "", errors.New("An error occured")
	}

	flag := os.Args[1]

	if constants.AppFlags[flag] == "--list" {
		return flag, "", nil
	}

	if constants.AppFlags[flag] == "--status" {
		return flag, "", nil
	}

	if constants.AppFlags[flag] == "--help" {
		return flag, "", nil
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

// ValidateEmail checks if the provided string is a valid email address
func ValidateEmail(email string) error {
	// Trim whitespace and newlines
	email = strings.TrimSpace(email)
	email = strings.Trim(email, "\n\r")

	if email == "" {
		return errors.New("email cannot be empty")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email format")
	}

	return nil
}
