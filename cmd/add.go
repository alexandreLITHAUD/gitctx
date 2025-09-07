/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/alexandreLITHAUD/gitctx/internal/ui"
	"github.com/alexandreLITHAUD/gitctx/internal/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new gitctx context",
	Long: `Add a new context to your gotcontext list.
You can set the name and path to the git config file to import. if the name is not set then a prompt will be displayed to enter the name.

You can use this command :
- gitctx add
- gitctx add --name <context-name>
- gitctx add --from <path-to-git-config-file>
- gitctx add --name <context-name> --from <path-to-git-config-file> --set-current
`,
	Run: func(cmd *cobra.Command, args []string) {

		from, _ := cmd.Flags().GetString("from")
		setCurrent, _ := cmd.Flags().GetBool("set-current")
		name, _ := cmd.Flags().GetString("name")

		if name == "" {
			name, _ = ui.PromptForName("Please enter a name for the context:")
		}

		if from == "" {
			err := utils.CreateGitctxContextFromScratch(name)
			if err != nil {
				panic(err)
			}
			if setCurrent {
				err := utils.UpdateCurrentContextName(name)
				if err != nil {
					panic(err)
				}
			}
		} else {
			err := utils.CreateGitctxContextFileFromPath(name, from)
			if err != nil {
				panic(err)
			}
			if setCurrent {
				err := utils.UpdateCurrentContextName(name)
				if err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("--name", "n", "", "Name of the context to create")
	addCmd.Flags().StringP("--from", "f", "", "Path to the git config file to import")
	addCmd.Flags().BoolP("--set-current", "s", false, "Set the imported context as the current context")
}
