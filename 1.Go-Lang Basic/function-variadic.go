package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	result := sumAll(10, 11, 12, 13, 14, 15)
	fmt.Println("Total : ", result)

	slice := []int{16, 17, 18, 19, 20}
	result = sumAll(slice...)
	fmt.Println("Total : ", result)
}
