package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "127.0.0.1", "Database hostname")
	user := flag.String("user", "root", "Database user")
	password := flag.String("password", "toor", "Database password")

	// Need to run this, to override flag from arguments
	flag.Parse()

	// Result is pointer, so need put * to get value
	fmt.Println("Host :", *host)
	fmt.Println("User :", *user)
	fmt.Println("Password :", *password)
}
