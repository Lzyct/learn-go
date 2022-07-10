package main

import "fmt"

func squareArea(width int, height int) int {
	if width == 0 {
		return -1
	}
	return width * height
}
func main() {
	result := squareArea(10, 6)
	result1 := squareArea(0, 6)

	fmt.Println(result)
	fmt.Println(result1)
}
