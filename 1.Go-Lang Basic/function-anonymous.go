package main

import "fmt"

type Register func(name string) bool

func registerUser(name string, register Register) {
	if register(name) {
		fmt.Println("Welcome back", name)
	} else {
		fmt.Println("User ", name, " is not registered")
	}
}

func main() {
	isRegister := func(name string) bool {
		return name == "Lzyct"
	}

	registerUser("Lzyct", isRegister)
	registerUser("Lazycat", func(name string) bool {
		return name == "Lzyct"
	})
}
