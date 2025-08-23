package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/types"
)

func GetCurrentContextName() (string, error) {
	configFilePath := path.Join(paths.GetGitctxConfigFolderPath(), ".config.json")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("config file does not exist at path: %s", configFilePath)
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}

	var config types.Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return "", err
	}

	return config.CurrentContext, nil
}

func UpdateCurrentContextName(newContextName string) error {
	configFilePath := path.Join(paths.GetGitctxConfigFolderPath(), ".config.json")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist at path: %s", configFilePath)
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	var config types.Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return err
	}

	config.CurrentContext = newContextName

	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
