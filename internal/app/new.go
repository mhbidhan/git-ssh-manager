package app

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mhbidhan/git-ssh-manager/internal/models"
	"github.com/mhbidhan/git-ssh-manager/internal/utils"
)

func newProfile(filePaths models.AppData, profileName string) error {
	if err := utils.ValidateProfileName(profileName); err != nil {
		return err
	}
	fileManager := utils.NewFileManager()

	sshName := profileName + "_github"

	existingProfiles, err := models.GetProfiles(filePaths.ProfileFilePath)
	if err != nil {
		// Profile file may be empty/uninitialized; ignore read errors here.
		existingProfiles = nil
	}

	inStream := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	profileUsernameInput, err := inStream.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read username: %w", err)
	}
	profileUsernameInput = strings.TrimSpace(profileUsernameInput)

	fmt.Print("Email: ")
	profileEmailInput, err := inStream.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}
	profileEmailInput = strings.TrimSpace(profileEmailInput)

	if err := utils.ValidateEmail(profileEmailInput); err != nil {
		return err
	}

	// Fail early if profile already exists.
	for _, p := range existingProfiles {
		if p.ProfileName == profileName {
			return errors.New("profile already exists")
		}
	}

	// Allow linking an existing GPG key instead of generating a new one.
	fmt.Print("GPG Key ID (leave empty to generate new): ")
	gpgKeyInput, _ := inStream.ReadString('\n')
	gpgKeyInput = strings.TrimSpace(gpgKeyInput)

	var gpgKeyID string
	if gpgKeyInput != "" {
		gpgKeyID = gpgKeyInput
	} else {
		fmt.Print("Expiry (optional, e.g. 0 for no expiry, 1y, 6m, 30d; default 0): ")
		expiryInput, _ := inStream.ReadString('\n')
		expiryInput = strings.TrimSpace(expiryInput)
		if expiryInput == "" {
			expiryInput = "0"
		}

		generated, err := generateGPGKey(profileUsernameInput, profileEmailInput, expiryInput)
		if err != nil {
			return fmt.Errorf("failed to generate GPG key: %w", err)
		}
		gpgKeyID = generated
	}

	prof, err := models.CreateProfile(profileName, profileUsernameInput, profileEmailInput, sshName, gpgKeyID, filePaths.ProfileFilePath)
	if err != nil {
		return err
	}
	// Ensure GPG key id and path are populated.
	prof.GPGKeyID = gpgKeyID

	var profiles []models.Profile
	profiles = append(profiles, existingProfiles...)
	profiles = append(profiles, prof)

	sshDirPath := filepath.Join(filePaths.SSHDirPath, sshName)
	sshFilePath := filepath.Join(sshDirPath, sshName)

	if err := fileManager.CreateDir(sshDirPath); err != nil {
		return fmt.Errorf("failed to create SSH dir: %w", err)
	}

	if err := exec.CommandContext(context.Background(), "ssh-keygen", "-t", "ed25519", "-C", prof.Email, "-f", sshFilePath).Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	if err := fileManager.WriteFile(filePaths.ProfileFilePath, profiles); err != nil {
		return err
	}

	utils.PrintSuccessMessage("\nCreated profile " + profileName)

	return nil
}

// generateGPGKey generates a new GPG signing key and returns its fingerprint.
func generateGPGKey(name, email, expiry string) (string, error) {
	paramsFile, err := os.CreateTemp("", "gpg-params-*")
	if err != nil {
		return "", err
	}
	paramsPath := paramsFile.Name()
	defer os.Remove(paramsPath)

	content := fmt.Sprintf(`%%no-protection
Key-Type: RSA
Key-Length: 2048
Key-Usage: sign
Name-Real: %s
Name-Email: %s
Expire-Date: %s
%%commit
`, name, email, expiry)
	if _, err := paramsFile.WriteString(content); err != nil {
		paramsFile.Close()
		return "", err
	}
	paramsFile.Close()

	cmd := exec.Command("gpg", "--batch", "--status-fd", "1", "--generate-key", paramsPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(out))
	}

	re := regexp.MustCompile(`KEY_CREATED \S+ (\S+)`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) == 2 {
		return strings.TrimSpace(matches[1]), nil
	}

	// Fallback: look up by email.
	return lookupGPGFingerprint(email)
}

func lookupGPGFingerprint(email string) (string, error) {
	out, err := exec.Command("gpg", "--batch", "--with-colons", "--list-secret-keys", email).Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(out), "\n")
	var currentFpr string
	for _, line := range lines {
		if strings.HasPrefix(line, "fpr:") {
			fields := strings.Split(line, ":")
			if len(fields) > 9 {
				currentFpr = fields[9]
			}
		}
		if strings.HasPrefix(line, "uid:") && strings.Contains(line, email) {
			if currentFpr != "" {
				return currentFpr, nil
			}
		}
	}
	// Last resort: return last fingerprint seen.
	if currentFpr != "" {
		return currentFpr, nil
	}
	return "", fmt.Errorf("could not determine GPG key fingerprint")
}

