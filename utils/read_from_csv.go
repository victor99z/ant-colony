package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	PosX, PosY float64
	Label      int
}

func ReadData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = ';'
	r.Comment = '#'

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func ParseData(fileName string) []Data {
	records, err := ReadData(fileName)
	if err != nil {
		panic(err)
	}

	listItems := []Data{}

	for _, record := range records {
		posx, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Printf("Error converting string: %v", err)
		}
		posy, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Printf("Error converting string: %v", err)
		}
		label, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Printf("Error converting string: %v", err)
		}
		listItems = append(listItems, Data{posx, posy, label})
	}
	return listItems
}
