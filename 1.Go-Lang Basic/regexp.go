package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Check if string have vowel char
	regex := regexp.MustCompile("x*?[aiueo]")

	fmt.Println(regex.MatchString("Lzyct"))   // false
	fmt.Println(regex.MatchString("Lazycat")) // true

}
