package main

import "fmt"

func getFullGithubInfo() (username string, stars int, ignoreMe string) {
	username = "Lzyct"
	stars = 69
	ignoreMe = ""
	return
}
func main() {
	name, stars, _ := getFullGithubInfo()

	fmt.Println("Username : ", name)
	fmt.Println("Stars : ", stars)
}
