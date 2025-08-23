/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alexandreLITHAUD/gitctx/internal/paths"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available gitctx contexts",
	Long: ` This command lists all available gitctx contexts that have been created.
It helps you see which git configurations are available for switching.

You can use this command :
- gitctx list
The current context will be marked with an asterisk (*)`,
	Run: func(cmd *cobra.Command, args []string) {
		contexts := paths.ListAllGitctxContexts()
		if len(contexts) == 0 {
			fmt.Println("No contexts available. You can import one using 'gitctx init' or create one using 'gitctx create <context-name>'.")
			return
		}

		currentContext, err := utils.GetCurrentContextName()
		if err != nil {
			currentContext = ""
		}

		for _, context := range contexts {
			if context == currentContext {
				fmt.Printf("* %s\n", context)
			} else {
				fmt.Printf("  %s\n", context)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
