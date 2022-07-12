package main

import "fmt"

func dynamic(i int) interface{} {
	if i == 1 {
		return 2
	} else if i == 2 {
		return "Ups"
	} else {
		return false
	}
}

func main() {
	dynamicData := dynamic(4)

	var isFalse bool
	isFalse = dynamicData.(bool) // define dynamic data as bool
	fmt.Println(isFalse)
}
