package app_git_ssh_manager

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
	"github.com/mhbidhan/git-ssh-manager/src/utils"
)

func bindProfile(filePaths models_file_paths.FilePaths, profileName string) error {
	profiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	profile, err := models_profile.GetProfile(profileName, profiles)
	if err != nil {
		return err
	}

	gitRoot, err := getGitRoot()
	if err != nil {
		return err
	}

	privateKey := filepath.Join(filePaths.SSHDirPath, profile.SSHDirPath, profile.SSHDirPath)
	if _, err := os.Stat(privateKey); err != nil {
		return err
	}

	sshCommand := "ssh -i " + privateKey + " -o IdentitiesOnly=yes"

	// Bind the SSH identity for this repo only
	if err := exec.Command("git", "-C", gitRoot, "config", "core.sshCommand", sshCommand).Run(); err != nil {
		return err
	}

	// Optional (but usually desired): bind git author identity for this repo only
	if name := strings.TrimSpace(profile.Username); name != "" {
		_ = exec.Command("git", "-C", gitRoot, "config", "user.name", name).Run()
	}
	if email := strings.TrimSpace(profile.Email); email != "" {
		_ = exec.Command("git", "-C", gitRoot, "config", "user.email", email).Run()
	}

	utils.PrintSuccessMessage("Bound repo " + gitRoot + " to profile " + profileName)
	return nil
}

func getGitRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", errors.New("Not a git repository (run this inside a repo you want to bind)")
	}

	root := strings.TrimSpace(string(out))
	if root == "" {
		return "", errors.New("Failed to detect git root directory")
	}
	return root, nil
}

