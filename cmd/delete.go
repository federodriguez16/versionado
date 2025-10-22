package cmd

import (
	"github.com/federodriguez16/versionado/pkg/datos"

	"github.com/spf13/cobra"
)

var clientDelete string

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&clientDelete, "client", "c", "", "Cliente")
}

var deleteCmd = &cobra.Command{
	Use:   "delete --client <nombre>",
	Short: "Remove the information of a client from the database",
	Long: `The delete command allows you to remove data associated with a specific client from the database.

It is mandatory to provide the client's name using the --client (or -c) flag.

Example usage:

versionado delete --client Juan
versionado delete -c Juan

Make sure the client's name is valid and exists in the database.`,

	Run: func(cmd *cobra.Command, args []string) {

		datos.EliminarDatos(clientDelete)

	},
}
