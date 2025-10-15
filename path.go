package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world.go"))
	fmt.Println(path.Split("hello/world.go"))
	fmt.Println(path.Clean("hello/../world.go"))
	fmt.Println(path.IsAbs("/hello/world.go"))
}