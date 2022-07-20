package main

import "fmt"

func getPerson() (string, int, string, string) {
	return "Lzyct", 17, "ignoreMe", "ignoreMeAgain"
}
func main() {
	name, age, _, _ := getPerson()
	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age, " years")
}
