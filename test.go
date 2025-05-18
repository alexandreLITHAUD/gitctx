package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetGitFolder(startingPath string) string {
	testedPath := filepath.Join(startingPath, ".git")

	if info, err := os.Stat(testedPath); err == nil && info.IsDir() {
		return testedPath
	}

	// Stop if we reached the root
	parent := filepath.Dir(startingPath)
	if parent == startingPath {
		return ""
	}

	return GetGitFolder(parent)
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("folder %s\n", GetGitFolder(path))
}
