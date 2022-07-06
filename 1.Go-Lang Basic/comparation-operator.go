package main

import "fmt"

func main() {
	firstName := "Mudassir"
	lastName := "Mudassir"

	valueA := 10
	valueB := 20

	isSame := firstName == lastName

	fmt.Println("First name and last name is same ", isSame)

	fmt.Println("Value A < Value B", valueA < valueB)
	fmt.Println("Value A > Value B", valueA > valueB)
	fmt.Println("Value A >= Value B", valueA >= valueB)
	fmt.Println("Value A <= Value B", valueA <= valueB)
	fmt.Println("Value A != Value B", valueA != valueB)
}
