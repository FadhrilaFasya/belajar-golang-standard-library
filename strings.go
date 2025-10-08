package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, World!", "World")) // true
	fmt.Println(strings.Split("a,b,c", ","))		  // [a b c]
	fmt.Println(strings.ToUpper("hello"))		  // HELLO
	fmt.Println(strings.TrimSpace("  hello  "))	  // "hello"
	fmt.Println(strings.ToLower("HELLO"))		  // hello
	fmt.Println(strings.Repeat("go", 3))		  // gogogo
	fmt.Println(strings.ReplaceAll("foo bar foo", "foo", "baz")) // baz bar baz
	fmt.Println(strings.HasPrefix("Golang", "Go")) // true
	fmt.Println(strings.HasSuffix("Golang", "lang")) // true
}