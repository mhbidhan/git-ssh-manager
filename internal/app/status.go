package app

import (
	"github.com/mhbidhan/git-ssh-manager/internal/models"
)

func status(filePaths models.AppData) error {
	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	activeProfile, err := models.GetActiveProfile(filePaths.ActiveProfileFilePath, profiles)
	if err != nil {
		return err
	}

	activeProfile.PrintProfileInfo()

	return nil
}
