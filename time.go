package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now() 
	fmt.Println("Current Time:", currentTime)
	fmt.Println("Year:", currentTime.Year())
	
	utc := currentTime.UTC()
	fmt.Println("UTC Time:", utc)

	formater := "2022-01-02 15:04:05" 

	value := "2023-08-01 10:30:00"
	parsedTime, err := time.Parse(formater, value) // time.ParseInLocation(formater, value, time.Local)
	if err != nil {
		fmt.Println("Error parsing time:", err) // jika format tidak sesuai
	} else {
		fmt.Println("Parsed Time:", parsedTime) // jika format sesuai
	}
}