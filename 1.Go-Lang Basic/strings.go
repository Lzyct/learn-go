package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Lzyct Labs"

	fmt.Println(strings.Contains(name, "Lzyct")) //true

	splitName := strings.Split(name, " ")

	for i, value := range splitName {
		fmt.Println(i, ":", value)
	}

	fmt.Println(strings.ToLower(name)) //lzyct labs
	fmt.Println(strings.ToUpper(name)) //LZYCT LABS
	// Need to set lowercase first to make Title works
	fmt.Println(strings.Title(strings.ToLower(name))) //Lzyct Labs

	// Remove all char in start and end of string by cutset
	fmt.Println(strings.Trim("  "+name+"   ", " "))

	fmt.Println(strings.ReplaceAll(name, " ", " ❤ ️"))
}
