package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "username for login")
	var password *string = flag.String("password", "root", "password for login")
	var host *string = flag.String("host", "localhost", "hostname for the server")
	var port *int = flag.Int("port", 8080, "port number for the server")

	flag.Parse()

	fmt.Println("Server Configuration:")
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
}