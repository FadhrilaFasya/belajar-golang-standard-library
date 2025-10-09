package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name, Address, Email string
}

// readField function to read struct fields using reflection
func readField(value any) { 
	valueType := reflect.TypeOf(value) // get the type of the value

	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ { // iterate over the fields
		valueField := valueType.Field(i) // get the field at index i
		fmt.Println(valueField, "with type", valueField.Type)
	}
}
func main() {
	readField(Sample{Name: "Eko"})
	readField(Person{Name: "Budi", Address: "Indonesia", Email: "budi@example.com"})
}