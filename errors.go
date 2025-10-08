package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error") // Example of a custom error variable
	NotFoundError   = errors.New("not found error")   // Another custom error variable
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "John" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("Doe")
	if err != nil {
		if errors.Is(err, ValidationError) { // Using errors.Is to check for specific error types
			fmt.Println("Caught a validation error:", err)
		} else if errors.Is(err, NotFoundError) { // Using errors.Is to check for another error type
			fmt.Println("Caught a not found error:", err)
		}
	}
}