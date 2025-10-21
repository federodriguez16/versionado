package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Archivo(a string) [][]string {
	file, err := os.Open(a)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	data, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return data
}
