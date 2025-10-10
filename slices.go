package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Jane", "Smith", "Doe"}
	values := []int{100, 90, 80, 95}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "John"))
	fmt.Println(slices.Index(names, "John"))
	fmt.Println(slices.Index(names, "Grey"))
}
