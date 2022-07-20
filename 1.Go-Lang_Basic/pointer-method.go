package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) marriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Lzyct"}
	man.married()
	// Not updated because using "pass by value"
	fmt.Println(man) // Lzyct

	manPointer := Man{"Lzyct"}
	manPointer.marriedPointer()
	// Updated because using "pass by reference"
	fmt.Println(manPointer) // Mr. Lzyct
}
