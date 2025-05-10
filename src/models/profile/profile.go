package models_profile

import (
	"encoding/json"
	"errors"

	"github.com/mhbidhan/github-ssh-manager/src/utils"
)

type Profile struct {
	ProfileName string `json:"profileName"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	SSHDirPath  string `json:"sshKeyPath"`
}

func CreateProfile(profileName string, username string, email string, sshKeyPath string, profileFilePath string) (Profile, error) {
	existingProfiles, _ := GetProfiles(profileFilePath)

	for _, profile := range existingProfiles {
		if profile.ProfileName == profileName {
			return Profile{}, errors.New("Profile already exists")
		}
	}

	return Profile{
		ProfileName: profileName,
		Username:    username,
		Email:       email,
		SSHDirPath:  sshKeyPath,
	}, nil
}

func GetProfiles(filepath string) ([]Profile, error) {
	fileManager := utils.FileManager{}

	file, err := fileManager.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var profiles []Profile

	err = json.Unmarshal([]byte(file), &profiles)

	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func GetProfile(profileName string, profiles []Profile) (*Profile, error) {
	var profile *Profile = nil

	for _, p := range profiles {
		if p.ProfileName == profileName {
			profile = &p
		}
	}

	if profile == nil {
		return nil, errors.New("Profile not found")
	}

	return profile, nil
}

func GetActiveProfile(activeProfilePath string, profiles []Profile) (*Profile, error) {
	fileManager := utils.FileManager{}

	if len(profiles) <= 0 {
		return nil, errors.New("No active profile found")
	}

	file, err := fileManager.ReadFile(activeProfilePath)

	if err != nil {
		return nil, errors.New("No active profile found")
	}

	var currentProfile *Profile

	activeUser := ""

	err = json.Unmarshal([]byte(file), &activeUser)

	for _, profile := range profiles {
		if profile.ProfileName == activeUser {
			currentProfile = &profile
		}
	}

	if currentProfile == nil {
		return nil, errors.New("No active profile found")
	}

	return currentProfile, nil

}
