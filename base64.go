package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "This is a sample code"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// Decode the base64 string
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		println("Error decoding:", err.Error())
		return
	}
	fmt.Println(string(decoded))
}