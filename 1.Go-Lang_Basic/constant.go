package main

import "fmt"

func main() {
	const firstName = "Lzyct"
	const lastName string = "Labs"

	//firstName = "Test" // error

	fmt.Println(firstName)
	fmt.Println(lastName)

	const(
		address = "Silicon Valley"
		phone = 123456
	)

	fmt.Println(address)
	fmt.Println(phone)
}