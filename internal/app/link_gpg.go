package app

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func linkGPGProfile(filePaths models.AppData, profileName string, gpgKeyID string) error {
	if err := utils.ValidateProfileName(profileName); err != nil {
		return err
	}
	fileManager := utils.NewFileManager()

	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return fmt.Errorf("failed to get profiles: %w", err)
	}

	prof, err := models.GetProfile(profileName, profiles)
	if err != nil {
		return err
	}

	if strings.TrimSpace(gpgKeyID) == "" {
		fmt.Print("GPG Key ID: ")
		inStream := bufio.NewReader(os.Stdin)
		input, err := inStream.ReadString('\n')
		if err != nil && strings.TrimSpace(input) == "" {
			return fmt.Errorf("failed to read GPG key id: %w", err)
		}
		gpgKeyID = strings.TrimSpace(input)
	}
	if gpgKeyID == "" {
		return fmt.Errorf("GPG key id cannot be empty")
	}

	prof.GPGKeyID = gpgKeyID

	// Update the profile in the slice.
	for i := range profiles {
		if profiles[i].ProfileName == profileName {
			profiles[i].GPGKeyID = gpgKeyID
			break
		}
	}

	if err := fileManager.WriteFile(filePaths.ProfileFilePath, profiles); err != nil {
		return err
	}

	// If this is the active profile, update git config too.
	file, err := fileManager.ReadFile(filePaths.ActiveProfileFilePath)
	if err == nil {
		var activeUser string
		if json.Unmarshal([]byte(file), &activeUser) == nil && activeUser == profileName {
			_ = exec.CommandContext(context.Background(), "git", "config", "--global", "user.signingkey", gpgKeyID).Run()
			status, _ := checkGPGKey(gpgKeyID)
			if strings.HasPrefix(status, "valid") {
				_ = exec.CommandContext(context.Background(), "git", "config", "--global", "commit.gpgsign", "true").Run()
			} else {
				_ = exec.CommandContext(context.Background(), "git", "config", "--global", "commit.gpgsign", "false").Run()
			}
		}
	}

	utils.PrintSuccessMessage("Linked GPG key " + gpgKeyID + " to profile " + profileName)
	return nil
}