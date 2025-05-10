package platforms

import (
	"github.com/mhbidhan/github-ssh-manager/src/app/github_ssh_manager"
	models_file_paths "github.com/mhbidhan/github-ssh-manager/src/models/file_paths"
)

type UnixInterfacePlatform struct {
	models_file_paths.FilePaths
}

func (app UnixInterfacePlatform) GetFilePaths() models_file_paths.FilePaths {
	return app.FilePaths
}

func (app UnixInterfacePlatform) Status() error {
	return github_ssh_manager.Status(app.FilePaths)
}

func (app UnixInterfacePlatform) GetKey(profileName string) error {
	return github_ssh_manager.GetKey(app.FilePaths, profileName)
}

func (app UnixInterfacePlatform) New(profileName string) error {
	return github_ssh_manager.New(app.FilePaths, profileName)
}

func (app UnixInterfacePlatform) Use(profileName string) error {
	return github_ssh_manager.Use(app.FilePaths, profileName)
}

func (app UnixInterfacePlatform) Remove(profileName string) error {
	return github_ssh_manager.Remove(app.FilePaths, profileName)
}
