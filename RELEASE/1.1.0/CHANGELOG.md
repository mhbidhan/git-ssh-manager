# Release Notes — Git SSH Manager

## v1.1.0

### Highlights

- Added GPG key support to profiles: each profile can store a unique GPG signing
  key ID/fingerprint.
- New `--link-gpg` / `-lg` command to link an existing GPG key to a profile.
- Profile creation now optionally generates a GPG key with configurable expiry,
  or links an existing one.
- `use` and `bind-profile` now update `user.signingkey` and `commit.gpgsign`
  based on key validity.
- `get-key` now prints both SSH and GPG public keys with clear labels.
- `status` now shows GPG key validity and expiry.

### New Features

- **Per-profile GPG keys**: `Profile` stores `GPGKeyID`; generated during
  `gsm -n`.
- **Link existing GPG key**: `gsm --link-gpg <profile> [key-id]` or prompted
  mode.
- **Expiry prompt**: during creation, optional expiry (`0`, `1y`, `6m`, `30d`,
  etc.).
- **GPG public key export on demand** via `gsm -gk <profile>`.

### Improvements

- Better `status` formatting and GPG validity check.
- `commit.gpgsign` enabled automatically when a valid GPG key is active.
- Updated Makefile for current `./cmd` layout; added `test-run` target using
  `GSM_HOME`.
- `.gitignore` updated for local test config/bin artifacts.

### Security Hardening

- Profile name validation to prevent path traversal.
- Shell-quoted SSH key paths in `core.sshCommand`.
- Profile metadata files written with `0600` permissions.
- Safer SSH key copy with `0600` initial permissions.

### Bug Fixes / Maintenance

- Hardened error handling for empty/uninitialized profile files.

