package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "Prince Raj"
	records := readCsvFile("./name_place.csv")

	// fmt.Println(records)
	place(input, records)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func place(inputs string, records [][]string) {

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			// fmt.Println(records[i][j], i, j)
			if records[i][j] == inputs {
				fmt.Println(records[i][j+1])
				break
			}
		}
	}

}
