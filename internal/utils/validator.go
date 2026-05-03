package utils

import (
	"errors"
	"net/mail"
	"os"
	"strings"

	"github.com/mhbidhan/git-ssh-manager/internal/constants"
)

// ValidateInput validates command line arguments and returns the flag, profile name, and any error.
func ValidateInput() (string, string, error) {
	if len(os.Args) < 2 {
		PrintErrorMessage("argument is required")
		PrintAllCommands()
		return "", "", errors.New("argument is required")
	}

	flag := os.Args[1]

	if constants.AppFlags[flag] == "--list" ||
		constants.AppFlags[flag] == "--status" ||
		constants.AppFlags[flag] == "--help" {
		return flag, "", nil
	}

	if _, valid := constants.AppFlags[flag]; !valid {
		PrintErrorMessage("invalid argument")
		PrintAllCommands()
		return "", "", errors.New("invalid argument")
	}

	if len(os.Args) < 3 {
		PrintErrorMessage("profile name is required")
		return "", "", errors.New("profile name is required")
	}

	return flag, os.Args[2], nil
}

// ValidateEmail checks if the provided string is a valid email address.
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("email cannot be empty")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}

	return nil
}
