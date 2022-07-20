package main

import "fmt"

func main() {
	var (
		isTrue  = true
		isFalse = false
	)

	fmt.Println("Both is True", isTrue && isFalse)
	fmt.Println("One of is True", isTrue || isFalse)
	fmt.Println("Negation of True is ", !isTrue)

}
