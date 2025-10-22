package cmd

import (
	"github.com/federodriguez16/versionado/pkg/datos"

	"github.com/spf13/cobra"
)

var clientList string

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&clientList, "client", "c", "", "Cliente")
}

var listCmd = &cobra.Command{
	Use:   "list --client <nombre>",
	Short: "List versions of a client",
	Long: `The list command allows you to display the stored data for a specific client or for all clients.

You can use the --client (or -c) flag to filter by a specific client. If this flag is not provided, data for all clients will be listed (if implemented that way).

Example usage:

versionado list --client Juan
versionado list -c Juan
versionado list    # Lists all clients`,
	Run: func(cmd *cobra.Command, args []string) {
		datos.ListarDatos(clientList)
	},
}
