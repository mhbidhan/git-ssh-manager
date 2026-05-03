package app

import (
	"github.com/mhbidhan/git-ssh-manager/internal/models"
)

func list(filePaths models.AppData) error {
	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	for _, prof := range profiles {
		prof.PrintProfileInfo()
	}

	return nil
}
