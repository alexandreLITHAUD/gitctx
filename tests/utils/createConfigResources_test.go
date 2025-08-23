package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
)

func TestCreateGitctxConfigFolder(t *testing.T) {
	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	// Ensure the gitctx config folder exists
	if err := utils.CreateGitctxConfigFolder(); err != nil {
		t.Fatalf("Failed to create gitctx config folder: %v", err)
	}

	// Check if the folder was created
	configDir := paths.GetGitctxConfigFolderPath()
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		t.Fatalf("Gitctx config folder does not exist at path: %s", configDir)
	}

	// Check if the .config.json file was created
	configFilePath := filepath.Join(configDir, ".config.json")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		t.Fatalf(".config.json file does not exist at path: %s", configFilePath)
	}

	// Test creating a context file
	contextName := "testContext"
	contextPath := filepath.Join(tempDir, "testContextFile")
	os.WriteFile(contextPath, []byte("test context data"), 0644)

	if err := utils.CreateGitctxContextFileFromPath(contextName, contextPath); err != nil {
		t.Fatalf("Failed to create context file: %v", err)
	}

	// Verify the context file was created
	if _, err := os.Stat(filepath.Join(configDir, contextName)); os.IsNotExist(err) {
		t.Fatalf("Context file '%s' was not created in the config folder", contextName)
	}

	fmt.Println("TestCreateGitctxConfigFolder completed successfully!")
}

func TestDoesConfigFolderExists(t *testing.T) {
	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	// Ensure the gitctx config folder exists
	if err := utils.CreateGitctxConfigFolder(); err != nil {
		t.Fatalf("Failed to create gitctx config folder: %v", err)
	}

	// Check if the config folder exists
	if !utils.DoesConfigFolderExists() {
		t.Fatalf("Config folder should exist at path: %s", paths.GetGitctxConfigFolderPath())
	}

	fmt.Println("TestDoesConfigFolderExists completed successfully!")
}

func TestCreateGitctxContextFileFromPath(t *testing.T) {

	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	// Create a temporary gitctx config folder
	if err := utils.CreateGitctxConfigFolder(); err != nil {
		t.Fatalf("Failed to create gitctx config folder: %v", err)
	}

	// Create a mock context file
	contextName := "testContext"
	contextPath := filepath.Join(tempDir, "testContextFile")
	if err := os.WriteFile(contextPath, []byte("test context data"), 0644); err != nil {
		t.Fatalf("Failed to create mock context file: %v", err)
	}

	// Test creating the context file from the path
	if err := utils.CreateGitctxContextFileFromPath(contextName, contextPath); err != nil {
		t.Fatalf("Failed to create context file from path: %v", err)
	}

	// Verify the context file was created in the config folder
	configDir := paths.GetGitctxConfigFolderPath()
	if _, err := os.Stat(filepath.Join(configDir, contextName)); os.IsNotExist(err) {
		t.Fatalf("Context file '%s' was not created in the config folder", contextName)
	}

	fmt.Println("TestCreateGitctxContextFileFromPath completed successfully!")
}
