package paths

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
)

func TestGetHomeFolderPath(t *testing.T) {
	// Test with default home directory
	homeDir := paths.GetHomeFolderPath()
	if homeDir == "" {
		t.Error("Expected non-empty home directory path")
	}
	t.Logf("Default home directory: %s", homeDir)

	// Override home directory path
	tempDir := t.TempDir()
	paths.OverrideHomeFolderPath(tempDir)
	if paths.GetHomeFolderPath() != tempDir {
		t.Errorf("Expected overridden home directory path %s, got %s", tempDir, paths.GetHomeFolderPath())
	}

	// Reset override
	paths.OverrideHomeFolderPath("")
	if (paths.GetHomeFolderPath() == tempDir) && (paths.GetHomeFolderPath() != homeDir) {
		t.Error("Expected home directory path to reset after clearing override")
	}
}

func TestGetConfigFolderPath(t *testing.T) {
	// Test with default config directory
	configDir := paths.GetConfigFolderPath()
	if configDir == "" {
		t.Error("Expected non-empty config directory path")
	}
	t.Logf("Default config directory: %s", configDir)

	// Override config directory path
	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)
	if paths.GetConfigFolderPath() != tempDir {
		t.Errorf("Expected overridden config directory path %s, got %s", tempDir, paths.GetConfigFolderPath())
	}

	// Reset override
	paths.OverrideConfigFolderPath("")
	if (paths.GetConfigFolderPath() == tempDir) && (paths.GetConfigFolderPath() != configDir) {
		t.Error("Expected config directory path to reset after clearing override")
	}
}

func TestGetGitctxConfigFolderPath(t *testing.T) {
	// Test gitctx config folder path
	gitctxConfigDir := paths.GetGitctxConfigFolderPath()
	if gitctxConfigDir == "" {
		t.Error("Expected non-empty gitctx config directory path")
	}
	t.Logf("Default gitctx config directory: %s", gitctxConfigDir)

	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)
	expectedPath := filepath.Join(tempDir, "gitctx")
	if paths.GetGitctxConfigFolderPath() != expectedPath {
		t.Errorf("Expected gitctx config directory path %s, got %s", expectedPath, paths.GetGitctxConfigFolderPath())
	}

	// Reset override
	paths.OverrideConfigFolderPath("")
	if (paths.GetGitctxConfigFolderPath() == expectedPath) && (paths.GetGitctxConfigFolderPath() != gitctxConfigDir) {
		t.Error("Expected gitctx config directory path to reset after clearing override")
	}
}

func TestFindGitConfigsFilePaths(t *testing.T) {
	// Test finding git config file paths

	tempHomeDir := t.TempDir()
	tempConfigDir := t.TempDir()
	paths.OverrideHomeFolderPath(tempHomeDir)
	paths.OverrideConfigFolderPath(tempConfigDir)

	// Create mock git config files
	gitHomeFilePath := filepath.Join(tempHomeDir, ".gitconfig")
	gitConfigFilePath := filepath.Join(tempConfigDir, "git", "config")
	if err := os.MkdirAll(filepath.Dir(gitConfigFilePath), 0755); err != nil {
		t.Fatalf("Failed to create directory for git config file: %v", err)
	}
	configText := "[user]\nname = Test User\nemail = toto@fortnite.com\n"
	if err := os.WriteFile(gitHomeFilePath, []byte(configText), 0644); err != nil {
		t.Fatalf("Failed to create git home config file: %v", err)
	}
	if err := os.WriteFile(gitConfigFilePath, []byte(configText), 0644); err != nil {
		t.Fatalf("Failed to create git config file: %v", err)
	}

	paths, err := paths.FindGitConfigsFilePaths()
	if err != nil {
		t.Errorf("Error finding git config file paths: %v", err)
	}

	if len(paths) != 2 {
		t.Error("Expected to find two git config files")
	} else {
		t.Logf("Found git config file paths: %v", paths)
	}

	for _, path := range paths {
		if path != gitHomeFilePath && path != gitConfigFilePath {
			t.Errorf("Unexpected git config file path found: %s", path)
		}
		t.Logf("Valid git config file path: %s", path)
	}
}

func TestGetAllGitctxContexts(t *testing.T) {

	// Test getting all gitctx contexts
	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	utils.CreateGitctxConfigFolder()

	contexts := paths.GetAllGitctxContexts()
	if len(contexts) != 0 {
		t.Error("Expected no contexts initially")
	}

	// Create mock context files
	context1 := filepath.Join(paths.GetGitctxConfigFolderPath(), "context1")
	context2 := filepath.Join(paths.GetGitctxConfigFolderPath(), "context2")
	if err := os.WriteFile(context1, []byte("context1 data"), 0644); err != nil {
		t.Fatalf("Failed to create context file 1: %v", err)
	}
	if err := os.WriteFile(context2, []byte("context2 data"), 0644); err != nil {
		t.Fatalf("Failed to create context file 2: %v", err)
	}

	contexts = paths.GetAllGitctxContexts()
	if len(contexts) != 2 {
		t.Errorf("Expected 2 contexts, got %d", len(contexts))
	} else {
		t.Logf("Found contexts: %v", contexts)
	}
}
