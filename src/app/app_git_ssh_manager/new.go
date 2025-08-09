package app_git_ssh_manager

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
	"github.com/mhbidhan/git-ssh-manager/src/utils"
)

func new(filePaths models_file_paths.FilePaths, profileName string) error {
	fileManager := utils.FileManager{}

	sshName := profileName + "_github"

	existingProfiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)

	if err != nil {
		return err
	}

	inStream := bufio.NewReader(os.Stdin)

	var profileEmialInput string

	fmt.Print("Username: ")
	profileNameInput, errNameInput := inStream.ReadString('\n')
	if errNameInput != nil {
		return errNameInput
	}

	fmt.Print("Email: ")
	profileEmialInput, errEmailInput := inStream.ReadString('\n')
	if errEmailInput != nil {
		return errEmailInput
	}

	profile, err := models_profile.CreateProfile(profileName, profileNameInput, profileEmialInput, sshName, filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

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
