package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required: true max: "10"`
}

type Person struct {
	Name     string `required: true max: "10"`
	Address  string `required: true max: "100"`
	Email    string `required: true max: "100"`
}

// readField function to read struct fields using reflection
func readField(value any) { 
	valueType := reflect.TypeOf(value) // get the type of the value
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ { // iterate over the fields
		structField := valueType.Field(i) // get the field at index i
		fmt.Println(structField, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}
func main() {
	readField(Sample{Name: "Eko"})
	readField(Person{Name: "Budi", Address: "Indonesia", Email: "budi@example.com"})
}