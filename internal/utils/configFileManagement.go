package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/types"
)

func GetConfigFields() (types.Config, error) {
	configFilePath := path.Join(paths.GetGitctxConfigFolderPath(), ".config.json")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return types.Config{}, fmt.Errorf("config file does not exist at path: %s", configFilePath)
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return types.Config{}, err
	}

	var config types.Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return types.Config{}, err
	}

	return config, nil
}

func GetCurrentContextName() (string, error) {
	config, err := GetConfigFields()
	if err != nil {
		return "", err
	}
	return config.CurrentContext, nil
}

func GetCurrentGlobalContextName() (string, error) {
	config, err := GetConfigFields()
	if err != nil {
		return "", err
	}
	return config.CurrentGlobalContext, nil
}

func UpdateConfigField(updateFn func(*types.Config)) error {
	configFilePath := path.Join(paths.GetGitctxConfigFolderPath(), ".config.json")

	// Check existence
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist at path: %s", configFilePath)
	}

	// Read file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal JSON
	var config types.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse config JSON: %w", err)
	}

	// Apply update
	updateFn(&config)

	// Marshal and write back
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	if err := os.WriteFile(configFilePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func UpdateCurrentContextName(newContextName string) error {
	return UpdateConfigField(func(c *types.Config) {
		c.CurrentContext = newContextName
	})
}

func UpdateCurrentGlobalContextName(newGlobalContextName string) error {
	return UpdateConfigField(func(c *types.Config) {
		c.CurrentGlobalContext = newGlobalContextName
	})
}

func ApplyContext(global bool) error {

	contextName, err := GetCurrentContextName()
	if err != nil {
		return err
	}

	contextFile, err := os.Open(filepath.Join(paths.GetGitctxConfigFolderPath(), contextName))
	if err != nil {
		return err
	}
	defer contextFile.Close()

	var configFile *os.File

	if global {
		configFile, err = os.OpenFile(filepath.Join(paths.GetHomeFolderPath(), ".gitconfig"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer configFile.Close()

		if err := UpdateCurrentGlobalContextName(contextName); err != nil {
			return err
		}
	} else {
		configFile, err = os.OpenFile(filepath.Join(paths.GetLocalGitFolderPath(), "config"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer configFile.Close()
	}

	_, err = io.Copy(configFile, contextFile)
	if err != nil {
		return err
	}

	return nil
}
