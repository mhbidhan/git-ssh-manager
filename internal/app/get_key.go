package app

import (
	"fmt"
	"path/filepath"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
)

func getKey(filePaths models.AppData, profileName string) error {
	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	prof, err := models.GetProfile(profileName, profiles)
	if err != nil {
		return err
	}

	fmt.Println(filepath.Join(filePaths.SSHDirPath, prof.SSHDirPath, prof.SSHDirPath+".pub"))
	return nil
}
