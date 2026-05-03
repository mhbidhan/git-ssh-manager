// Package app provides the main SSH manager application logic.
package app

import "github.com/mhbidhan/git-ssh-manager/internal/models"

// GitSSHManager manages git SSH profiles.
type GitSSHManager struct {
	appData models.AppData
}

// NewGitSSHManager creates a new GitSSHManager with the given app data.
func NewGitSSHManager(ad models.AppData) GitSSHManager {
	return GitSSHManager{appData: ad}
}

// FilePaths returns the file paths used by the manager.
func (app GitSSHManager) FilePaths() models.AppData {
	return app.appData
}

// Status displays the currently active profile.
func (app GitSSHManager) Status() error {
	return status(app.FilePaths())
}

// List outputs all available profiles.
func (app GitSSHManager) List() error {
	return list(app.FilePaths())
}

// GetKey prints the public key for the given profile.
func (app GitSSHManager) GetKey(profileName string) error {
	return getKey(app.FilePaths(), profileName)
}

// New creates a new SSH profile.
func (app GitSSHManager) New(profileName string) error {
	return newProfile(app.FilePaths(), profileName)
}

// Use switches to the specified profile.
func (app GitSSHManager) Use(profileName string) error {
	return useProfile(app.FilePaths(), profileName)
}

// Remove deletes the specified profile.
func (app GitSSHManager) Remove(profileName string) error {
	return removeProfile(app.FilePaths(), profileName)
}

// BindProfile binds a profile to the current git repository.
func (app GitSSHManager) BindProfile(profileName string) error {
	return bindProfile(app.FilePaths(), profileName)
}
