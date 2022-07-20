package main

import "fmt"

func main() {
	// 1
	var name string
	name = "Lazycat"
	fmt.Println(name)

	name = "Lazycat Labs"
	fmt.Println(name)

	var alias = "Lzyct"
	fmt.Println(alias)

	stars := 100
	fmt.Println(stars)

	var (
		firstName = "Lzyct"
		lastName  = "Labs"
		age       = 17
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}
