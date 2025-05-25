package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"gopkg.in/ini.v1"
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

func GetGnuGpgKeys() []string {

	// Step 1: Run `gpg --list-signatures`
	cmd := exec.Command("gpg", "--list-signatures")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running gpg:", err)
		return nil
	}

	// Step 2: Process output line by line
	var results []string
	lines := strings.SplitSeq(out.String(), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "sig 3") {
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				results = append(results, parts[2])
			}
		}
	}

	var gpgKeys []string
	gpgKeys = append(gpgKeys, results...)

	return gpgKeys
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("folder %s\n", GetGitFolder(path))

	cfg, err := ini.Load("test.ini")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", cfg)

	gpgKeys := GetGnuGpgKeys()
	fmt.Printf("%#v\n", gpgKeys)

	prompt := promptui.Select{
		Label: "Slect one of your gpg keys",
		Items: gpgKeys,
		Size:  4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, gpgKeys[i])

}
