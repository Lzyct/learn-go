package main

import "fmt"

func main() {
	var (
		valueA = 10
		valueB = 20
		valueC = valueA + valueB
	)

	fmt.Println(valueC)

	valueC++ // valueC + 1
	fmt.Println(valueC)

	valueA += 25 // valueA = valueA + 25
	fmt.Println(valueA)
}
