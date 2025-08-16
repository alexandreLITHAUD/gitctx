package paths

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
)

func BenchmarkGetGitctxFolderConfig(b *testing.B) {

	tempDir := b.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	// Benchmark the GetGitctxConfigFolderPath function
	for i := 0; i < b.N; i++ {
		_ = paths.GetGitctxConfigFolderPath()
	}
}

func BenchmarkGetGitctxContexts(b *testing.B) {
	tempDir := b.TempDir()
	paths.OverrideConfigFolderPath(tempDir)

	// Create a temporary gitctx config folder
	gitctxConfigFolder := paths.GetGitctxConfigFolderPath()
	if err := utils.CreateGitctxConfigFolder(); err != nil {
		b.Fatalf("Failed to create gitctx config folder: %v", err)
	}

	// Create mock context files
	context1 := filepath.Join(gitctxConfigFolder, "context1")
	context2 := filepath.Join(gitctxConfigFolder, "context2")
	if err := os.WriteFile(context1, []byte("context1 data"), 0644); err != nil {
		b.Fatalf("Failed to create context file 1: %v", err)
	}
	if err := os.WriteFile(context2, []byte("context2 data"), 0644); err != nil {
		b.Fatalf("Failed to create context file 2: %v", err)
	}

	// Benchmark the GetAllGitctxContexts function
	for i := 0; i < b.N; i++ {
		_ = paths.GetAllGitctxContexts()
	}
}
