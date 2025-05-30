# Git SSH Manager

**Git SSH Manager** is a simple command-line utility that helps you manage multiple Git SSH profiles with ease. It allows you to create, switch, and remove SSH keys and Git global user configurations (username and email) using straightforward commands.

<blockquote>

The Story Behind `git-ssh-manager` [**My Journey Building My First CLI Tool**](https://medium.com/@mhbidhan/my-journey-building-my-first-cli-tool-0b37b2b8ef8c)

</blockquote>

## Features

- Quickly generate new SSH keys and Git profiles
- Seamlessly switch between profiles for different projects
- View the current profile and its public key
- Remove profiles you no longer need

## Supported Platforms

- ✅ Linux
- ✅ macOS
- ✅ Windows

## Installation

<details>
  <summary><strong>Linux/MacOS</strong></summary>

  <h3>🔧 Option 1: Build from Source</h3>

  <blockquote>
    <strong>Note:</strong> <code>git-ssh-manager</code> requires <a href="https://golang.org/dl/">Go</a> version <strong>1.24.2</strong> or higher if building from source.
  </blockquote>

  <ol>
    <li>
      Clone the repository:
      <pre><code class="language-bash">
$ git clone https://github.com/mhbidhan/git-ssh-manager.git
$ cd git-ssh-manager
      </code></pre>
    </li>
    <li>
      Build and install:
      <pre><code class="language-bash">
make install
      </code></pre>
    </li>
  </ol>

  <hr>

  <h3>📦 Option 2: Install Precompiled Binary</h3>

  <ol>
    <li>
      Download the latest precompiled binary:
      👉 <a href="https://drive.google.com/drive/folders/16Y5vDB5KJk1wbbk4NLPM6jUgYaYO2woH?usp=sharing">Download from Google Drive</a>
    </li>
    <li>
      Unzip the archive (this shows the file structure):
      <pre><code class="language-text">
files
├── gsm
└── install.sh
      </code></pre>
    </li>
    <li>
      Run the install script:
      <pre><code class="language-bash">
$ chmod +x install.sh
$ ./install.sh
      </code></pre>
    </li>
  </ol>

  <hr>

</details>

<details>
  <summary><strong>Windows</strong></summary>

  <h3>🔧 Option 1: Build from Source</h3>

  <blockquote>
    <strong>Note:</strong> <code>git-ssh-manager</code> requires <a href="https://golang.org/dl/">Go</a> version <strong>1.24.2</strong> or higher if building from source.
  </blockquote>

  <ol>
    <li>
      Clone the repository:
      <pre><code class="language-bash">
$ git clone https://github.com/mhbidhan/git-ssh-manager.git
$ cd git-ssh-manager
      </code></pre>
    </li>
    <li>
      Build and install:
      <pre><code class="language-text">
git-ssh-manager
└── src
    └── build_utils
        └── install.bat
      </code></pre>
    </li>
    <li>
      Run the install.bat as an Administrator
    </li>

  </ol>

  <hr>

  <h3>📦 Option 2: Install Precompiled Binary</h3>

  <ol>
    <li>
      Download the latest precompiled binary:
      👉 <a href="https://drive.google.com/drive/folders/1eYxDfsi55_bHe_qOLiARfiMi8_n1-dvW?usp=sharing">Download from Google Drive</a>
    </li>
    <li>
      Unzip the archive (this shows the file structure):
      <pre><code class="language-text">
/files
├── gsm
└── install.bat
└── update_path.bat
      </code></pre>
    </li>
    <li>
      Run the install.bat as an Administrator
    </li>
  </ol>

  <hr>

</details>

### ✅ Verify Installation

<blockquote>
  <strong>Note:</strong> The program might not work on Windows correctly. If it doesn't produce expected results, manually add <code>C:\Program Files\GitSSHManager</code> to your Windows environment paths. Windows sometimes fails to update this automatically.
</blockquote>

After installation, check that it works:

```bash
$ gsm -h
```

Expected output:

```text
-gk, --get-key          prints the public key
-l, --list              outputs a list of all available profiles
-s, --status            outputs the active profile
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
| `-l`, `--list`     | Displays a list of all available profiles.           |
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
