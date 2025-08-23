package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/types"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
)

func TestGetCurrentContextName(t *testing.T) {

	tmpDir := t.TempDir()
	paths.OverrideConfigFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	// Initially, the current context should be empty
	currentContext, err := utils.GetCurrentContextName()
	if err != nil {
		t.Fatalf("Error getting current context name: %v", err)
	}
	if currentContext != "" {
		t.Errorf("Expected initial current context to be empty, but got: '%s'", currentContext)
	}

	// Update the .config.json file to set a current context
	configFilePath := filepath.Join(paths.GetGitctxConfigFolderPath(), ".config.json")
	config := types.Config{CurrentContext: "testContext"}
	jsonData, err := json.Marshal(config)
	err = os.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to update config file: %v", err)
	}

	currentContext, err = utils.GetCurrentContextName()
	if err != nil {
		t.Fatalf("Error getting current context name after update: %v", err)
	}
	if currentContext != "testContext" {
		t.Errorf("Expected current context to be 'testContext', but got: '%s'", currentContext)
	}

	t.Logf("TestGetCurrentContextName Passed !")
}

func TestUpdateCurrentContextName(t *testing.T) {

	tmpDir := t.TempDir()
	paths.OverrideConfigFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	// Update the current context to "initialContext"
	err := utils.UpdateCurrentContextName("initialContext")
	if err != nil {
		t.Fatalf("Error updating current context name: %v", err)
	}

	currentContext, err := utils.GetCurrentContextName()
	if err != nil {
		t.Fatalf("Error getting current context name: %v", err)
	}
	if currentContext != "initialContext" {
		t.Errorf("Expected current context to be 'initialContext', but got: '%s'", currentContext)
	}

	// Update the current context to "newContext"
	err = utils.UpdateCurrentContextName("newContext")
	if err != nil {
		t.Fatalf("Error updating current context name: %v", err)
	}

	currentContext, err = utils.GetCurrentContextName()
	if err != nil {
		t.Fatalf("Error getting current context name: %v", err)
	}
	if currentContext != "newContext" {
		t.Errorf("Expected current context to be 'newContext', but got: '%s'", currentContext)
	}

	t.Logf("TestUpdateCurrentContextName Passed !")
}
