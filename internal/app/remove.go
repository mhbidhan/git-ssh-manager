package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bobg/go-generics/slices"
	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func removeProfile(filePaths models.AppData, profileName string) error {
	fileManager := utils.NewFileManager()

	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return fmt.Errorf("failed to get profiles: %w", err)
	}

	prof, err := models.GetProfile(profileName, profiles)
	if err != nil {
		return err
	}

	file, err := fileManager.ReadFile(filePaths.ActiveProfileFilePath)
	if err != nil {
		return errors.New("no active profile found")
	}

	var activeUser string
	if err := json.Unmarshal([]byte(file), &activeUser); err != nil {
		return errors.New("no active profile found")
	}

	if prof.ProfileName == activeUser {
		return errors.New("profile is currently in use")
	}

	filtered, err := slices.Filter(profiles, func(val models.Profile) (bool, error) {
		return val.ProfileName != profileName, nil
	})
	if err != nil {
		return fmt.Errorf("failed to filter profiles: %w", err)
	}

	if err := fileManager.WriteFile(filePaths.ProfileFilePath, filtered); err != nil {
		return err
	}

	sshProfileDirPath := profileName + "_github"
	dir := filepath.Join(filePaths.SSHDirPath, sshProfileDirPath)

	_ = os.RemoveAll(dir)

	return nil
}
