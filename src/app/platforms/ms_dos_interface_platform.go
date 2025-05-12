package platforms

import (
	"github.com/mhbidhan/git-ssh-manager/src/app/git_ssh_manager"
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
)

type MSDOSInterfacePlatform struct {
	models_file_paths.FilePaths
}

func (app MSDOSInterfacePlatform) GetFilePaths() models_file_paths.FilePaths {
	return app.FilePaths
}

func (app MSDOSInterfacePlatform) Status() error {
	return git_ssh_manager.Status(app.FilePaths)
}

func (app MSDOSInterfacePlatform) GetKey(profileName string) error {
	return git_ssh_manager.GetKey(app.FilePaths, profileName)
}

func (app MSDOSInterfacePlatform) New(profileName string) error {
	return git_ssh_manager.New(app.FilePaths, profileName)
}

func (app MSDOSInterfacePlatform) Use(profileName string) error {
	return git_ssh_manager.Use(app.FilePaths, profileName)
}

func (app MSDOSInterfacePlatform) Remove(profileName string) error {
	return git_ssh_manager.Remove(app.FilePaths, profileName)
}
