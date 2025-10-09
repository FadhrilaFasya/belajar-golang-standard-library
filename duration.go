package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := 100 * time.Second
	duration2 := 10 * time.Millisecond
	duration := duration1 + duration2

	fmt.Println("Duration:", duration)
	fmt.Printf("Duration in Hours: %.2f\n", duration.Hours())

}