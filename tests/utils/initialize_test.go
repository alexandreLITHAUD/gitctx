package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
)

func TestInitializeGitctxConfigFolder(t *testing.T) {
	tempDir := t.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	utils.InitializeGitctxConfigFolder(true)

	// Ensure the gitctx config folder exists
	if _, err := os.Stat(paths.GetGitctxConfigFolderPath()); os.IsNotExist(err) {
		t.Fatalf("Gitctx config folder does not exist at path: %s", paths.GetGitctxConfigFolderPath())
	}

	fmt.Println("TestInitializeGitctxConfigFolder completed successfully!")
}
