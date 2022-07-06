package main

import "fmt"

func main() {
	var numbers = [5]int{
		1, 2, 3, 4,
	}

	fmt.Println(numbers)      // [1 2 3 4 0]
	fmt.Println(numbers[3])   // 4
	fmt.Println(numbers[4])   // 0
	fmt.Println(len(numbers)) // 5 : lenght from numbers
}
