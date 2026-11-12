package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func getKey(filePaths models.AppData, profileName string) error {
	if err := utils.ValidateProfileName(profileName); err != nil {
		return err
	}
	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	prof, err := models.GetProfile(profileName, profiles)
	if err != nil {
		return err
	}

	// SSH public key
	sshPubPath := filepath.Join(filePaths.SSHDirPath, prof.SSHDirPath, prof.SSHDirPath+".pub")
	sshData, err := os.ReadFile(sshPubPath)
	if err != nil {
		return fmt.Errorf("failed to read SSH public key: %w", err)
	}

	fmt.Println("SSH Public Key:")
	fmt.Println(strings.TrimSpace(string(sshData)))
	fmt.Println()

	// GPG public key
	if prof.GPGKeyID != "" {
		out, err := exec.Command("gpg", "--batch", "--armor", "--export", prof.GPGKeyID).Output()
		if err != nil {
			return fmt.Errorf("failed to export GPG public key: %w", err)
		}
		if len(strings.TrimSpace(string(out))) == 0 {
			return fmt.Errorf("GPG public key not found for %s", prof.GPGKeyID)
		}
		fmt.Println("GPG Public Key:")
		fmt.Println(strings.TrimSpace(string(out)))
	}

	return nil
}