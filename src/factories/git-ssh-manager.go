package factories

import (
	"errors"
	"os"

	app "github.com/mhbidhan/git-ssh-manager/src/app/platforms"
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
		return app.UnixInterfacePlatform{
			FilePaths: models_file_paths.GenerateFilePaths(homeDir, "/.config/git-ssh-manager"),
		}, nil
	default:
		return nil, errors.New("Platform not supported")
	}
}
