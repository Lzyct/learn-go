package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println("Index : ", i)
		if i == 6 {
			break
		}

	}
}
