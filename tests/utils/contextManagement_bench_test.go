package utils

import (
	"strconv"
	"testing"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
)

func BenchmarkCreateGitctxContextFromScratch(b *testing.B) {

	tmpDir := b.TempDir()
	paths.OverrideConfigFolderPath(tmpDir)
	utils.CreateGitctxConfigFolder()

	for i := 0; i < b.N; i++ {
		str := "benchmarkContext_" + strconv.Itoa(i)
		err := utils.CreateGitctxContextFromScratch(str)
		if err != nil {
			b.Fatalf("Error creating context from scratch: %v", err)
		}
	}

}
