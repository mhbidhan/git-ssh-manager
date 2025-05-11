package git_ssh_manager

import (
	"os/exec"
	"path/filepath"

	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
	models_profile "github.com/mhbidhan/git-ssh-manager/src/models/profile"
	"github.com/mhbidhan/git-ssh-manager/src/utils"
)

func Use(filePaths models_file_paths.FilePaths, profileName string) error {
	fileManager := utils.FileManager{}

	profiles, err := models_profile.GetProfiles(filePaths.ProfileFilePath)

	err = fileManager.WriteFile(filePaths.ActiveProfileFilePath, profileName)

	profileToBeActive, err := models_profile.GetProfile(profileName, profiles)

	if err != nil {
		return err
	}

	osSSHDir := filepath.Join(filePaths.HomeDir, ".ssh")

	exec.Command("rm", filepath.Join(osSSHDir, "id_rsa")).Run()
	exec.Command("rm", filepath.Join(osSSHDir, "id_rsa.pub")).Run()

	exec.Command("git", "config", "--global", "user.name", profileToBeActive.Username).Run()
	exec.Command("git", "config", "--global", "user.email", profileToBeActive.Email).Run()

	privateKey := filepath.Join(filePaths.SSHDirPath, profileToBeActive.SSHDirPath, profileToBeActive.SSHDirPath)
	publicKey := filepath.Join(filePaths.SSHDirPath, profileToBeActive.SSHDirPath, profileToBeActive.SSHDirPath+".pub")

	exec.Command("cp", privateKey, osSSHDir).Run()
	exec.Command("cp", publicKey, osSSHDir).Run()

	privateKey = filepath.Join(osSSHDir, profileToBeActive.SSHDirPath)
	publicKey = filepath.Join(osSSHDir, profileToBeActive.SSHDirPath+".pub")

	exec.Command("mv", privateKey, filepath.Join(osSSHDir, "id_rsa")).Run()
	exec.Command("mv", publicKey, filepath.Join(osSSHDir, "id_rsa.pub")).Run()

	exec.Command("rm", privateKey).Run()
	exec.Command("rm", publicKey).Run()

	return nil
}
