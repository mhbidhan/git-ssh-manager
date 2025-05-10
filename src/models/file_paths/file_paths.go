package models_file_paths

type FilePaths struct {
	HomeDir               string
	RootDir               string
	SSHDirPath            string
	ProfileFilePath       string
	ActiveProfileFilePath string
}

func GenerateFilePaths(homeDir string, rootDir string) FilePaths {
	return FilePaths{

		HomeDir:               homeDir,
		RootDir:               homeDir + rootDir,
		ProfileFilePath:       homeDir + rootDir + "/profiles.json",
		SSHDirPath:            homeDir + rootDir + "/ssh_keys",
		ActiveProfileFilePath: homeDir + rootDir + "/active-profile.txt",
	}
}
