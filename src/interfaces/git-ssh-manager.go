package interfaces

import models_file_paths "github.com/mhbidhan/git-ssh-manager/src/models/file_paths"

type GithubSshManager interface {
	GetFilePaths() models_file_paths.FilePaths

	List() error
	Status() error
	GetKey(profileName string) error
	New(profileName string) error
	Use(profileName string) error
	Remove(profileName string) error
}
