package app_git_ssh_manager

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/bobg/go-generics/slices"
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
	"github.com/mhbidhan/git-ssh-manager/src/utils"
)

func remove(filePaths models_file_paths.FilePaths, profileName string) error {
	fileManager := utils.FileManager{}

	profiles, _ := models_profile.GetProfiles(filePaths.ProfileFilePath)

	profile, err := models_profile.GetProfile(profileName, profiles)

	if err != nil {
		return err
	}

	file, err := fileManager.ReadFile(filePaths.ActiveProfileFilePath)

	if err != nil {
		return errors.New("No active profile found")
	}

	activeUser := ""

	err = json.Unmarshal([]byte(file), &activeUser)

	if profile.ProfileName == activeUser {
		return errors.New("Profile is currently in use")
	}

	filtered, err := slices.Filter(profiles, func(val models_profile.Profile) (bool, error) { return val.ProfileName != profileName, nil })

	if err != nil {
		return err
	}

	fileManager.WriteFile(filePaths.ProfileFilePath, filtered)

	sshProfileDirPath := profileName + "_github"

	dir := filepath.Join(filePaths.SSHDirPath, sshProfileDirPath)

	os.Remove(dir)

	return nil
}
