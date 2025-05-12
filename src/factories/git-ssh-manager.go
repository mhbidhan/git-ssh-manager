package factories

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mhbidhan/git-ssh-manager/src/app/app_git_ssh_manager"
	"github.com/mhbidhan/git-ssh-manager/src/interfaces"
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
)

func CreateGithubSshManager(platform string) (interfaces.GithubSshManager, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	switch platform {
	case "linux":
		return app_git_ssh_manager.GitSSHManager{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, filepath.Join(".config", "git-ssh-manager")),
		}, nil
	case "darwin":
		return app_git_ssh_manager.GitSSHManager{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, filepath.Join(".config", "git-ssh-manager")),
		}, nil
	case "windows":
		return app_git_ssh_manager.GitSSHManager{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, filepath.Join("AppData", "Local", "git-ssh-manager")),
		}, nil
	default:
		return nil, errors.New("Platform not supported")
	}
}
