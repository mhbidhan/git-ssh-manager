package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func useProfile(filePaths models.AppData, profileName string) error {
	fileManager := utils.NewFileManager()

	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return fmt.Errorf("failed to get profiles: %w", err)
	}

	if err := fileManager.WriteFile(filePaths.ActiveProfileFilePath, profileName); err != nil {
		return fmt.Errorf("failed to write active profile: %w", err)
	}

	prof, err := models.GetProfile(profileName, profiles)
	if err != nil {
		return err
	}

	osSSHDir := filepath.Join(filePaths.HomeDir, ".ssh")

	_ = os.Remove(filepath.Join(osSSHDir, "id_rsa"))
	_ = os.Remove(filepath.Join(osSSHDir, "id_rsa.pub"))

	if err := exec.CommandContext(context.Background(), "git", "config", "--global", "user.name", prof.Username).Run(); err != nil {
		return fmt.Errorf("failed to set git user.name: %w", err)
	}

	if err := exec.CommandContext(context.Background(), "git", "config", "--global", "user.email", prof.Email).Run(); err != nil {
		return fmt.Errorf("failed to set git user.email: %w", err)
	}

	privateKey := filepath.Join(filePaths.SSHDirPath, prof.SSHDirPath, prof.SSHDirPath)
	publicKey := privateKey + ".pub"

	if err := fileManager.CopyFileToDirectory(publicKey, osSSHDir); err != nil {
		return err
	}
	if err := fileManager.CopyFileToDirectory(privateKey, osSSHDir); err != nil {
		return err
	}

	osPrivateKey := filepath.Join(osSSHDir, prof.SSHDirPath)
	osPublicKey := osPrivateKey + ".pub"

	if err := fileManager.RenameFile(osPublicKey, "id_rsa.pub"); err != nil {
		return err
	}

	if err := fileManager.RenameFile(osPrivateKey, "id_rsa"); err != nil {
		return err
	}

	return nil
}
