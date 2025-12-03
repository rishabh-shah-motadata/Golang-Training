package day6

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type jsonReader struct {
	Rows []map[string]any `json:"rows"`
}

func convertToCSV(json jsonReader) error {
	if len(json.Rows) == 0 {
		log.Println("no data to convert")
		return nil
	}

	var (
		headers      []string
		csvRows      []string
		tempHeaders  = make([]string, 10)
		headerLength = 0
	)

	for _, row := range json.Rows {
		for key := range row {
			if !slices.Contains(tempHeaders, key) {
				tempHeaders[headerLength] = key
				headerLength++
			}
		}
	}

	headers = make([]string, headerLength)
	csvRows = make([]string, len(json.Rows)+1)
	copy(headers, tempHeaders[:headerLength])
	csvRows[0] = strings.Join(headers, ",")

	for i, row := range json.Rows {
		csvRow := make([]string, headerLength)
		for key, value := range row {
			index := slices.Index(headers, key)
			if index != -1 {
				csvRow[index] = fmt.Sprintf("%v", value)
			}
		}
		csvRows[i+1] = strings.Join(csvRow, ",")
	}

	csvContent := strings.Join(csvRows, "\n")
	if err := os.WriteFile("day_6/data.csv", []byte(csvContent), 0644); err != nil {
		return err
	}

	return nil
}

func Day6() {
	var jsonReader jsonReader

	fileContent, err := os.ReadFile("day_6/data.json")
	if err != nil {
		log.Println("failed to read json file", err)
	}

	if err := json.Unmarshal(fileContent, &jsonReader); err != nil {
		log.Println("failed to unmarshal json data", err)
	}

	convertToCSV(jsonReader)
}
