// Package app provides the main SSH manager application logic.
package app

import "github.com/mhbidhan/git-ssh-manager/internal/models"

// GitHubSSHManager defines the interface for SSH manager operations.
type GitHubSSHManager interface {
	FilePaths() models.AppData
	List() error
	Status() error
	GetKey(profileName string) error
	New(profileName string) error
	Use(profileName string) error
	Remove(profileName string) error
	BindProfile(profileName string) error
}
