package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	gitctxTypes "github.com/alexandreLITHAUD/gitctx/internal/types"
	"github.com/alexandreLITHAUD/gitctx/internal/ui"
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

		var config gitctxTypes.Config = gitctxTypes.Config{
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

func CreateGitctxContextFileFromPath(name string, path string) error {

	contextFilePath := filepath.Join(paths.GetGitctxConfigFolderPath(), name)

	if DoesContextFileExists(name) {
		fmt.Printf("Context '%s' already exists. Would you like to override it ?.\n", name)
		confirmed, err := ui.PromptForConfirmation("Answer (y)es or (n)o :")
		if err != nil {
			return fmt.Errorf("error reading confirmation input: %w", err)
		}
		if !confirmed {
			fmt.Println("Skipping creation of context file.")
			return nil
		}
	}

	srcFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening source file '%s': %w", path, err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(contextFilePath)
	if err != nil {
		return fmt.Errorf("error creating destination file '%s': %w", contextFilePath, err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("error copying content from '%s' to '%s': %w", path, contextFilePath, err)
	}
	fmt.Printf("Context '%s' created successfully in path %s from '%s'.\n", name, contextFilePath, path)

	return nil
}
