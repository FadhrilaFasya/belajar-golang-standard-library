package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(4) // Create a new ring with 4 elements
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i+1)
		data = data.Next()
	}

	// Iterate through the ring and print each element
	data.Do(func(value interface{}) { 
		fmt.Println(value)
	})
}