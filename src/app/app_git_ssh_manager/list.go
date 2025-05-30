package app_git_ssh_manager

import (
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
)

func list(filePaths models_file_paths.FilePaths) error {
	profiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)

	if err != nil {
		return err
	}

	for _, profile := range profiles {
		profile.PrintProfileInfo()
	}

	return nil
}
