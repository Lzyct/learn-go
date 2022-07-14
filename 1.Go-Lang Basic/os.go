package main

import (
	"fmt"
	"os"
)

/**
Read more on https://pkg.go.dev/os#pkg-overview
for documentation and example
*/
func main() {
	args := os.Args
	fmt.Println("Arguments : ")
	fmt.Println(args)

	// get hostname
	name, err := os.Hostname()

	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error :", err.Error())
	}

	// get os env
	flutterHome := os.Getenv("FLUTTER_HOME")
	androidHome := os.Getenv("ANDROID_HOME")

	fmt.Println("Flutter home :", flutterHome)
	fmt.Println("Android home :", androidHome)

}
