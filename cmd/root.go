/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gitctx",
	Short: "gitctx is a tool to manage git contexts",
	Long: `gitctx helps you manage multiple git configurations and contexts easily.
It allows you to switch between different git user profiles, repositories, and settings quickly.

If this is your first time using gitctx, you can initialize it with the 'gitctx init' command.
It will set up the necessary configuration files and directories and store your current git configuration.

For more information, visit https://alexandrelithaud.github.io/gitctx/`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			os.Exit(1)
		}

		if Verbose {
			fmt.Println("Verbose mode enabled")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose output")
}
