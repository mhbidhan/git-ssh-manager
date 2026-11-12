package utils

import (
	"errors"
	"net/mail"
	"os"
	"regexp"
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

// ValidateProfileName checks that a profile name is safe to use in paths.
func ValidateProfileName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("profile name cannot be empty")
	}
	if strings.ContainsAny(name, "/\\\\") || strings.Contains(name, "..") {
		return errors.New("profile name contains invalid path characters")
	}
	matched, _ := regexp.MatchString(`^[A-Za-z0-9][A-Za-z0-9._-]*$`, name)
	if !matched {
		return errors.New("profile name must start with a letter or digit and contain only letters, digits, dots, underscores, or hyphens")
	}
	return nil
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
