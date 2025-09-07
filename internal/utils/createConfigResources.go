package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	gitctxTypes "github.com/alexandreLITHAUD/gitctx/internal/types"
)

func DoesConfigFolderExists() bool {
	configDir := paths.GetGitctxConfigFolderPath()
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateGitctxConfigFolder() error {
	configDir := paths.GetGitctxConfigFolderPath()

	if configDir == "" {
		return os.ErrNotExist
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(filepath.Join(configDir, ".config.json")); os.IsNotExist(err) {
		configFile, err := os.Create(filepath.Join(configDir, ".config.json"))
		if err != nil {
			return err
		}
		defer configFile.Close()

		var config = gitctxTypes.Config{
			CurrentContext: "",
		}

		jsonData, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return err
		}

		_, err = configFile.WriteString(string(jsonData))
		if err != nil {
			return err
		}
	}

	return nil
}
