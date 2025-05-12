package factories

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mhbidhan/git-ssh-manager/src/app/platforms"
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
		return platforms.UnixInterfacePlatform{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, filepath.Join(".config", "git-ssh-manager")),
		}, nil
	case "darwin":
		return platforms.UnixInterfacePlatform{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, filepath.Join(".config", "git-ssh-manager")),
		}, nil
	case "windows":
		return platforms.MSDOSInterfacePlatform{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, filepath.Join("AppData", "Local")),
		}, nil
	default:
		return nil, errors.New("Platform not supported")
	}
}
