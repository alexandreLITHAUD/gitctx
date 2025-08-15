/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gitctx configuration folder and contexts",
	Long: `This command initializes the gitctx configuration folder and contexts by searching for
existing git configuration files in common locations:

Here are the common locations:
- ~/.gitconfig
- ~/.config/git/config

Here are some examples of how to use the init command:
- gitctx init
- gitctx init --no-fetch-context (to skip fetching contexts from git config files)`,
	Run: func(cmd *cobra.Command, args []string) {
		excludeFetch, err := cmd.Flags().GetBool("no-fetch-context")
		if err != nil {
			excludeFetch = false
		}
		utils.InitializeGitctxConfigFolder(excludeFetch)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("no-fetch-context", "n", false, "Skip fetching contexts from git config files")
}
