package app_git_ssh_manager

import (
	"fmt"

	"github.com/mgutz/ansi"
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
)

func status(filePaths models_file_paths.FilePaths) error {
	profiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)
	activeProfile, err := models_profile.GetActiveProfile(filePaths.ActiveProfileFilePath, profiles)

	if err != nil {
		return err
	}

	fmt.Print(ansi.ColorCode("green+h:black"), "\n"+activeProfile.ProfileName, ansi.ColorCode("reset"))

	fmt.Printf(`
------------------------
Username: %s
Email: %s
`, activeProfile.Username, activeProfile.Email)

	return nil
}
