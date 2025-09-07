/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alexandreLITHAUD/gitctx/internal/utils"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the current gitctx context",
	Long: `This command applies the current gitctx context to your git configuration.
It will replace your current git configuration with the settings from the current gitctx context. You can choose to apply the changes locally or globally and chose the context to apply.

If you want to apply the changes locally, use the --local flag. It will then replace the config file in your current git repository.

You can use this command :
- gitctx apply
- gitctx apply --local
- gitctx apply --name <context-name> --local
`,
	Run: func(cmd *cobra.Command, args []string) {

		local, _ := cmd.Flags().GetBool("local")
		name, _ := cmd.Flags().GetString("name")

		if name != "" {
			err := utils.UpdateCurrentContextName(name)
			if err != nil {
				fmt.Println("Context name not found")
				return
			}
		}

		err := utils.ApplyContext(!local)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)
	applyCmd.Flags().BoolP("local", "l", false, "Apply the changes locally")
	applyCmd.Flags().StringP("name", "n", "", "Name of the context to apply")
}
