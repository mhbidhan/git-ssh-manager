package github_ssh_manager

import (
	"fmt"
	"os/exec"
	"path/filepath"

	models_file_paths "github.com/mhbidhan/github-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/github-ssh-manager/src/models/profile"
	"github.com/mhbidhan/github-ssh-manager/src/utils"
)

func New(filePaths models_file_paths.FilePaths, profileName string) error {
	fileManager := utils.FileManager{}

	sshName := profileName + "_github"

	profile, err := models_profile.CreateProfile(profileName, "", "", sshName, filePaths.ProfileFilePath)

	if err != nil {
		return err
	}

	existingProfiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)

	if err != nil {
		return err
	}

	fmt.Print("Username: ")
	fmt.Scanln(&profile.Username)

	fmt.Print("Email: ")
	fmt.Scanln(&profile.Email)

	profiles := make([]models_profile.Profile, 0)
	if len(existingProfiles) <= 0 {

		profiles = append(profiles, profile)

	} else {
		profiles = append(profiles, profile)
		profiles = append(profiles, existingProfiles...)

	}

	sshDirPath := filepath.Join(filePaths.SSHDirPath, sshName)
	sshFilePath := filepath.Join(sshDirPath, sshName)

	err = fileManager.CreateDir(sshDirPath)

	if err != nil {
		fmt.Println(err)
	}

	exec.Command("ssh-keygen", "-t", "ed25519", "-C", profile.Email, "-f", sshFilePath).Run()

	fileManager.WriteFile(filePaths.ProfileFilePath, profiles)

	message := "\nCreated profile " + profileName

	utils.PrintSuccessMessage(message)

	return nil

}
