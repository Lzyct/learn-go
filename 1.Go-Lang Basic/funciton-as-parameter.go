package main

import (
	"fmt"
	"strings"
)

// Create type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

// Create spam filter
func spamFilter(name string) string {
	if name == "fuck" {
		return "****"
	} else {
		return name
	}
}

func uppercase(name string) string {
	return strings.ToUpper(name)
}

func main() {
	sayHelloWithFilter("Lzyct", spamFilter)
	sayHelloWithFilter("fuck", spamFilter)

	sayHelloWithFilter("Lzyct Labs", uppercase)
}
