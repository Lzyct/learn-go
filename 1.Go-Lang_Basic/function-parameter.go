package main

import "fmt"

func githubInfo(username string, email string) {
	fmt.Println("Username :", username)
	fmt.Println("Email :", email)
}
func main() {
	githubInfo("Lzyct", "hey.mudassir@gmail.com")
}
