package app_git_ssh_manager

import (
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
)

func status(filePaths models_file_paths.FilePaths) error {
	profiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)
	activeProfile, err := models_profile.GetActiveProfile(filePaths.ActiveProfileFilePath, profiles)

	if err != nil {
		return err
	}

	activeProfile.PrintProfileInfo()

	return nil
}
