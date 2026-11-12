package app

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func bindProfile(filePaths models.AppData, profileName string) error {
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

	gitRoot, err := getGitRoot()
	if err != nil {
		return err
	}

	privateKey := filepath.Join(filePaths.SSHDirPath, prof.SSHDirPath, prof.SSHDirPath)
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("git not found: %w", err)
	}

	sshCommand := "ssh -i " + shellQuote(privateKey) + " -o IdentitiesOnly=yes"

	if err := exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "core.sshCommand", sshCommand).Run(); err != nil {
		return fmt.Errorf("failed to set core.sshCommand: %w", err)
	}

	if name := strings.TrimSpace(prof.Username); name != "" {
		_ = exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "user.name", name).Run()
	}
	if email := strings.TrimSpace(prof.Email); email != "" {
		_ = exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "user.email", email).Run()
	}
	gpgKeyID := prof.GPGKeyID
	if gpgKeyID != "" {
		_ = exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "user.signingkey", strings.TrimSpace(gpgKeyID)).Run()
		status, _ := checkGPGKey(gpgKeyID)
		if strings.HasPrefix(status, "valid") {
			_ = exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "commit.gpgsign", "true").Run()
		} else {
			_ = exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "commit.gpgsign", "false").Run()
		}
	} else {
		_ = exec.CommandContext(context.Background(), "git", "-C", gitRoot, "config", "--unset", "commit.gpgsign").Run()
	}

	utils.PrintSuccessMessage("Bound repo " + gitRoot + " to profile " + profileName)
	return nil
}

// shellQuote wraps s in single quotes, escaping embedded single quotes.
func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

func getGitRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", errors.New("not a git repository (run this inside a repo you want to bind)")
	}

	root := strings.TrimSpace(string(out))
	if root == "" {
		return "", errors.New("failed to detect git root directory")
	}
	return root, nil
}
