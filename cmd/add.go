package cmd

import (
	"fmt"
	"os"

	"github.com/federodriguez16/versionado/pkg/datos"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [archivo.csv]",
	Short: "Create clients from a CSV file",
	Long: `The add command allows you to load records from a specified file and store them in the database.

You must provide the file path as an argument. The command validates the database connection and processes the file, inserting the data into the corresponding table.

Example usage:

versionado add datos.csv


The file must have a valid format, as the second row will be used to extract the records.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		result := datos.CargarDatos(file)

		if result != nil {
			fmt.Println("Hubo un error en la carga de datos!")
			os.Exit(1)
		}
	},
}
