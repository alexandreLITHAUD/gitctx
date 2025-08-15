package utils

import (
	"fmt"
	"os"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/ui"
)

func InitializeGitctxConfigFolder(excludeFetch bool) error {
	// Ensure the gitctx config folder exists
	if _, err := os.Stat(paths.GetGitctxConfigFolderPath()); os.IsNotExist(err) {
		CreateGitctxConfigFolder()
	}

	if excludeFetch {
		fmt.Println("Initialization complete !")
		return nil
	}

	// Check or recheck the git configs file paths
	gitConfigsPaths, err := paths.FindGitConfigsFilePaths()
	if err != nil {
		return err
	}

	if len(gitConfigsPaths) == 0 {
		fmt.Printf("No git config files found in the common locations.\n")
	} else {
		for _, path := range gitConfigsPaths {

			//TODO: Add a ignore system (as a param)
			//TODO: Add prompttui ui for that ! (in a dedicated function)

			fmt.Printf("Found git config file: %s\n", path)

			confirmed, err := ui.PromptForConfirmation("Would you like to import it as a context ?\n  Answer (y)es or (n)o :")
			if err != nil {
				return fmt.Errorf("error reading confirmation input: %w", err)
			}
			if !confirmed {
				fmt.Println("Skipping import of this config file.")
				continue
			}

			name, err := ui.PromptForName("What should be the name of this context ?\n Enter context name:")
			if err != nil {
				return fmt.Errorf("error reading context name input: %w", err)
			}

			CreateGitctxContextFile(name, path)
			fmt.Printf("Context '%s' created from config file: %s\n", name, path)
		}
	}
	fmt.Println("Initialization complete !")
	return nil
}
