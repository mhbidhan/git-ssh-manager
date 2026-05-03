package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Profile represents a git SSH profile.
type Profile struct {
	ProfileName string `json:"profileName"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	SSHDirPath  string `json:"sshDirPath"`
}

// PrintProfileInfo prints profile information to stdout.
func (p Profile) PrintProfileInfo() {
	fmt.Print("\033[32m", "\n"+p.ProfileName, "\033[0m")

	fmt.Printf(`
------------------------
Username: %s
Email: %s
`, strings.Trim(p.Username, "\n"), strings.Trim(p.Email, "\n"))
}

// CreateProfile creates a new profile after validating it doesn't already exist.
func CreateProfile(profileName string, username string, email string, sshDirPath string, profileFilePath string) (Profile, error) {
	existingProfiles, err := GetProfiles(profileFilePath)
	if err != nil {
		return Profile{}, fmt.Errorf("failed to get existing profiles: %w", err)
	}

	for _, p := range existingProfiles {
		if p.ProfileName == profileName {
			return Profile{}, errors.New("profile already exists")
		}
	}

	return Profile{
		ProfileName: profileName,
		Username:    username,
		Email:       email,
		SSHDirPath:  sshDirPath,
	}, nil
}

// GetProfiles reads and parses profiles from the given file path.
func GetProfiles(path string) ([]Profile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, fmt.Errorf("failed to unmarshal profiles: %w", err)
	}

	return profiles, nil
}

// GetProfile finds a profile by name from the provided slice.
func GetProfile(profileName string, profiles []Profile) (*Profile, error) {
	for i := range profiles {
		if profiles[i].ProfileName == profileName {
			return &profiles[i], nil
		}
	}

	return nil, errors.New("profile not found")
}

// GetActiveProfile returns the currently active profile.
func GetActiveProfile(activeProfilePath string, profiles []Profile) (*Profile, error) {
	if len(profiles) == 0 {
		return nil, errors.New("no active profile found")
	}

	data, err := os.ReadFile(activeProfilePath)
	if err != nil {
		return nil, errors.New("no active profile found")
	}

	var activeUser string
	if err := json.Unmarshal(data, &activeUser); err != nil {
		return nil, errors.New("no active profile found")
	}

	for i := range profiles {
		if profiles[i].ProfileName == activeUser {
			return &profiles[i], nil
		}
	}

	return nil, errors.New("no active profile found")
}
