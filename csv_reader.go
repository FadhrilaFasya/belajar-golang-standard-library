package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "John Doe\n" +
		"Jane Smith\n" +
		"Emily Davis"

	reader := csv.NewReader(strings.NewReader(csvString)) // Create a new CSV reader

	for {
		record, err := reader.Read() // Read each record
		if err == io.EOF {
			break // End of file
		}

		fmt.Println(record) // Print the record
	}
}