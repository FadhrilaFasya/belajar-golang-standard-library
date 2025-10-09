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

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField, "with type", valueField.Type)
	}
}
func main() {
	readField(Sample{Name: "Eko"})
	readField(Person{Name: "Budi", Address: "Indonesia", Email: "budi@example.com"})
}