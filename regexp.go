package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))      // true
	fmt.Println(regex.MatchString("Eko"))	  // false
	fmt.Println(regex.MatchString("eka"))      // false

	fmt.Println(regex.FindAllString("eko edo egi eto eta itu", 10)) // [eko edo egi eto eta itu]
}