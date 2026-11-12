package app

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
)

func status(filePaths models.AppData) error {
	profiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		return err
	}

	activeProfile, err := models.GetActiveProfile(filePaths.ActiveProfileFilePath, profiles)
	if err != nil {
		return err
	}

	gpgStatus, gpgExpiry := checkGPGKey(activeProfile.GPGKeyID)

	fmt.Printf("\n\033[32m%s\033[0m\n", activeProfile.ProfileName)
	fmt.Println("----------------------------")
	fmt.Printf("Username : %s\n", strings.TrimSpace(activeProfile.Username))
	fmt.Printf("Email    : %s\n", strings.TrimSpace(activeProfile.Email))
	if activeProfile.GPGKeyID != "" {
		fmt.Printf("GPG Key  : %s\n", activeProfile.GPGKeyID)
		fmt.Printf("GPG Status: %s\n", gpgStatus)
		if gpgExpiry != "" {
			fmt.Printf("GPG Expiry: %s\n", gpgExpiry)
		}
	} else {
		fmt.Println("GPG Key  : (none)")
	}
	fmt.Println()

	return nil
}

// checkGPGKey reports the status and expiry of the given GPG key.
func checkGPGKey(keyID string) (string, string) {
	keyID = strings.TrimSpace(keyID)
	if keyID == "" {
		return "not set", ""
	}

	out, err := exec.Command("gpg", "--batch", "--with-colons", "--list-keys", keyID).Output()
	if err != nil {
		return "not found", ""
	}

	for _, line := range strings.Split(string(out), "\n") {
		if !strings.HasPrefix(line, "pub:") {
			continue
		}
		fields := strings.Split(line, ":")
		if len(fields) < 7 {
			continue
		}
		expiryField := strings.TrimSpace(fields[6])
		if expiryField == "" || expiryField == "0" {
			return "valid (no expiry)", "never"
		}
		ts, err := strconv.ParseInt(expiryField, 10, 64)
		if err != nil {
			return "unknown", ""
		}
		expiryTime := time.Unix(ts, 0)
		if time.Now().After(expiryTime) {
			return "expired", expiryTime.Format("2006-01-02")
		}
		return "valid", expiryTime.Format("2006-01-02")
	}

	return "not found", ""
}