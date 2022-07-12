package main

import "fmt"

func main() {
	name := "Lzyct"

	patch := func() {
		// create local variable
		name := "Labs"
		fmt.Println(name)
	}
	patch()

	fmt.Println(name)
}
