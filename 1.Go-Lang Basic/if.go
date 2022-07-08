package main

import "fmt"

func main() {
	name := "LazyCat Labs"

	if name == "LazyCat" {
		fmt.Println("Hello ", name)
	} else if name == "Lzyct" {
		fmt.Println("Whatsup ", name)
	} else {
		fmt.Println("Who are you?")
	}

	/// short hand condition
	if length := len(name); length < 6 {
		fmt.Println("Name too short")
	} else {
		fmt.Println("Good name ", name)
	}
}
