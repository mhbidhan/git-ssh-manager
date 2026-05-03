package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func newProfile(filePaths models.AppData, profileName string) error {
	fileManager := utils.NewFileManager()

	sshName := profileName + "_github"

	existingProfiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return fmt.Errorf("failed to get existing profiles: %w", err)
	}

	inStream := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	profileNameInput, err := inStream.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read username: %w", err)
	}

	fmt.Print("Email: ")
	profileEmailInput, err := inStream.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}

	if err := utils.ValidateEmail(profileEmailInput); err != nil {
		return err
	}

	prof, err := models.CreateProfile(profileName, profileNameInput, profileEmailInput, sshName, filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	var profiles []models.Profile
	profiles = append(profiles, existingProfiles...)
	profiles = append(profiles, prof)

	sshDirPath := filepath.Join(filePaths.SSHDirPath, sshName)
	sshFilePath := filepath.Join(sshDirPath, sshName)

	if err := fileManager.CreateDir(sshDirPath); err != nil {
		return fmt.Errorf("failed to create SSH dir: %w", err)
	}

	if err := exec.CommandContext(context.Background(), "ssh-keygen", "-t", "ed25519", "-C", prof.Email, "-f", sshFilePath).Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	if err := fileManager.WriteFile(filePaths.ProfileFilePath, profiles); err != nil {
		return err
	}

	utils.PrintSuccessMessage("\nCreated profile " + profileName)

	return nil
}
