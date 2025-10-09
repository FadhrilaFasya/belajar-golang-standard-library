package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New() // Create a new list

	// Add elements to the list
	data.PushBack("John")
	data.PushBack("Doe")
	data.PushBack("Jane")
	data.PushBack("Smith")

	// Iterate through the list and print each element
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}