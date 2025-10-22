package cmd

import (
	"github.com/federodriguez16/versionado/pkg/datos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update <cliente> <servicio> <version>",
	Short: "Update a client's service",
	Long: `The update command allows you to modify a specific service of a client in the database.

You must provide the client identifier, the name of the field to update, and the new value.

Example usage:

versionado update Juan producto-a v2.0.0

This will update the producto-a field of the client Juan with the value v2.0.0.`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		datos.ActualizarDatos(args[0], args[1], args[2])
	},
}
