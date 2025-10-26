package cmd

import (
	"fmt"

	"github.com/federodriguez16/versionado/pkg/datos"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(urlCmd)
}

var urlCmd = &cobra.Command{
	Use:   "url [client]",
	Short: "Display the URL from a specific client",
	Long: `Displays the URL corresponding to the specified client.

This command fetches and returns the FQDN  associated with the client provided as an argument.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := args[0]

		url := datos.MostrarUrl(client)

		if url == "" {
			fmt.Println("No se pudo obtener el dominio")
		} else {
			fmt.Printf("El dominio de %s es: %s\n", client, url)
		}

	},
}
