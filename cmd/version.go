package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the CLI version",
	Long: `The version command prints the current version of this command-line tool (CLI).

It is useful to verify compatibility or to ensure you are using the expected version.

Example usage:

versionado version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli version v1.0.0")
	},
}
