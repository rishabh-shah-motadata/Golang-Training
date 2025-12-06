package day6

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
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
		inputWg      sync.WaitGroup
		resultWg     sync.WaitGroup
		inputChan    = make(chan map[string]any, len(json.Rows)/2)
		resultChan   = make(chan string)
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
	copy(headers, tempHeaders[:headerLength])
	tempHeaders = nil

	csvRows = make([]string, len(json.Rows)+1)
	csvRows[0] = strings.Join(headers, ",")

	for range 10 {
		inputWg.Go(func() {
			csvConverterWorker(headerLength, headers, resultChan, inputChan)
		})
	}

	for _, row := range json.Rows {
		inputChan <- row
	}
	close(inputChan)

	resultWg.Go(func() {
		inputWg.Wait()
		close(resultChan)
	})

	index := 1
	for csvRow := range resultChan {
		csvRows[index] = csvRow
		index++
	}

	resultWg.Wait()

	csvContent := strings.Join(csvRows, "\n")
	if err := os.WriteFile("day_6/data.csv", []byte(csvContent), 0644); err != nil {
		return err
	}

	return nil
}

func csvConverterWorker(headerLength int, headers []string, resultChan chan<- string, row <-chan map[string]any) {
	csvRow := make([]string, headerLength)
	for data := range row {
		for key, value := range data {
			index := slices.Index(headers, key)
			if index != -1 {
				csvRow[index] = fmt.Sprintf("%v", value)
			}
		}
		fmt.Println("", csvRow)
		resultChan <- strings.Join(csvRow, ",")
		csvRow = make([]string, headerLength)
	}
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
