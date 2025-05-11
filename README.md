# Git SSH Manager

**Git SSH Manager** is a simple command-line utility that helps you manage multiple Git SSH profiles with ease. It allows you to create, switch, and remove SSH keys and Git global user configurations (username and email) using straightforward commands.

## Features

- Quickly generate new SSH keys and Git profiles
- Seamlessly switch between profiles for different projects
- View the current profile and its public key
- Remove profiles you no longer need

## Supported Platforms

- ✅ Linux
- ✅ macOS

## Installation

### Linux/MacOS

> **Note:** `git-ssh-manager` requires [Go](https://golang.org/dl/) version **1.24.2** or higher if building from source.

---

### 🔧 Option 1: Build from Source

1. Clone the repository:

   ```bash
   $ git clone https://github.com/mhbidhan/git-ssh-manager.git
   $ cd git-ssh-manager
   ```

2. Build and install:
   ```bash
   make install
   ```

---

### 📦 Option 2: Install Precompiled Binary

1. Download the latest precompiled binary:

   👉 [Download from Google Drive](https://drive.google.com/file/d/1nLueHClz0-1YVyYJvc0-Y3trZWBHKsfM/view?usp=drive_link)

2. Unzip the archive:

   ```text
   /files
   ├── gsm
   └── install.sh
   ```

3. Run the install script:

   ```bash
   $ chmod +x install.sh
   $ ./install.sh
   ```

---

### ✅ Verify Installation

After installation, check that it works:

```bash
$ gsm -h
```

Expected output:

```text
-gk, --get-key          prints the public key
-s, --status            show active profile
-n, --new               create new profile
-u, --use               use to profile
-r, --remove            remove profile
-h, --help              list all command
```

## Usage

To use `git-ssh-manager`, simply run the following command:

```bash
# gsm --(COMMAND)
$ gsm -s
# or
$ gsm --status
```

Below are the supported commands:

| Command            | Description                                          |
| ------------------ | ---------------------------------------------------- |
| `-gk`, `--get-key` | Prints the public key of the current profile.        |
| `-s`, `--status`   | Displays the currently active profile.               |
| `-n`, `--new`      | Creates a new profile.                               |
| `-u`, `--use`      | Switches to the specified profile.                   |
| `-r`, `--remove`   | Removes an existing profile.                         |
| `-h`, `--help`     | Lists all available commands and their descriptions. |

## Examples

```bash
# Create a new profile
$ gsm -n profile-name # Promts for username and email

# Switch to a profile
$ gsm -u work

# View current active profile
$ gsm -s

# Get the public key of the profile
$ gsm -gk profile-name

# Remove a profile
$ gsm -r work # Works for non-active profiles
```

## Command Breakdown

### `gsm -n`, `--new`

Prompts for username and email to create a new SSH profile.  
The username and email should match your GitHub (or other remote Git service) identity.

```bash
$ gsm -n work
```

<pre>
Username: work_user
Email: user@work.com
</pre>

---

### `gsm -u`, `--use`

Activates the specified SSH profile and updates the global Git username and email.

```bash
$ gsm -u work
```

```bash
$ git config --list

user.name=work_user
user.email=user@work.com
```

---

### `gsm -gk`, `--get-key`

Prints the public key of the specified profile.  
You can copy and add this to your GitHub or other Git host.

```bash
gsm -gk work
```

<pre>
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFK... user@work.com
</pre>
