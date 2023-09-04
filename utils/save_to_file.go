package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func SaveToFile(environment *Enviroment, filename string) {
	// Create a new CSV file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the matrix data to the CSV file
	for _, row := range environment.GetAll() {
		// Convert each integer to a string before writing to CSV
		stringRow := make([]string, len(row))
		for i, val := range row {
			stringRow[i] = strconv.Itoa(val)
		}

		if err := writer.Write(stringRow); err != nil {
			log.Fatal("Error writing row to CSV:", err)
		}
	}

	log.Println("CSV file 'output.csv' has been created successfully.")
}
