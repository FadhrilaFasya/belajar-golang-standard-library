package main

import "os"

func main() {
	args := os.Args
	for _, arg := range args {
		println(arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		println("Error getting hostname:", err)
	} else {
		println("Hostname:", hostname)
	}
}