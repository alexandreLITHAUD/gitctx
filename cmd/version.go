/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var Version string = "undefined"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return your gitctx version and architecture",
	Long: `This command return your current gitctx version, operating system, and architecture.
It is useful to check if you are using the correct version of gitctx for your system.

Here are tge commands to catch the version and architecture:
For version : gitctx version -q'
For operating system : gitctx version | awk '{print $4}' | cut -d/ -f1
For architecture : gitctx version | awk '{print $4}' | cut -d/ -f2`,
	Run: func(cmd *cobra.Command, args []string) {
		if quiet, _ := cmd.Flags().GetBool("quiet"); quiet {
			fmt.Println(Version)
			return
		}
		fmt.Printf("gitctx version %s %s/%s\n", Version, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("quiet", "q", false, "Suppress output to just version")
}
