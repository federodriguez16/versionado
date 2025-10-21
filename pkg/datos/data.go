package datos

import (
	"fmt"
	"os"

	"github.com/federodriguez16/versionado/internal/database"
	"github.com/federodriguez16/versionado/internal/handler"
	"github.com/federodriguez16/versionado/internal/models"
)

func strPtr(s string) *string {
	return &s
}

func CargarDatos(f string) error {

	db, err := database.GetConnection()

	if err != nil {
		fmt.Println("Hubo un error de conexion!")
		os.Exit(1)
	}

	file := f
	records := handler.Archivo(file)

	clientsToCreate := make([]models.Client, 0, len(records)-1)

	for i, d := range records {
		if i == 0 {
			continue // Salta el header del archivo
		}

		client := models.Client{
			Name: strPtr(d[0]),
			Url:  d[1],
			Versions: models.Versions{
				ProductA: d[2],
				ProductB: d[3],
				ProductC: d[4],
				ProductD: d[5],
				ProductE: d[6],
			},
		}

		clientsToCreate = append(clientsToCreate, client)
	}

	if len(clientsToCreate) > 0 {
		result := db.Create(&clientsToCreate)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
