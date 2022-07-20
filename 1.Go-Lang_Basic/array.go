package main

import "fmt"

func main() {
	// if you define len from array
	// you can't append new data in numbers
	var numbers = [5]int{
		1, 2, 3, 4,
	}

	// if you not define len from array
	// you can append new data in numbers
	var unknownNumbers = [...]int{
		5, 4, 3, 2, 1,
	}

	fmt.Println(numbers) // [1 2 3 4 0]
	numbers[4] = 9
	fmt.Println(numbers[3])   // 4
	fmt.Println(numbers[4])   // 9
	fmt.Println(len(numbers)) // 5 : lenght from numbers

	fmt.Println(unknownNumbers)

}
