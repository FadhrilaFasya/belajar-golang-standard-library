package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout) // Create a new CSV writer that writes to standard output
	_ = writer.Write([]string{"John Doe", "john@example.com", "123 Main St"})   // Write a record
	_ = writer.Write([]string{"Jane Smith", "jane@example.com", "456 Elm St"})
	_ = writer.Write([]string{"Emily Davis", "emily@example.com", "789 Oak St"})

	writer.Flush() // Flush the writer to ensure all data is written
}