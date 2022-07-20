package main

import "fmt"

func main() {
	name := "Lazycat Labs"

	switch name {
	case "Lazycat Labs", "Lzyct":
		fmt.Println("Hello ", name)
	default:
		fmt.Println("Who are you ?")
	}

	// switch with short hand
	switch length := len(name); length < 6 {
	case true:
		fmt.Println("Name too short")
	case false:
		fmt.Println("Good name ", name)
	}

	// switch without value
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name too long")
	case length > 6:
		fmt.Println("Name is long")
	default:
		fmt.Println("Good name ", name)
	}
}
