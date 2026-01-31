package app_git_ssh_manager

import (
	models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"
)

type GitSSHManager struct {
	models_file_paths.FilePaths
}

func (app GitSSHManager) GetFilePaths() models_file_paths.FilePaths {
	return app.FilePaths
}

func (app GitSSHManager) Status() error {
	return status(app.FilePaths)
}

func (app GitSSHManager) List() error {
	return list(app.FilePaths)
}

func (app GitSSHManager) GetKey(profileName string) error {
	return getKey(app.FilePaths, profileName)
}

func (app GitSSHManager) New(profileName string) error {
	return new(app.FilePaths, profileName)
}

func (app GitSSHManager) Use(profileName string) error {
	return use(app.FilePaths, profileName)
}

func (app GitSSHManager) Remove(profileName string) error {
	return remove(app.FilePaths, profileName)
}

func (app GitSSHManager) BindProfile(profileName string) error {
	return bindProfile(app.FilePaths, profileName)
}
