package paths

import (
	"os"
	"path/filepath"
)

var configFolderPathOverride string
var homeFolderPathOverride string

func OverrideConfigFolderPath(path string) {
	configFolderPathOverride = path
}

func OverrideHomeFolderPath(path string) {
	homeFolderPathOverride = path
}

func GetHomeFolderPath() string {
	var homeDirPath string

	if homeFolderPathOverride != "" {
		homeDirPath = homeFolderPathOverride
	} else {
		var err error
		homeDirPath, err = os.UserHomeDir()
		if err != nil {
			homeDirPath = "."
		}
	}
	return homeDirPath
}

func GetConfigFolderPath() string {
	var configDirPath string

	if configFolderPathOverride != "" {
		configDirPath = configFolderPathOverride
	} else {
		var err error
		configDirPath, err = os.UserConfigDir()
		if err != nil {
			configDirPath = "."
		}
	}
	return configDirPath
}

func GetGitctxConfigFolderPath() string {
	return filepath.Join(GetConfigFolderPath(), "gitctx")
}

func FindGitConfigsFilePaths() ([]string, error) {

	var gitConfigsPaths []string = make([]string, 0)

	// Check both common user git config file locations
	gitHomeFilePath := filepath.Join(GetHomeFolderPath(), ".gitconfig")
	gitConfigFilePath := filepath.Join(GetConfigFolderPath(), "git", "config")

	// Check if the files exist and add them to the list
	if _, err := os.Stat(gitHomeFilePath); err == nil {
		gitConfigsPaths = append(gitConfigsPaths, gitHomeFilePath)
	}
	if _, err := os.Stat(gitConfigFilePath); err == nil {
		gitConfigsPaths = append(gitConfigsPaths, gitConfigFilePath)
	}

	return gitConfigsPaths, nil
}

func ListAllGitctxContexts() []string {

	var contexts []string = make([]string, 0)

	entries, err := os.ReadDir(GetGitctxConfigFolderPath())
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		if !entry.IsDir() && entry.Name()[0] != '.' {
			contexts = append(contexts, entry.Name())
		}
	}
	return contexts
}
