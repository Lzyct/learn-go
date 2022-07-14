package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) // 2
	fmt.Println(math.Round(1.3)) // 1

	fmt.Println(math.Floor(3.9)) // 3
	fmt.Println(math.Ceil(2.2))  //3

	fmt.Println(math.Max(10, 99)) //99
	fmt.Println(math.Min(10, 99)) //10
}
