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

func TestGetCurrentGlobalContextName(t *testing.T) {

	tmpDir := t.TempDir()
	paths.OverrideConfigFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	// Initially, the current global context should be empty
	currentGlobalContext, err := utils.GetCurrentGlobalContextName()
	if err != nil {
		t.Fatalf("Error getting current global context name: %v", err)
	}
	if currentGlobalContext != "" {
		t.Errorf("Expected initial current global context to be empty, but got: '%s'", currentGlobalContext)
	}

	updateErr := utils.UpdateCurrentGlobalContextName("testGlobalContext")
	if updateErr != nil {
		t.Fatalf("Error updating current global context name: %v", updateErr)
	}

	currentGlobalContext, err = utils.GetCurrentGlobalContextName()
	if err != nil {
		t.Fatalf("Error getting current global context name after update: %v", err)
	}
	if currentGlobalContext != "testGlobalContext" {
		t.Errorf("Expected current global context to be 'testGlobalContext', but got: '%s'", currentGlobalContext)
	}

	t.Logf("TestGetCurrentGlobalContextName Passed !")
}

func TestUpdateCurrentGlobalContextName(t *testing.T) {

	tmpDir := t.TempDir()
	paths.OverrideConfigFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	// Update the current global context to "initialGlobalContext"
	updateErr := utils.UpdateCurrentGlobalContextName("initialGlobalContext")
	if updateErr != nil {
		t.Fatalf("Error updating current global context name: %v", updateErr)
	}

	currentGlobalContext, err := utils.GetCurrentGlobalContextName()
	if err != nil {
		t.Fatalf("Error getting current global context name: %v", err)
	}
	if currentGlobalContext != "initialGlobalContext" {
		t.Errorf("Expected current global context to be 'initialGlobalContext', but got: '%s'", currentGlobalContext)
	}

	// Update the current global context to "newGlobalContext"
	updateErr = utils.UpdateCurrentGlobalContextName("newGlobalContext")
	if updateErr != nil {
		t.Fatalf("Error updating current global context name: %v", updateErr)
	}

	currentGlobalContext, err = utils.GetCurrentGlobalContextName()
	if err != nil {
		t.Fatalf("Error getting current global context name: %v", err)
	}
	if currentGlobalContext != "newGlobalContext" {
		t.Errorf("Expected current global context to be 'newGlobalContext', but got: '%s'", currentGlobalContext)
	}

	t.Logf("TestUpdateCurrentGlobalContextName Passed !")
}

func TestGetConfigFields(t *testing.T) {

	tmpDir := t.TempDir()
	paths.OverrideConfigFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	// Update the current global context to "initialGlobalContext"
	updateErr := utils.UpdateCurrentGlobalContextName("initialGlobalContext")
	if updateErr != nil {
		t.Fatalf("Error updating current global context name: %v", updateErr)
	}

	// Update the current global context to "newGlobalContext"
	updateErr = utils.UpdateCurrentGlobalContextName("newGlobalContext")
	if updateErr != nil {
		t.Fatalf("Error updating current global context name: %v", updateErr)
	}

	currentGlobalContext, err := utils.GetCurrentGlobalContextName()
	if err != nil {
		t.Fatalf("Error getting current global context name: %v", err)
	}
	if currentGlobalContext != "newGlobalContext" {
		t.Errorf("Expected current global context to be 'newGlobalContext', but got: '%s'", currentGlobalContext)
	}

	config, err := utils.GetConfigFields()
	if err != nil {
		t.Fatalf("Error getting config fields: %v", err)
	}

	if config.CurrentGlobalContext != "newGlobalContext" {
		t.Errorf("Expected current global context to be 'newGlobalContext', but got: '%s'", config.CurrentGlobalContext)
	}

	t.Logf("TestGetConfigFields Passed !")
}

func TestApplyContext(t *testing.T) {

	tmpDir := t.TempDir()
	err := os.Mkdir(filepath.Join(tmpDir, ".git"), 0755)
	if err != nil {
		t.Fatalf("Failed to create .git directory: %v", err)
	}
	paths.OverrideConfigFolderPath(tmpDir)
	paths.OverrideLocalGitFolderPath(tmpDir)
	paths.OverrideHomeFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	err = utils.CreateGitctxContextFromScratch("testContext")
	if err != nil {
		t.Fatalf("Failed to create test context: %v", err)
	}
	err = os.WriteFile(filepath.Join(paths.GetGitctxConfigFolderPath(), "testContext"), []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test context file: %v", err)
	}

	err = utils.UpdateCurrentContextName("testContext")
	if err != nil {
		t.Fatalf("Failed to update current context: %v", err)
	}

	// Test ApplyConfig function // local
	err = utils.ApplyContext(false)
	if err != nil {
		t.Errorf("Error applying config: %v", err)
	}

	_, err = os.Stat(filepath.Join(paths.GetLocalGitFolderPath(), "config"))
	if err != nil {
		t.Errorf("Error applying config: %v", err)
	}

	// Test ApplyConfig function // global
	err = utils.ApplyContext(true)
	if err != nil {
		t.Errorf("Error applying config: %v", err)
	}

	_, err = os.Stat(filepath.Join(paths.GetHomeFolderPath(), ".gitconfig"))
	if err != nil {
		t.Errorf("Error applying config: %v", err)
	}

	t.Logf("TestApplyContext Passed !")
}
